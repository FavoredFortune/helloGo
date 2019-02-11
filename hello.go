package main

import (
	"fmt"
	"math"

	"github.com/FavoredFortune/stringutil"

	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("A fun number is", rand.Int())
	fmt.Println("Who doesn't love", math.Pi, "?")
	fmt.Println("Let's do some addition!", add(433, 226), "is 433 plus 226")
}
