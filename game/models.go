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

type BattleResult string

const CONQUERED BattleResult = "CONQUERED"
const DEFENDED BattleResult = "DEFENDED"

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

// Moves armies from one nation to another
//
// Return error if:
// 	1) There is no player by this name
// 	2) There are no nations by this names
// 	3) The army count is smaller than one
//	4) The nations are not occupied by the player
//	5) There is no connecting occupied region between the nations
func (game *Game) MoveArmies(playerName PlayerName, from NationName, to NationName, amount ArmyCount) error {
	if err := game.validatePlayer(playerName); err != nil {
		return err
	}

	return errors.New("not yet implemented")
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

func (game Game) getNation(name NationName) (*Nation, error) {
	for _, nation := range game.world.nations {
		if nation.name == name {
			return nation, nil
		}
	}

	return nil, fmt.Errorf("no nation by name $s", name)
}

func (game Game) validatePlayer(playerName PlayerName) error {
	if game.activePlayer.name != playerName {
		return fmt.Errorf("player %s is not active", playerName)
	}

	return nil
}

func (nation Nation) validateArmyCount(count ArmyCount) error {
	if nation.armies < count {
		return fmt.Errorf("not enough armies")
	}
	return nil
}

func validateArmyCountForBattle(count ArmyCount) error {
	if count < 1 || count > 3 {
		return fmt.Errorf("invalid army count for battle: %d", count)
	}
	return nil
}
