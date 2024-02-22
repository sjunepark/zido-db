-- +goose Up
INSERT INTO location.code_names (sd_code, sgg_code, emd_code, sd_name, sgg_name, emd_name)
SELECT SUBSTRING(hjd_code FROM 1 FOR 2) AS sd_code,
       SUBSTRING(hjd_code FROM 3 FOR 3) AS sgg_code,
       SUBSTRING(hjd_code FROM 6 FOR 3) AS emd_code,
       sd_name,
       sgg_name,
       emd_name
FROM raw.code_names r
WHERE SUBSTRING(hjd_code FROM 1 FOR 2) != '00'
  AND SUBSTRING(hjd_code FROM 3 FOR 3) != '000'
  AND SUBSTRING(hjd_code FROM 6 FOR 3) != '000'
  AND create_date IS NOT NULL
  AND expire_date IS NULL;


-- +goose Down
-- +goose StatementBegin
-- noinspection SqlWithoutWhere
DELETE
FROM location.code_names;
DROP SCHEMA location;
-- +goose StatementEnd