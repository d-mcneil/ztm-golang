//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() float32 {
	var total Bytes
	for _, v := range b.amount {
		total += v
	}
	return float32(total) / float32(len(b.amount))
}

func (c *CpuTemp) AverageTemp() float32 {
	var total Celcius
	for _, v := range c.temp {
		total += v
	}
	return float32(total) / float32(len(c.temp))
}

func (m *MemoryUsage) AverageMemory() float32 {
	var total Bytes
	for _, v := range m.amount {
		total += v
	}
	return float32(total) / float32(len(m.amount))
}

func PrintAverage(n float32) {
	fmt.Printf("Average: %v\n", n)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		bandwidth,
		temp,
		memory,
	}

	fmt.Println("5-second average:")
	fmt.Println("Bandwidth:")
	PrintAverage(dashboard.AverageBandwidth())
	fmt.Println("Temperature:")
	PrintAverage(dashboard.AverageTemp())
	fmt.Println("Memory:")
	PrintAverage(dashboard.AverageMemory())
}
