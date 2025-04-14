package handlers

import (
	"strings"

	"github.com/cjack0416/rivals-picker/api"
	"github.com/cjack0416/rivals-picker/internal/tools"
	m "github.com/cjack0416/rivals-picker/internal/model"
	"github.com/gofiber/fiber/v2"
)

const duelist = "DUELIST"
const vanguard = "VANGUARD"
const strategist = "STRATEGIST"

func CompetitivePicker(c *fiber.Ctx) error {
	var params = api.PickHeroesCompetitiveParams{}

	q := c.Queries()
	params.TeamHeroes = strings.Split(q["team_heroes"], ",")
	params.EnemyHeroes = strings.Split(q["enemy_heroes"], ",")

	heroPicks := tools.GetAllHeroes()
	if len(heroPicks) == 0 {
		resp := api.Error{Code: 500, Message: "All heroes cache is empty."}
		return c.Status(500).JSON(resp)
	}
	heroPicks = filterHeroRoles(heroPicks, params.TeamHeroes)
	heroPicks = filterTeamHeroes(heroPicks, params.TeamHeroes)

	return c.Status(200).JSON(heroPicks)
}

func filterHeroRoles(heroMap map[string]m.Hero, teamHeroes []string) map[string]m.Hero {
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

	var filteredHeroMap map[string]m.Hero = make(map[string]m.Hero)

	for _, hero := range heroMap {
		if filterRoles[hero.HeroRole] < 3 {
			hero.HeroScore += 3 - filterRoles[hero.HeroRole]
			filteredHeroMap[hero.HeroName] = hero
		}
	}

	return filteredHeroMap
}

func filterTeamHeroes(heroMap map[string]m.Hero, teamHeroes []string) map[string]m.Hero {
	var filteredHeroMap map[string]m.Hero = make(map[string]m.Hero)
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