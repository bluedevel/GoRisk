package game

import (
	"testing"
)

func TestFightBattle(t *testing.T) {
	t.Run("with same loss on both sides", func(t *testing.T) {
		attackNation := Nation{
			armies: ArmyCount(5),
		}
		defendNation := Nation{
			armies: ArmyCount(4),
		}

		attackPoints := []int{5, 5, 3}
		defendPoints := []int{6, 2}

		attackArmies := ArmyCount(3)

		fightBattle(&attackNation, &defendNation, attackPoints, defendPoints, &attackArmies)

		if attackArmies != 2 {
			t.Errorf("attacking armies expected to be 2 but was %d", attackArmies)
		}

		if attackNation.armies != 4 {
			t.Errorf("attacking nations armies were expected to be 4 but were %d", attackNation.armies)
		}

		if defendNation.armies != 3 {
			t.Errorf("defending nations armies were expected to be 3 but were %d", defendNation.armies)
		}
	})

	t.Run("with more loss on attacking side", func(t *testing.T) {
		attackNation := Nation{
			armies: ArmyCount(5),
		}
		defendNation := Nation{
			armies: ArmyCount(4),
		}

		attackPoints := []int{4, 3}
		defendPoints := []int{5, 4, 3}

		attackArmies := ArmyCount(2)

		fightBattle(&attackNation, &defendNation, attackPoints, defendPoints, &attackArmies)

		if attackArmies != 0 {
			t.Errorf("attacking armies expected to be 0 but was %d", attackArmies)
		}

		if attackNation.armies != 3 {
			t.Errorf("attacking nations armies were expected to be 3 but were %d", attackNation.armies)
		}

		if defendNation.armies != 4 {
			t.Errorf("defending nations armies were expected to be 4 but were %d", defendNation.armies)
		}
	})
}

func TestOccupyNationIfDefeated(t *testing.T) {
	t.Run("with armies left", func(t *testing.T) {
		playerA := Player{}
		playerB := Player{}

		attackNation := Nation{
			armies:   ArmyCount(5),
			occupant: &playerA,
		}
		defendNation := Nation{
			armies:   ArmyCount(4),
			occupant: &playerB,
		}

		occupyNationIfDefeated(&attackNation, &defendNation, ArmyCount(3))

		if attackNation.armies != 5 {
			t.Errorf("attacking nations army count must not change. changed to: %d", attackNation.armies)
		}

		if attackNation.occupant != &playerA {
			t.Errorf("attacking nations player must not change")
		}

		if defendNation.armies != 4 {
			t.Errorf("defending nations army count must not change. changed to: %d", defendNation.armies)
		}

		if defendNation.occupant != &playerB {
			t.Errorf("defending nations player must not change")
		}
	})

	t.Run("with no armies left", func(t *testing.T) {
		playerA := Player{}
		playerB := Player{}

		attackNation := Nation{
			armies:   ArmyCount(2),
			occupant: &playerA,
		}
		defendNation := Nation{
			armies:   ArmyCount(0),
			occupant: &playerB,
		}

		occupyNationIfDefeated(&attackNation, &defendNation, ArmyCount(2))

		if attackNation.armies != 0 {
			t.Errorf("attacking nations armies expected to be 0 but were %d", attackNation.armies)
		}

		if attackNation.occupant != &playerA {
			t.Errorf("occupant of attackking nation must not change")
		}

		if defendNation.armies != 2 {
			t.Errorf("defending nations army count expected to be 2 but was %d", defendNation.armies)
		}

		if defendNation.occupant != &playerA {
			t.Errorf("defending nations occupant didn't change")
		}
	})
}

func TestDrawBattlePoints(t *testing.T) {
	t.Run("with 0 armies", func(t *testing.T) {
		if len(drawBattlePoints(0)) != 0 {
			t.Fail()
		}
	})

	t.Run("with 1 armies", func(t *testing.T) {
		if len(drawBattlePoints(1)) != 1 {
			t.Fail()
		}
	})

	t.Run("with 2 armies", func(t *testing.T) {
		if len(drawBattlePoints(2)) != 2 {
			t.Fail()
		}
	})

	t.Run("with 3 armies", func(t *testing.T) {
		if len(drawBattlePoints(3)) != 3 {
			t.Fail()
		}
	})
}
