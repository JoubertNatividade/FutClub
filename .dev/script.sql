
CREATE TABLE player (
	player_id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	position VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE game (
	game_id INT AUTO_INCREMENT PRIMARY KEY,
	game_date DATE NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE player_game_stats (
	stat_id INT AUTO_INCREMENT PRIMARY KEY,
	player_id INT NOT NULL,
	game_id INT NOT NULL,
	goals INT DEFAULT 0,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (player_id) REFERENCES player(player_id) ON DELETE CASCADE,
	FOREIGN KEY (game_id) REFERENCES game(game_id) ON DELETE CASCADE
);
