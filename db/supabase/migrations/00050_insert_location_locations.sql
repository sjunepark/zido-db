-- +goose Up
CREATE TEMP TABLE staging_locations AS
SELECT *
FROM location.locations WITH NO DATA;

INSERT INTO staging_locations (sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code, sd_name,
                               sgg_name, emd_name, road_name, building_name, location_5179, location_4326)
SELECT LEFT(sgg_code_raw, 2),
       RIGHT(sgg_code_raw, 3),
       LEFT(RIGHT(bjd_code, 6), 3),
       RIGHT(road_code_raw, 7),
       building_main_code,
       NULLIF(NULLIF(building_sub_code, '0'), ''),
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

WITH ranked_and_counted AS (SELECT *,
                                   ROW_NUMBER() OVER (
                                       PARTITION BY sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code, sd_name, emd_name, road_name
                                       ORDER BY (building_name IS NOT NULL) DESC, id
                                       ) AS rn,
                                   COUNT(*) OVER (
                                       PARTITION BY sd_code, sgg_code, emd_code, road_code, building_main_code, building_sub_code, sd_name, emd_name, road_name
                                       ) AS cnt
                            FROM staging_locations),
     filtered_locations AS (SELECT *, (cnt > 1) AS is_duplicate
                            FROM ranked_and_counted
                            WHERE rn = 1)

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


-- +goose Down
-- noinspection SqlWithoutWhere
DELETE
FROM location.locations;