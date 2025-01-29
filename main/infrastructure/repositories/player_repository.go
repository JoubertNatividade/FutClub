package repositories

import (
	"database/sql"
	"github.com/JoubertNatividade/FutClub/main/domain/entities"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(
	db *sql.DB,
) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(player *entities.Player) error {
	log.Info("starting create repository...")
	log.Infof("--player received: %+v", player)
	_, err := r.db.Query("INSERT INTO player(name, last_name, position) VALUES(?,?,?)", player.Name, player.LastName, player.Position)
	if err != nil {
		log.Errorf("error on create player: %s", err)
		return err
	}
	log.Infof("player %s created successfully", player.Name)
	return nil
}
func (r *PlayerRepository) List() ([]entities.Player, error) {
	log.Info("Starting list repository...")

	query := "SELECT player_id, name, last_name, position, created_at FROM player"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Errorf("Error listing players: %s", err)
		return nil, err
	}
	defer rows.Close()

	var players []entities.Player

	for rows.Next() {
		var player entities.Player
		err = rows.Scan(&player.PlayerID, &player.Name, &player.LastName, &player.Position, &player.CreatedAt)
		if err != nil {
			log.Errorf("Error scanning player: %s", err)
			continue
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("Error iterating over players: %s", err)
		return nil, err
	}

	log.Infof("Quantity of players found: %d", len(players))

	return players, nil
}
