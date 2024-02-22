-- +goose Up
CREATE TEMP TABLE staged_locations AS
SELECT *
FROM location.locations WITH NO DATA;

INSERT INTO staged_locations (sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code, sd_name,
                              sgg_name, emd_name, road_name, building_name, location_5179, location_4326)
SELECT LEFT(sgg_code_raw, 2)                      as sd_code,
       RIGHT(sgg_code_raw, 3)                     as sgg_code,
       SUBSTRING(bjd_code FROM 6 FOR 3)           as emd_code,
       RIGHT(road_code_raw, 7)                    as road_code,
       building_main_code,
       NULLIF(NULLIF(building_sub_code, '0'), '') as building_sub_code,
       sd_name,
       sgg_name,
       emd_name,
       road_name,
       building_name,
       gis.ST_SetSRID(gis.ST_MakePoint(x, y), 5179),
       gis.ST_Transform(gis.ST_SetSRID(gis.ST_MakePoint(x, y), 5179), 4326)
FROM raw.locations_summary
WHERE x IS NOT NULL
  AND y IS NOT NULL;

WITH generated_staged_locations AS (SELECT *,
                                           (sd_name || COALESCE(' ' || sgg_name, '') || CASE
                                                                                            WHEN RIGHT(emd_name, 1) = '동'
                                                                                                THEN ''
                                                                                            ELSE ' ' || emd_name END) AS sse_name,
                                           (road_name || ' ' || building_main_code || CASE
                                                                                          WHEN building_sub_code IS NOT NULL
                                                                                              THEN '-' || building_sub_code
                                                                                          ELSE '' END)                AS rb_name
                                    FROM staged_locations),
     ranked_and_counted AS (SELECT *,
                                   ROW_NUMBER() OVER (
                                       PARTITION BY sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code
                                       ORDER BY (building_name IS NOT NULL) DESC, id
                                       ) AS code_row_number,
                                   ROW_NUMBER() OVER (
                                       PARTITION BY sse_name, rb_name
                                       ORDER BY (building_name IS NOT NULL) DESC, id
                                       ) AS name_row_number,
                                   COUNT(*) OVER (
                                       PARTITION BY sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code
                                       ) AS code_count,
                                   COUNT(*) OVER (
                                       PARTITION BY sse_name, rb_name
                                       ) AS name_count
                            FROM generated_staged_locations),
     filtered_locations AS (SELECT *, ((code_count > 1) OR (name_count > 1)) AS is_duplicate
                            FROM ranked_and_counted
                            WHERE code_row_number = 1
                              AND name_row_number = 1)

INSERT
INTO location.locations (sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code, sd_name,
                         sgg_name, emd_name, road_name, building_name, location_5179, location_4326, duplicate_flag)
SELECT sd_code,
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
       location_4326,
       is_duplicate
FROM filtered_locations;

DROP TABLE staged_locations;

-- +goose Down
-- noinspection SqlWithoutWhere
DELETE
FROM location.locations;