package tools

import (
	"context"
	"errors"

	m "github.com/cjack0416/rivals-picker/internal/model"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn
var allHeroes map[string]m.Hero

func InitializeDatabaseConn(dbURI string) (*pgx.Conn, error) {
	log.Info("Connecting to database")

	var connErr error
	conn, connErr = pgx.Connect(context.Background(), dbURI)
	if connErr != nil {
		log.Fatalf("Error connecting to database: %s", connErr)
		return nil, connErr
	}
	log.Info("Successfully connected to database")

	allHeroes = make(map[string]m.Hero)
	queryErr := SelectAllHeroesQuery()
	if queryErr != nil {
		log.Fatalf("Error selecting all heroes: %s", queryErr)
		return nil, queryErr
	}
	return conn, nil
}

func GetAllHeroes() map[string]m.Hero {
	return allHeroes
}

func SelectAllHeroesQuery() error {
	log.Info("Selecting heroes from heroes table")
	query := `SELECT heroes.hero_id, heroes.hero_name, roles.role_name FROM heroes LEFT JOIN roles ON heroes.role_id = roles.role_id`
	if conn == nil {
		return errors.New("no database connection to run query")
	}
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Errorf("Error querying heroes from heroes table with error=%s", err)
		return err
	}
	log.Info("Successfully made query from heroes table")

	for rows.Next() {
		var hero = m.Hero{}
		err := rows.Scan(&hero.HeroId, &hero.HeroName, &hero.HeroRole)
		if err != nil {
			log.Error("Error scanning row: ", err)
			return err
		}
		allHeroes[hero.HeroName] = hero
	}
	return nil
}