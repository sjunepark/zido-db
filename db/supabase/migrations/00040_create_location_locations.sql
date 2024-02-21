-- +goose Up
CREATE SCHEMA location;
CREATE SCHEMA gis;
CREATE EXTENSION postgis WITH SCHEMA gis;

CREATE TABLE location.code_name
(
    sd_code  CHAR(2)     NOT NULL CHECK (sd_code ~ '^[0-9]{2}$'),
    sgg_code CHAR(3)     NOT NULL CHECK (sgg_code ~ '^[0-9]{3}$'),
    emd_code CHAR(3)     NOT NULL CHECK (emd_code ~ '^[0-9]{3}$'),
    sd_name  VARCHAR(40) NOT NULL CHECK (sd_name = TRIM(sd_name)),
    sgg_name VARCHAR(40) CHECK (sgg_name = TRIM(sgg_name)),
    emd_name VARCHAR(40) NOT NULL CHECK (emd_name = TRIM(emd_name)),
    PRIMARY KEY (sd_code, sgg_code, emd_code)
);


CREATE TABLE location.locations
(
    id                 SERIAL PRIMARY KEY,
    sd_code            CHAR(2)                   NOT NULL CHECK (sd_code ~ '^[0-9]{2}$'),
    sgg_code           CHAR(3)                   NOT NULL CHECK (sgg_code ~ '^[0-9]{3}$'),
    emd_code           CHAR(3)                   NOT NULL CHECK (emd_code ~ '^[0-9]{3}$'),
    road_code          CHAR(7)                   NOT NULL CHECK (road_code ~ '^[0-9]{7}$'),
    building_main_code VARCHAR(5)                NOT NULL CHECK (building_main_code ~ '^[0-9]{1,5}$'),
    building_sub_code  VARCHAR(5) CHECK (building_sub_code ~ '^[0-9]{1,5}$'),
    sd_name            VARCHAR(40)               NOT NULL CHECK (sd_name = TRIM(sd_name)),
    sgg_name           VARCHAR(40) CHECK (sgg_name = TRIM(sgg_name)),
    emd_name           VARCHAR(40)               NOT NULL CHECK (emd_name = TRIM(emd_name)),
    road_name          VARCHAR(40)               NOT NULL CHECK (road_name = TRIM(road_name)),
    building_name      VARCHAR(40) CHECK (building_name = TRIM(building_name)),
    location_5179      gis.GEOMETRY(POINT, 5179) NOT NULL,
    location_4326      gis.GEOMETRY(POINT, 4326) NOT NULL,
    sd_sgg_em_name     VARCHAR(100) GENERATED ALWAYS AS (sd_name || COALESCE(' ' || sgg_name, '') || CASE
                                                                                                         WHEN RIGHT(emd_name, 1) = '동'
                                                                                                             THEN ''
                                                                                                         ELSE ' ' || emd_name END) STORED,
    road_building_name VARCHAR(100) GENERATED ALWAYS AS (road_name || ' ' || building_main_code || CASE
                                                                                                       WHEN building_sub_code IS NOT NULL
                                                                                                           THEN '-' || building_sub_code
                                                                                                       ELSE '' END) STORED,
    duplicate_flag     BOOLEAN                   NOT NULL DEFAULT FALSE
);

ALTER TABLE location.locations
    ADD CONSTRAINT chk_locations_unique_code_name UNIQUE (sd_code, sgg_code, emd_code, road_code, building_main_code,
                                                          building_sub_code, sd_sgg_em_name, road_building_name);

CREATE INDEX ON location.locations (sd_sgg_em_name);


-- +goose Down
-- +goose StatementBegin
DROP TABLE location.locations;
DROP EXTENSION postgis;
DROP SCHEMA gis;
DROP SCHEMA location;
-- +goose StatementEnd