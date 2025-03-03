-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS round (
    round_id INT AUTO_INCREMENT PRIMARY KEY,
    round_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE round;
-- +goose StatementEnd
