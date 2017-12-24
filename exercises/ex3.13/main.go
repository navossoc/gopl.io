package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Printf("KB = %.3e\n", float64(KB))
	fmt.Printf("MB = %.3e\n", float64(MB))
	fmt.Printf("GB = %.3e\n", float64(GB))
	fmt.Printf("TB = %.3e\n", float64(TB))
	fmt.Printf("PB = %.3e\n", float64(PB))
	fmt.Printf("EB = %.3e\n", float64(EB))
	fmt.Printf("ZB = %.3e\n", float64(ZB))
	fmt.Printf("YB = %.3e\n", float64(YB))
}
