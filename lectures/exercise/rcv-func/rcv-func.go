//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health uint
type Energy uint
type Name string

type Player struct {
	name              Name
	health, maxHealth Health
	energy, maxEnergy Energy
}

func (player *Player) increaseHealth(healthToIncrease Health) {
	maxHealth := player.maxHealth
	totalHealth := player.health + healthToIncrease
	if totalHealth > maxHealth {
		totalHealth = maxHealth
	}
	healthIncrease := totalHealth - player.health
	player.health = totalHealth
	fmt.Println("Health increased by", healthIncrease)
	fmt.Println(*player, "\n")
}

func (player *Player) decreaseHealth(healthToDecrease Health) {
	totalHealth := player.health - healthToDecrease
	if totalHealth > player.health {
		totalHealth = 0
	}
	healthDecrease := player.health - totalHealth
	player.health = totalHealth
	fmt.Println("Health decreased by", healthDecrease)
	fmt.Println(*player, "\n")
}

func (player *Player) increaseEnergy(energyToIncrease Energy) {
	maxEnergy := player.maxEnergy
	totalEnergy := player.energy + energyToIncrease
	if totalEnergy > maxEnergy {
		totalEnergy = maxEnergy
	}
	energyIncrease := totalEnergy - player.energy
	player.energy = totalEnergy
	fmt.Println("Energy increased by", energyIncrease)
	fmt.Println(*player, "\n")
}

func (player *Player) decreaseEnergy(energyToDecrease Energy) {
	totalEnergy := player.energy - energyToDecrease
	if totalEnergy > player.energy {
		totalEnergy = 0
	}
	energyDecrease := player.energy - totalEnergy
	player.energy = totalEnergy
	fmt.Println("Energy decreased by", energyDecrease)
	fmt.Println(*player, "\n")
}

func main() {
	player := Player{
		name:      "John Doe",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
	}
	fmt.Println(player, "\n")
	player.decreaseHealth(50)
	player.increaseHealth(70)
	player.decreaseEnergy(30)
	player.increaseEnergy(50)
	player.decreaseHealth(90)
	player.decreaseEnergy(75)
	player.increaseHealth(25)
	player.increaseEnergy(12)
}
