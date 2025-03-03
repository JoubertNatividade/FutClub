-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS podium (
    podium_id INT AUTO_INCREMENT PRIMARY KEY,
    season_id INT NOT NULL,
    player_id INT NOT NULL,
    classification INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (season_id) REFERENCES season(season_id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES player(player_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE podium;
-- +goose StatementEnd
