package api

type PickHeroesCompetitiveParams struct {
	TeamHeroes []string
	EnemyHeroes []string
}

type PickHeroesCompetitiveResponse struct {
	HeroPick string `json:"heroPick"`
	PickReasons []string `json:"pickReasons"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}