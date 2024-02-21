-- +goose Up


-- +goose Down
-- +goose StatementBegin
-- noinspection SqlWithoutWhere
DELETE
FROM location.code_nmaes
-- +goose StatementEnd