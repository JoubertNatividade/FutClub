-- +goose Up
-- +goose StatementBegin
INSERT INTO player (name, last_name, position_id, avatar_url) VALUES
('Joubert', 'Natividade', 1, 'https://example.com/avatars/carlos.jpg'),
('Ramon', 'Bonfim', 2, 'https://example.com/avatars/joao.jpg'),
('Diego', 'Pereira', 3, 'https://example.com/avatars/lucas.jpg'),
('Jair', 'Souza', 4, 'https://example.com/avatars/ricardo.jpg');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM player WHERE first_name IN ('Joubert', 'Ramon', 'Diego', 'Jair');
-- +goose StatementEnd
