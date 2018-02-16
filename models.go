package main

type World struct {
	nations []*Nation
}

type NationName string
type Nation struct {
	name       NationName
	neighbours []*Nation
	occupant   *Player
	armies     int
}

type ConquestCardName string
type ConquestCardCount int

type Player struct {
	conquestCards map[ConquestCardName]ConquestCardCount
}

type BattleResult string
const CONQUERED BattleResult = "CONQUERED"
const DEFENDED BattleResult = "DEFENDED"

/*
 * World logic
 */

func (world *World) attack(attacker NationName, defender NationName, attackArmies int, defendArmies int) {

}
