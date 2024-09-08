package main

import "fmt"

type Room struct {
	name      string
	isCleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.isCleaned {
			fmt.Println(room.name, "is cleaned")
		} else {
			fmt.Println(room.name, "is not cleaned")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Living Room"},
		{name: "Bedroom"},
		{name: "Kitchen"},
		{name: "Bathroom"},
	}

	checkCleanliness(rooms)

	fmt.Println("Cleaning the rooms...")

	rooms[0].isCleaned = true
	rooms[1].isCleaned = true

	checkCleanliness(rooms)
}
