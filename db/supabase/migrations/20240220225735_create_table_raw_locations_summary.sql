-- +goose Up

CREATE SCHEMA raw;

CREATE TABLE raw.locations_summary
(
    sgg_number            CHAR(5)     NOT NULL CHECK (sgg_number ~ '^[0-9]{5}$'),
    entrance_number       CHAR(10) CHECK (entrance_number ~ '^[0-9]{1,10}\s+$'),
    bjd_number            CHAR(10)    NOT NULL CHECK (bjd_number ~ '^[0-9]{10}$'),
    sd_name               VARCHAR(40) NOT NULL CHECK ( sd_name = TRIM(sd_name) ),
    sgg_name              VARCHAR(40) CHECK ( sgg_name = TRIM(sgg_name) ),
    emd_name              VARCHAR(40) NOT NULL CHECK ( emd_name = TRIM(emd_name) ),
    road_number           CHAR(12)    NOT NULL CHECK (road_number ~ '^[0-9]{12}$'),
    road_name             VARCHAR(40) NOT NULL CHECK ( road_name = TRIM(road_name) ),
    underground_flag      CHAR(1)     NOT NULL CHECK (underground_flag ~ '^[012]$'),
    building_main_number  VARCHAR(5)  NOT NULL CHECK (building_main_number ~ '^[0-9]{1,5}$'),
    building_sub_number   VARCHAR(5) CHECK (building_sub_number ~ '^[0-9]{1,5}$'),
    building_name         VARCHAR(40) CHECK ( building_name = TRIM(building_name) ),
    postal_number         CHAR(5)     NOT NULL CHECK (postal_number ~ '^[0-9]{5}$'),
    building_use_category VARCHAR(100) CHECK ( building_use_category = TRIM(building_use_category) ),
    building_group_flag   CHAR(1)     NOT NULL CHECK (building_group_flag ~ '^[01]$'),
    jurisdiction_hjd      VARCHAR(8)  NOT NULL,
    x                     REAL,
    y                     REAL,
    CONSTRAINT pk_primary_key PRIMARY KEY (bjd_number, road_number, underground_flag, building_main_number,
                                           building_sub_number)
);

-- +goose Down
-- +goose StatementBegin
DO
$$
    BEGIN
        IF (SELECT COUNT(*) FROM raw.locations_summary) = 0 THEN
            DROP TABLE raw.locations_summary;
        END IF;
    END;
$$;

DROP SCHEMA raw;
-- +goose StatementEnd