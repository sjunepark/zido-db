-- +goose Up
CREATE MATERIALIZED VIEW location.address_summary AS
SELECT DISTINCT sd_sgg_em_name,
                sd_sgg_em_name || ' ' || road_building_name as address
FROM location.locations
WITH DATA;

CREATE INDEX ON location.address_summary (sd_sgg_em_name);
CREATE UNIQUE INDEX ON location.address_summary (address);

-- +goose Down
-- +goose StatementBegin
DROP MATERIALIZED VIEW location.address_summary;
-- +goose StatementEnd