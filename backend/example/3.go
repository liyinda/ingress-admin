package main

import (
	"encoding/json"
	"fmt"
	//"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	//group := ColorGroup{
	//	ID:     1,
	//	Name:   "Reds",
	//	Colors: []string{"Crimson", "Red", "Ruby", "liyinda"},
	//}
        group := new(ColorGroup)
        group.ID = 1

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)
	fmt.Println(string(b))
}
