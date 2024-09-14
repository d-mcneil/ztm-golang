//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import (
	"fmt"
	"testing"
)

func newPlayer(playerNumber int) Player {
	return Player{
		name:      Name(fmt.Sprintf("Player%d", playerNumber)),
		health:    50,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 100,
	}
}

func TestIncreaseHealth(t *testing.T) {
	player := newPlayer(1)
	player2 := newPlayer(2)

	player.IncreaseHealth(60)
	if player.health > player.maxHealth {
		t.Errorf("Health should not go above maximum health. Current health: %d -> Max health: %d", player.health, player.maxHealth)
	}

	player2.IncreaseHealth(25)
	if player2.health != 75 {
		t.Errorf("Health should be 75. Current health: %d", player2.health)
	}
}

func TestDecreaseHealth(t *testing.T) {
	player := newPlayer(1)
	player2 := newPlayer(2)

	player.DecreaseHealth(60)
	if player.health > player.maxHealth {
		t.Errorf("Health should not go below 0. Current health: %d", player.health)
	}
	if player.health < 0 {
		t.Errorf("Health should not go below 0. Current health: %d", player.health)
	}

	player2.DecreaseHealth(25)
	if player2.health != 25 {
		t.Errorf("Health should be 25. Current health: %d", player2.health)
	}
}

func TestIncreaseEnergy(t *testing.T) {
	player := newPlayer(1)
	player2 := newPlayer(2)

	player.IncreaseEnergy(60)
	if player.energy > player.maxEnergy {
		t.Errorf("Energy should not go above maximum energy. Current energy: %d -> Max energy: %d", player.energy, player.maxEnergy)
	}

	player2.IncreaseEnergy(25)
	if player2.energy != 75 {
		t.Errorf("Energy should be 75. Current energy: %d", player2.energy)
	}
}

func TestDecreaseEnergy(t *testing.T) {
	player := newPlayer(1)
	player2 := newPlayer(2)

	player.DecreaseEnergy(60)
	if player.energy > player.maxEnergy {
		t.Errorf("Energy should not go below 0. Current energy: %d", player.energy)
	}
	if player.energy < 0 {
		t.Errorf("Energy should not go below 0. Current energy: %d", player.energy)
	}

	player2.DecreaseEnergy(25)
	if player2.energy != 25 {
		t.Errorf("Energy should be 25. Current energy: %d", player2.energy)
	}
}
