-- +goose Up

CREATE SCHEMA raw;

CREATE TABLE raw.locations_summary
(
    sgg_code_raw          CHAR(5)     NOT NULL CHECK (sgg_code_raw ~ '^[0-9]{5}$'),    -- sd_code(2) + sgg_code(3)
    entrance_code         CHAR(10) CHECK (entrance_code ~ '^[0-9]{1,10}\s+$'),
    bjd_code              CHAR(10)    NOT NULL CHECK (bjd_code ~ '^[0-9]{10}$'),       -- sgg_code(5) + emd_code(3) + 00
    sd_name               VARCHAR(40) NOT NULL CHECK ( sd_name = TRIM(sd_name) ),
    sgg_name              VARCHAR(40) CHECK ( sgg_name = TRIM(sgg_name) ),
    emd_name              VARCHAR(40) NOT NULL CHECK ( emd_name = TRIM(emd_name) ),
    road_code_raw         CHAR(12)    NOT NULL CHECK (road_code_raw ~ '^[0-9]{12}$'),  -- sgg_code(5) + road_code(7)
    road_name             VARCHAR(40) NOT NULL CHECK ( road_name = TRIM(road_name) ),
    underground_flag      CHAR(1)     NOT NULL CHECK (underground_flag ~ '^[012]$'),
    building_main_code    VARCHAR(5)  NOT NULL CHECK (building_main_code ~ '^[0-9]{1,5}$'),
    building_sub_code     VARCHAR(5) CHECK (building_sub_code ~ '^[0-9]{1,5}$'),
    building_name         VARCHAR(40) CHECK ( building_name = TRIM(building_name) ),
    postal_code           CHAR(5)     NOT NULL CHECK (postal_code ~ '^[0-9]{5}$'),
    building_use_category VARCHAR(100) CHECK ( building_use_category = TRIM(building_use_category) ),
    building_group_flag   CHAR(1)     NOT NULL CHECK (building_group_flag ~ '^[01]$'), -- 0: 단독건물, 1: 건물그군
    jurisdiction_hjd      VARCHAR(20),
    x                     REAL,
    y                     REAL,
    CONSTRAINT pk_primary_key PRIMARY KEY (bjd_code, road_code_raw, underground_flag, building_main_code,
                                           building_sub_code)
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE raw.locations_summary;

DROP SCHEMA raw;
-- +goose StatementEnd