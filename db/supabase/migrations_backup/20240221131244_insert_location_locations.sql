-- +goose Up
INSERT INTO location.locations (sd_code,
                                sgg_code,
                                emd_code,
                                road_code,
                                building_main_code,
                                building_sub_code,
                                sd_name,
                                sgg_name,
                                emd_name,
                                road_name,
                                building_name,
                                location_5179,
                                location_4326)
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
ON CONFLICT ON CONSTRAINT chk_locations_unique_code_name
    DO UPDATE SET sd_code            = EXCLUDED.sd_code,
                  sgg_code           = EXCLUDED.sgg_code,
                  emd_code           = EXCLUDED.emd_code,
                  road_code          = EXCLUDED.road_code,
                  building_main_code = EXCLUDED.building_main_code,
                  building_sub_code  = EXCLUDED.building_sub_code,
                  sd_name            = EXCLUDED.sd_name,
                  sgg_name           = EXCLUDED.sgg_name,
                  emd_name           = EXCLUDED.emd_name,
                  road_name          = EXCLUDED.road_name,
                  building_name      = EXCLUDED.building_name,
                  location_5179      = EXCLUDED.location_5179,
                  location_4326      = EXCLUDED.location_4326,
                  duplicate_flag     = TRUE;

-- +goose Down
-- noinspection SqlWithoutWhere
DELETE
FROM location.locations;