package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits++
	fmt.Println("Counter hits:", counter.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	counter.hits++
}

func main() {
	counter := Counter{}

	hello := "Hello"
	world := "World"

	fmt.Println("Before replace:", hello, world)

	replace(&hello, "Hi", &counter)

	fmt.Println("After replace:", hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)
	replace(&phrase[1], "Go", &counter)
	fmt.Println(phrase)
	fmt.Println("After slice without pointer:", hello, world)

	phrase2 := []*string{&hello, &world}
	fmt.Println(phrase2)
	replace(phrase2[1], "Go", &counter)
	fmt.Println(phrase2)
	fmt.Println("After slice with pointer:", hello, world)

	fmt.Println(counter.hits)
}
