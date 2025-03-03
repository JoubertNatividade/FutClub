-- +goose Up
-- +goose StatementBegin
INSERT INTO season (name) VALUES
('Temporada 1'),
('Temporada 2');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM season WHERE name IN ('Temporada 1', 'Temporada 2');
-- +goose StatementEnd
