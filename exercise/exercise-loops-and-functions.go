package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// 1. Loop 10 times
	//z := 1.0
	//for i := 0; i < 10; i++ {
	//	fmt.Println(z)
	//	z -= (z * z - x) / (2 * z)
	//}
	//return z

	// 2. Stop loop when value does not change
	//z := 1.0
	//bz := 0.0
	//for cnt := 0; z != bz && math.Abs(z -bz) > 0.0000000000001; cnt++ {
	//	fmt.Println("z:", z, "bz:", bz, "cnt:", cnt)
	//	bz = z
	//	z -= (z * z - x) / (2 * z)
	//}
	//return z

	// 3. z = x
	//z := float64(x)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(z)
	//	z -= (z * z - x) / (2 * z)
	//}
	//return z

	// 4. z = x / 2
	z := float64(x / 2)
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
