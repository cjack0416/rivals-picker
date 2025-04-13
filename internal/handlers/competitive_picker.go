package handlers

import (
	"context"
	"strings"

	"github.com/cjack0416/rivals-picker/api"
	"github.com/cjack0416/rivals-picker/internal/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Hero struct {
	HeroId int
	HeroName string
	HeroRole string
	HeroScore int
}

const duelist = "DUELIST"
const vanguard = "VANGUARD"
const strategist = "STRATEGIST"

func CompetitivePicker(c *fiber.Ctx) error {
	conn := tools.GetDatabaseConn()

	var params = api.PickHeroesCompetitiveParams{}

	m := c.Queries()
	params.TeamHeroes = strings.Split(m["team_heroes"], ",")
	params.EnemyHeroes = strings.Split(m["enemy_heroes"], ",")

	log.Info("Making query in heroes table")
	rows, err := conn.Query(context.Background(), "SELECT heroes.hero_id, heroes.hero_name, roles.role_name FROM heroes LEFT JOIN roles ON heroes.role_id = roles.role_id")
	if err != nil {
		log.Errorf("Error querying heroes from heroes table with query params team_heroes=%s, and error=%s", params.TeamHeroes, err)
		return c.Status(500).JSON("{}")
	}
	log.Info("Successfully made query from heroes table")

	var heroPicks map[string]Hero = make(map[string]Hero)
	for rows.Next() {
		var hero = Hero{}
		err := rows.Scan(&hero.HeroId, &hero.HeroName, &hero.HeroRole)
		if err != nil {
			log.Error("Error scanning row: ", err)
			return c.Status(500).JSON("{}")
		}
		heroPicks[hero.HeroName] = hero
	}
	log.Info("Result from heroes table query: ", heroPicks)

	heroPicks = filterHeroRoles(heroPicks, params.TeamHeroes)
	heroPicks = filterTeamHeroes(heroPicks, params.TeamHeroes)

	return c.Status(200).JSON(heroPicks)
}

func filterHeroRoles(heroMap map[string]Hero, teamHeroes []string) map[string]Hero {
	var filterRoles map[string]int = map[string]int{duelist: 0, vanguard: 0, strategist: 0}

	for _, hero := range teamHeroes {
		heroRole := heroMap[hero].HeroRole
		if heroRole == duelist {
			filterRoles[duelist]++
		} else if heroRole == vanguard {
			filterRoles[vanguard]++
		} else {
			filterRoles[strategist]++
		}
	}

	var filteredHeroMap map[string]Hero = make(map[string]Hero)

	for _, hero := range heroMap {
		if filterRoles[hero.HeroRole] < 3 {
			hero.HeroScore += 3 - filterRoles[hero.HeroRole]
			filteredHeroMap[hero.HeroName] = hero
		}
	}

	return filteredHeroMap
}

func filterTeamHeroes(heroMap map[string]Hero, teamHeroes []string) map[string]Hero {
	var filteredHeroMap map[string]Hero = make(map[string]Hero)
	var teamHeroMap map[string]string = make(map[string]string)

	for _, heroName := range teamHeroes {
		teamHeroMap[heroName] = heroName
	}

	for _, hero := range heroMap {
		var _, ok = teamHeroMap[hero.HeroName]
		if !ok {
			filteredHeroMap[hero.HeroName] = hero
		}
	}

	return filteredHeroMap
}