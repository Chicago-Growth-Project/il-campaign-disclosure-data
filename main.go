package main

import "fmt"

func main() {
	for _, table := range AllTables {
		err := table.Create()
		if err != nil {
			fmt.Println(err)
		}
	}
}
