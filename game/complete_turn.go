package game

func (game *Game) completeTurn() {
	game.activePlayer = game.getNextPlayer()
	game.turn++

	game.calculateBonusArmies()
}

func (game *Game) calculateBonusArmies() {
	for _, nation := range game.world.nations {
		if nation.occupant == game.activePlayer {
			game.activePlayer.freeArmies++
		}
	}

	for _, region := range game.world.regions {
		ok, player := region.SingleOccupant();
		if ok && player == game.activePlayer {
			game.activePlayer.freeArmies += region.bonusArmies
		}
	}
}
