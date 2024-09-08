//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(assemblyLine []Part) {
	fmt.Println("=== Assembly Line ===")
	for i := 0; i < len(assemblyLine); i++ {
		j := i
		fmt.Println("Part", j+1, "->", assemblyLine[j])
	}
}

func main() {

	assemblyLine := make([]Part, 3)

	// assemblyLine := []Part{
	// 	"Thermocouple",
	// 	"Pressure Sensor",
	// 	"Valve",
	// }

	assemblyLine[0] = "Thermocouple"
	assemblyLine[1] = "Pressure Sensor"
	assemblyLine[2] = "Valve"

	printAssemblyLine(assemblyLine)

	assemblyLine2 := []Part{
		"Motor",
		"Pump",
	}

	assemblyLine = append(assemblyLine, assemblyLine2...)

	printAssemblyLine(assemblyLine)

	assemblyLine = assemblyLine[3:]

	printAssemblyLine(assemblyLine)
}
