package api

type PickHeroesCompetitiveParams struct {
	TeamHeroes []string
	EnemyHeroes []string
}

type PickHeroesCompetitiveResponse struct {
	HeroPick string `json:"heroPick"`
	PickReasons []string `json:"pickReasons"`
}