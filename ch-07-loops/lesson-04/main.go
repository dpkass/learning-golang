package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		out := ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}

		if out == "" {
			fmt.Println(i)
		} else {
			fmt.Println(out)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
