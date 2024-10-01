package main

import (
	// "errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		// return 0, errors.New(fmt.Sprintf("Index %d out of range", index))
		return 0, fmt.Errorf("index %d out of range", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{values: []int{1, 2, 3, 4, 5}}
	value, err := stuff.Get(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
