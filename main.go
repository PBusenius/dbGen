package main

import (
	"fmt"
)

func main() {
	fmt.Println("start dbGen")
	frqList := getEntries("data/DB_CHN.xml")
	fmt.Println(frqList.Name)

	for _, en := range frqList.Entries {
		fmt.Println(en.Transmission)
		fmt.Println(en.Comment)
	}
}
