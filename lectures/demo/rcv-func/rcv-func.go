package main

import "fmt"

type Space struct {
	isOccupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(parkingLot *ParkingLot, spaceIndex int) {
	parkingLot.spaces[spaceIndex].isOccupied = true
}

func (parkingLot *ParkingLot) occupySpace(spaceIndex int) {
	parkingLot.spaces[spaceIndex].isOccupied = true
}

func (parkingLot *ParkingLot) vacateSpace(spaceIndex int) {
	parkingLot.spaces[spaceIndex].isOccupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 4)}

	fmt.Println(lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println(lot)

	lot.vacateSpace(2)
	fmt.Println(lot)
}
