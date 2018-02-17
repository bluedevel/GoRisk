package game

import "testing"



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
