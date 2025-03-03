-- +goose Up
-- +goose StatementBegin
INSERT INTO podium (season_id, player_id, classification) VALUES
(1, 1, 1),
(1, 2, 2),
(1, 3, 3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM podium WHERE season_id = 1;
-- +goose StatementEnd
