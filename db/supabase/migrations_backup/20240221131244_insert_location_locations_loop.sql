-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION upsert_location(
    p_sd_code CHAR(2),
    p_sgg_code CHAR(3),
    p_emd_code CHAR(3),
    p_road_code CHAR(7),
    p_building_main_code VARCHAR(5),
    p_building_sub_code VARCHAR(5),
    p_sd_name VARCHAR(40),
    p_sgg_name VARCHAR(40),
    p_emd_name VARCHAR(40),
    p_road_name VARCHAR(40),
    p_building_name VARCHAR(40),
    p_location_5179 gis.GEOMETRY(POINT, 5179),
    p_location_4326 gis.GEOMETRY(POINT, 4326)
) RETURNS VOID AS
$$
DECLARE
    v_constraint_name TEXT;
BEGIN
    LOOP
        BEGIN
            INSERT INTO location.locations (sd_code, sgg_code, emd_code, road_code, building_main_code,
                                            building_sub_code,
                                            sd_name, sgg_name, emd_name, road_name, building_name, location_5179,
                                            location_4326)
            VALUES (p_sd_code, p_sgg_code, p_emd_code, p_road_code, p_building_main_code, p_building_sub_code,
                    p_sd_name, p_sgg_name, p_emd_name, p_road_name, p_building_name, p_location_5179, p_location_4326);
            EXIT; -- Exit loop if INSERT succeeds
        EXCEPTION
            WHEN unique_violation THEN
                GET STACKED DIAGNOSTICS v_constraint_name := CONSTRAINT_NAME;
                IF (v_constraint_name = 'chk_unique_code') OR (v_constraint_name = 'chk_unique_name') THEN
                    -- Handle the violation for the first unique constraint
                    UPDATE location.locations
                    SET sd_code            = p_sd_code,
                        sgg_code           = p_sgg_code,
                        emd_code           = p_emd_code,
                        road_code          = p_road_code,
                        building_main_code = p_building_main_code,
                        building_sub_code  = p_building_sub_code,
                        sd_name            = p_sd_name,
                        sgg_name           = p_sgg_name,
                        emd_name           = p_emd_name,
                        road_name          = p_road_name,
                        building_name      = p_building_name,
                        location_5179      = p_location_5179,
                        location_4326      = p_location_4326
                    WHERE sd_code = p_sd_code
                      AND sgg_code = p_sgg_code
                      AND emd_code = p_emd_code
                      AND road_code = p_road_code
                      AND building_main_code = p_building_main_code
                      AND building_sub_code = p_building_sub_code;
                ELSE
                    RAISE NOTICE 'Unexpected unique constraint violation: %', v_constraint_name;
                    EXIT; -- Optionally, exit or take other actions
                END IF;
                EXIT WHEN FOUND; -- Exit loop if UPDATE succeeds
        END;
    END LOOP;
END;
$$ LANGUAGE plpgsql;


DO
$$
    DECLARE
        r RECORD;
    BEGIN
        FOR r IN
            SELECT LEFT(sgg_code_raw, 2)                                                AS sd_code,
                   RIGHT(sgg_code_raw, 3)                                               AS sgg_code,
                   LEFT(RIGHT(bjd_code, 6), 3)                                          AS emd_code,
                   RIGHT(road_code_raw, 7)                                              AS road_code,
                   building_main_code,
                   NULLIF(NULLIF(building_sub_code, '0'), '')                           AS building_sub_code,
                   sd_name,
                   sgg_name,
                   emd_name,
                   road_name,
                   building_name,
                   gis.ST_SetSRID(gis.ST_MakePoint(x, y), 5179)                         AS location_5179,
                   gis.ST_Transform(gis.ST_SetSRID(gis.ST_MakePoint(x, y), 5179), 4326) AS location_4326
            FROM raw.locations_summary
            WHERE x IS NOT NULL
              AND y IS NOT NULL
            LOOP
                PERFORM upsert_location(r.sd_code, r.sgg_code, r.emd_code, r.road_code, r.building_main_code,
                                        r.building_sub_code,
                                        r.sd_name, r.sgg_name, r.emd_name, r.road_name, r.building_name,
                                        r.location_5179, r.location_4326);
            END LOOP;
    END;
$$;

DROP FUNCTION IF EXISTS upsert_location(p_sd_code CHAR(2), p_sgg_code CHAR(3), p_emd_code CHAR(3), p_road_code CHAR(7),
                                        p_building_main_code VARCHAR(5), p_building_sub_code VARCHAR(5),
                                        p_sd_name VARCHAR(40), p_sgg_name VARCHAR(40), p_emd_name VARCHAR(40),
                                        p_road_name VARCHAR(40), p_building_name VARCHAR(40),
                                        p_location_5179 gis.GEOMETRY(POINT, 5179),
                                        p_location_4326 gis.GEOMETRY(POINT, 4326));
-- +goose StatementEnd

-- +goose Down
-- noinspection SqlWithoutWhere
DELETE
FROM location.locations;

DROP FUNCTION IF EXISTS upsert_location(p_sd_code CHAR(2), p_sgg_code CHAR(3), p_emd_code CHAR(3), p_road_code CHAR(7),
                                        p_building_main_code VARCHAR(5), p_building_sub_code VARCHAR(5),
                                        p_sd_name VARCHAR(40), p_sgg_name VARCHAR(40), p_emd_name VARCHAR(40),
                                        p_road_name VARCHAR(40), p_building_name VARCHAR(40),
                                        p_location_5179 gis.GEOMETRY(POINT, 5179),
                                        p_location_4326 gis.GEOMETRY(POINT, 4326));