-- +goose Up
-- +goose StatementBegin
INSERT INTO round (round_date) VALUES
('2025-03-01'),
('2025-03-08'),
('2025-03-15');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM round WHERE round_date IN ('2025-03-01', '2025-03-08', '2025-03-15');
-- +goose StatementEnd
