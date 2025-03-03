-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS player_round_status (
    id INT AUTO_INCREMENT PRIMARY KEY,
    player_id INT NOT NULL,
    round_id INT NOT NULL,
    goals INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player_id) REFERENCES player(player_id) ON DELETE CASCADE,
    FOREIGN KEY (round_id) REFERENCES round(round_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE player_round_status;
-- +goose StatementEnd
