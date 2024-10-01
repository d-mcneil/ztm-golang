//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	TruckLift = iota
	CarLift
	MotorcycleLift
)

type Lift int

type Lifter interface {
	Lift() Lift
}

type Vehicle struct {
	modelName string
}

type Truck Vehicle

type Car Vehicle

type Motorcycle Vehicle

func (t Truck) Lift() Lift {
	return TruckLift
}

func (c Car) Lift() Lift {
	return CarLift
}

func (m Motorcycle) Lift() Lift {
	return MotorcycleLift
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", t.modelName)
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", c.modelName)
}

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", m.modelName)
}

func sendToLift(vehicle Lifter) {
	switch vehicle.Lift() {
	case TruckLift:
		fmt.Println("Directing to truck lift")
	case CarLift:
		fmt.Println("Directing to car lift")
	case MotorcycleLift:
		fmt.Println("Directing to motorcycle lift")
	}
}

func main() {
	vehicles := []Lifter{
		Truck{"Road Devourer"},
		Car{"Speedy"},
		Motorcycle{"Racer"},
	}

	for _, vehicle := range vehicles {
		fmt.Printf("-- %v --\n", vehicle)
		sendToLift(vehicle)
	}
}
