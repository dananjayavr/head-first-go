package main

import (
	"fmt"
	"os"
)

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	for _, n := range numbers {
		fmt.Printf("%d: %04b\n", n, n)
	}

	fmt.Printf("%d => %04b & %04b = %04b => %d\n", 2,2,1,2 & 1,2 & 1)
	fmt.Printf("%d => %04b | %04b = %04b => %d\n", 2,2,1,2 | 1, 2 | 1)

	fmt.Printf("RDONLY = %016b\n", os.O_RDONLY)
	fmt.Printf("WRONLY = %016b\n", os.O_WRONLY)
	fmt.Printf("RDWR = %016b\n", os.O_RDWR)
	fmt.Printf("CREATE = %016b\n", os.O_CREATE)
	fmt.Printf("APPEND = %016b\n", os.O_APPEND)

	fmt.Printf("%016b\n", os.O_CREATE|os.O_APPEND)
}
