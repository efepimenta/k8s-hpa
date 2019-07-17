package main

import (
	"fmt"
	"math"
)

func dosqrt(a float64) float64 {
    return math.Sqrt(a)
}

func main() {

    sum := 0.0
    for i := 0.0; i < 1000000000.0; i++ {
    	sum += dosqrt(i)
    }

	fmt.Printf("Code.education Rocks!\n");
	// fmt.Printf("%.1f\n", sum);
}
