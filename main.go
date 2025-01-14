package main

import "fmt"

func main() {
	err := Candidates.Create()
	if err != nil {
		fmt.Println(err)
	}
}
