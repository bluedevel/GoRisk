package game

import (
	"fmt"
)

// Moves armies from one nation to another
//
// Return error if:
// 	1) There is no player by this name
// 	2) There are no nations by this names
// 	3) The army count is smaller than one
//	4) The nations are not occupied by the player
//	5) There is no connecting occupied region between the nations
func (game *Game) MoveArmies(from NationName, to NationName, amount ArmyCount) error {

	fromNation :=  &Nation{}
	toNation := &Nation{}

	if err := validateMoveArmiesInput(&fromNation, &toNation, *game, from, to, amount); err != nil {
		return err
	}

	fromNation.armies -= amount
	toNation.armies += amount

	return nil
}

func validateMoveArmiesInput(fromNation **Nation, toNation **Nation, game Game, from NationName, to NationName, amount ArmyCount) error {

	var nation *Nation
	var err error

	if nation, err = game.getNation(from); err != nil {
		return err
	}
	*fromNation = nation

	if (*fromNation).armies < amount {
		return fmt.Errorf("unable to move more armies then present")
	}

	if err := game.validatePlayer((*fromNation).occupant.name); err != nil {
		return err
	}

	if nation, err = game.getNation(to); err != nil {
		return err
	}
	*toNation = nation

	if (*fromNation).occupant != (*toNation).occupant {
		return fmt.Errorf("unable to move armies betweeen nation of different occupants")
	}

	// implement region path finding here

	return nil
}
