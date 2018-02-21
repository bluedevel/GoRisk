package game

import (
	"errors"
	"fmt"
)

type Game struct {
	world        *World
	turn         int
	activePlayer *Player
	players      []*Player
}

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

type PlayerName string
type Player struct {
	name          PlayerName
	conquestCards map[ConquestCardType]ConquestCardCount
}

type BattleResult struct {
	attackPoints []int
	defendPoints []int
}

func (game *Game) completeTurn() {
	game.activePlayer = game.getNextPlayer()
	game.turn++
}

// Returns the next player in line or the first of none is set yet.
func (game Game) getNextPlayer() *Player {
	if game.activePlayer == game.players[len(game.players)-1] {
		return game.players[0]
	}

	for i, player := range game.players {
		if game.activePlayer == player {
			return game.players[i+1]
		}
	}

	return game.players[0]
}

func (game Game) getNation(name NationName) (*Nation, error) {
	for _, nation := range game.world.nations {
		if nation.name == name {
			return nation, nil
		}
	}

	return nil, fmt.Errorf("no nation by name $s", name)
}

// Trades a certain conquest card type for troops and places them as specified.
//
// Returns error if:
//	1) There is no player by this name
// 	2) There are not enough conquest cards of this type
//	3) There is no nation by a supplied name
//	4) An army count is smaller than one
func (game *Game) TradeConquestCards(playerName PlayerName, cardType ConquestCardType, positions map[NationName]ArmyCount) error {
	return errors.New("not yet implemented")
}

// Calculates the amount of armies received if the conquest card type was traded in.
//
// Returns error if:
//	1) There is no player by this name
func (game Game) ConquestCardArmies(playerName PlayerName, cardType ConquestCardType) (ArmyCount, error) {
	return 0, errors.New("not yet implemented")
}
