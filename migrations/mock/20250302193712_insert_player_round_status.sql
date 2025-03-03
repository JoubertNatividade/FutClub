-- +goose Up
-- +goose StatementBegin
INSERT INTO player_round_status (player_id, round_id, goals) VALUES
(1, 1, 2),
(2, 1, 1),
(3, 2, 3),
(4, 2, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM player_round_status WHERE round_id IN (1, 2);
-- +goose StatementEnd
