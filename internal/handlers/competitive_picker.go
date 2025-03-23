package handlers

import (
	"strings"

	"github.com/cjack0416/rivals-picker/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CompetitivePicker(c *fiber.Ctx) error {
	var params = api.PickHeroesCompetitiveParams{}

	m := c.Queries()
	params.TeamHeroes = strings.Split(m["team_heroes"], ",")
	params.EnemyHeroes = strings.Split(m["enemy_heroes"], ",")

	log.Info("Your teams heroes are: ", params.TeamHeroes)
	log.Info("Your first team hero is: ", params.TeamHeroes[0])
	log.Info("The enemy heroes are: ", params.EnemyHeroes)

	return c.Status(200).JSON("{}")
}