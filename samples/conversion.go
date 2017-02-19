package main

import "fmt"

func main() {
	type T1 struct {
		X int `json:"foo"`
	}
	type T2 struct {
		X int `json:"bar"`
	}
	var v1 T1
	var v2 T2
	v2.X = 12
	v1 = T1(v2) // now legal // HL
	fmt.Println(v1)
}
