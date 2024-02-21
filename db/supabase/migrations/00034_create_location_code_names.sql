-- +goose Up
CREATE SCHEMA location;

CREATE TABLE location.code_names
(
    sd_code  CHAR(2)     NOT NULL CHECK (sd_code ~ '^[0-9]{2}$'),
    sgg_code CHAR(3)     NOT NULL CHECK (sgg_code ~ '^[0-9]{3}$'),
    emd_code CHAR(3)     NOT NULL CHECK (emd_code ~ '^[0-9]{3}$'),
    sd_name  VARCHAR(40) NOT NULL CHECK (sd_name = TRIM(sd_name)),
    sgg_name VARCHAR(40) CHECK (sgg_name = TRIM(sgg_name)),
    emd_name VARCHAR(40) NOT NULL CHECK (emd_name = TRIM(emd_name)),
    PRIMARY KEY (sd_code, sgg_code, emd_code)
);


-- +goose Down
-- +goose StatementBegin
DROP TABLE location.code_names;
DROP SCHEMA location;
-- +goose StatementEnd