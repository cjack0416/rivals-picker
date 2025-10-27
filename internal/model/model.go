package model

type Hero struct {
	HeroId int
	HeroName string
	HeroRole string
	HeroScore int
	PickReasons []string
}

type EngineQueryResult interface {
	AbilityCounterResult
}

type AbilityCounterResult struct {
	HeroId int
	AbilityName string
	EnemyId int
	EnemyAbilityName string
	Weight int
}