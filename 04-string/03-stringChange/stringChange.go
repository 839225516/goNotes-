package main

import "fmt"

func main() {
	angel := "heros never die"
	angelByes := []byte(angel)

	for i := 5; i <= 10; i++ {
		angelByes[i] = ' '
	}

	fmt.Println(string(angelByes))
}
