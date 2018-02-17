package game

import (
	"math/rand"
	"sort"
)

// Attacks a neighbouring nation. Removes defeated armies and sets the occupant
// of the attacked nation if no defending troops are left.
//
// Returns the result of the battle.
// An error is returned if:
// 	1) There are no nations by the supplied names
// 	2) There are not enough armies in the battling nations
// 	3) The army counts are smaller than one or greater then three
func (game *Game) Attack(attacker NationName, defender NationName, attackArmies ArmyCount, defendArmies ArmyCount) (BattleResult, error) {
	var result BattleResult

	var attackingNation *Nation
	var defendingNation *Nation
	if err := validateAttackInput(attackingNation, defendingNation, *game, attacker, defender, attackArmies, defendArmies); err != nil {
		return result, err
	}

	attackPoints := drawBattlePoints(attackArmies)
	defendPoints := drawBattlePoints(defendArmies)

	sort.Sort(sort.Reverse(sort.IntSlice(attackPoints)))
	sort.Sort(sort.Reverse(sort.IntSlice(defendPoints)))

	fightBattle(attackingNation, defendingNation, attackPoints, defendPoints, &attackArmies)
	occupyNationIfDefeated(attackingNation, defendingNation, attackArmies)

	result.attackPoints = attackPoints
	result.defendPoints = defendPoints

	return result, nil
}

func fightBattle(attackingNation *Nation, defendingNation *Nation, attackPoints []int, defendPoints []int, attackArmies *ArmyCount) {
	for i := 0; i < min(len(attackPoints), len(defendPoints)); i++ {
		attack := attackPoints[i]
		defend := defendPoints[i]

		if attack > defend {
			defendingNation.armies--
		} else {
			attackingNation.armies--
			(*attackArmies)--
		}
	}
}

func occupyNationIfDefeated(attackingNation *Nation, defendingNation *Nation, attackArmies ArmyCount) {
	if defendingNation.armies <= 0 {
		defendingNation.occupant = attackingNation.occupant
		attackingNation.armies -= attackArmies
		defendingNation.armies = attackArmies
	}
}

func drawBattlePoints(amount ArmyCount) []int {
	res := make([]int, amount)

	for i := 0; i < int(amount); i++ {
		res[i] = rand.Int()
	}

	return res
}

func validateAttackInput(attackingNation *Nation, defendingNation *Nation, game Game, attacker NationName, defender NationName, attackArmies ArmyCount, defendArmies ArmyCount) error {
	var err error
	var nation *Nation

	if err = validateArmyCountForBattle(attackArmies); err != nil {
		return err
	}

	if err = validateArmyCountForBattle(defendArmies); err != nil {
		return err
	}

	if nation, err = game.getNation(attacker); err != nil {
		return err
	}
	*attackingNation = *nation

	if err := game.validatePlayer(attackingNation.occupant.name); err != nil {
		return err
	}

	if nation, err = game.getNation(defender); err != nil {
		return err
	}
	*defendingNation = *nation

	if err = attackingNation.validateArmyCount(attackArmies); err != nil {
		return err
	}

	if err = defendingNation.validateArmyCount(defendArmies); err != nil {
		return err
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
