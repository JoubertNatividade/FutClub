-- +goose Up
-- +goose StatementBegin
INSERT INTO player_position (name) VALUES
('Atacante'),
('Meio-Campo'),
('Zagueiro'),
('Goleiro');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM player_position WHERE name IN ('Atacante', 'Meio-Campo', 'Zagueiro', 'Goleiro');

-- +goose StatementEnd
