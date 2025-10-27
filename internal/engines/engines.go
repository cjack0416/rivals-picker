package engines

import (
	"fmt"

	"github.com/cjack0416/rivals-picker/internal/model"
	"github.com/cjack0416/rivals-picker/internal/tools"
)

type Engine[T model.EngineQueryResult] interface {
	Run(heroPicks []model.Hero, args [][]int) ([]model.Hero, error)
	addPickReason(rows T)
}

type AbilityCounterEngine struct {}

func (ac AbilityCounterEngine) Run(heroPicks []model.Hero, args [][]int) ([]model.Hero, error) {
	var heroIdList = make([]int, 0, len(heroPicks))

	for _, hero := range heroPicks {
		heroIdList = append(heroIdList, hero.HeroId)
	}
	acRows, err := tools.GetHeroAbilityCounters(heroIdList, args[0])
	if err != nil {
		return nil, err
	}

	for _, row := range acRows {
		fmt.Println(row)
	}


	return []model.Hero{}, nil
}

func (ac AbilityCounterEngine) addPickReason(rows model.AbilityCounterResult) {

}