package common

import (
	"testing"
)

func TestRandomFlat(t *testing.T) {
	puns := P.Random("Flat")
	expectedPuns := P["Flat"]

	found := false
	for _, pun := range expectedPuns {
		if puns == pun {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Random() returned unexpected pun for Flat: got %v, expected one of %v", puns, expectedPuns)
	}
}

func TestRandomRound(t *testing.T) {
	puns := P.Random("Round")
	expectedPuns := P["Round"]

	found := false
	for _, pun := range expectedPuns {
		if puns == pun {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Random() returned unexpected pun for Round: got %v, expected one of %v", puns, expectedPuns)
	}
}

func TestRandomWaiting(t *testing.T) {
	puns := P.Random("Waiting")
	expectedPuns := P["Waiting"]

	found := false
	for _, pun := range expectedPuns {
		if puns == pun {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Random() returned unexpected pun for Waiting: got %v, expected one of %v", puns, expectedPuns)
	}
}

func TestRandomInvalidCategory(t *testing.T) {
	puns := P.Random("NonExistent")
	expected := "Invalid category!"

	if puns != expected {
		t.Errorf("Random() returned wrong output for invalid category: got %v, want %v", puns, expected)
	}
}
