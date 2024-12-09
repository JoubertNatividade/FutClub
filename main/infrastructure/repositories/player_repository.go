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
	_, err := r.db.Query("INSERT INTO player(name, last_name, position) VALUES(?,?,?)", player.Name, player.LastName, player.Position)
	if err != nil {
		return err
	}
	return nil
}
