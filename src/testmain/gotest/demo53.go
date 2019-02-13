package gotest

import (
	"errors"
	// "flag"
	"fmt"
	"math"
)

var name string

func GetPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

// func init() {
// 	flag.StringVar(&name, "name", "everyone", "The greeting object.")
// }

// func main() {
// 	flag.Parse()
// 	greeting, err := hello(name)
// 	if err != nil {
// 		fmt.Printf("error: %s\n", err)
// 		return
// 	}
// 	fmt.Println(greeting, introduce())
// }

// hello 用于生成问候内容。
func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

// introduce 用于生成介绍内容。
func introduce() string {
	return "Welcome to my Golang column."
}
