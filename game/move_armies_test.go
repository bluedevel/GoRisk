package game

import (
	"testing"
)

func TestGame_MoveArmies(t *testing.T) {
	t.Run("same occupants", func(t *testing.T) {
		player := Player{name: PlayerName("Player1")}

		nationA := Nation{name: NationName("NationA"), occupant: &player, armies: ArmyCount(10)}
		nationB := Nation{name: NationName("NationB"), occupant: &player, armies: ArmyCount(5)}

		game := Game{
			activePlayer: &player,
			world: &World{
				nations: []*Nation{&nationA, &nationB},
			},
		}

		err := game.MoveArmies("NationA", "NationB", ArmyCount(5))

		if err != nil {
			t.Errorf("expected no error but received %s", err)
		}

		if nationA.armies != 5 {
			t.Errorf("expected NationA armies to be 5 but were %d", nationA.armies)
		}

		if nationB.armies != 10 {
			t.Errorf("expected NationB armies to increase to 10 but was %d", nationB.armies)
		}
	})
}
