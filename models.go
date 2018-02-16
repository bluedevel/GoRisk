package main

import "errors"

type World struct {
	nations []*Nation
}

type NationName string
type ArmyCount int
type Nation struct {
	name       NationName
	neighbours []*Nation
	occupant   *Player
	armies     ArmyCount
}

type ConquestCardType string
type ConquestCardCount int

type Player struct {
	conquestCards map[ConquestCardType]ConquestCardCount
}

type BattleResult string

const CONQUERED BattleResult = "CONQUERED"
const DEFENDED BattleResult = "DEFENDED"

/*
 * World logic
 */


// Attacks a neighbouring nation. Removes killed troops and sets the occupant
// of the attacked nation if no defending troops are left.
//
// Returns the result of the battle.
// An error is returned if:
// 	1) There is no neighbouring nation by the supplied name
// 	2) There are not enough armies in the battling nations or the inputs are smaller than one
func (nation *Nation) Attack(defender NationName, attackArmies int, defendArmies int) (BattleResult, error) {
	var result BattleResult
	return result, errors.New("not yet implemented")
}

// Trades a certain conquest card type for troops and places them as specified.
//
// Returns error if:
// 	1) There are not enough conquest cards of this type
//	2) There is no nation by a supplied name
//	3) A army count is smaller than one
func (player *Player) TradeConquestCards(cardType ConquestCardType, positions map[NationName]ArmyCount) error {
	return errors.New("not yet implemented")
}

// Calculates the amount of armies received if the conquest card type was traded in.
func (player Player) ConquestCardArmies(cardType ConquestCardType) ArmyCount {
	return 0
}
