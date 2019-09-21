package project

import "fmt"

//go:generate go run gen.go
func PrintContributers() {
	for _, c := range Contributors {
		fmt.Println(c)
	}
}
