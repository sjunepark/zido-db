-- +goose Up

CREATE MATERIALIZED VIEW location.locations_summary AS
SELECT l.id                                            AS location_id,
       l.sd_sgg_em_name || ' ' || l.road_building_name AS address,
       gis.ST_X(l.location_4326)                       AS longitude,
       gis.ST_Y(l.location_4326)                       AS latitude
FROM location.locations l
WITH DATA;

CREATE UNIQUE INDEX ON location.locations_summary (location_id);
CREATE UNIQUE INDEX ON location.locations_summary (address);


-- +goose Down
-- +goose StatementBegin
DROP MATERIALIZED VIEW location.locations_summary;
-- +goose StatementEnd