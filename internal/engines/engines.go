package engines

import "github.com/cjack0416/rivals-picker/internal/model"

type Engine[T model.EngineQueryRow] interface {
	Run(heroPicks []model.Hero) []model.Hero
	addPickReason(rows T)
}

type AbilityCounterEngine struct {}

func (ac AbilityCounterEngine) Run(heroPicks []model.Hero) []model.Hero {
	return []model.Hero{}
}

func (ac AbilityCounterEngine) addPickReason(rows model.AbilityCounterRow) {

}