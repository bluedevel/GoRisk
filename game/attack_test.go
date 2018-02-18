package game

import "testing"

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

	t.Run("with winning defender", func(t *testing.T) {

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
