package main

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
	"log"
)
type db struct {
	Name string `xml:"name"`
	Generator string `xml:"generator"`
	LastUpdate string `xml:"lastupdate"`
	Comment string `xml:"comment"`
	Entries []entry `xml:"entry"`
}

type entry struct {
	Freq string `xml:"freq"`
	Transmission string `xml:"transmission"`
	Network string `xml:"network"`
	FrqID string `xml:"freqID"`
	Comment string `xml:"comment"`
	IgnoreBearing string `xml:"ignorebearing"`
	IgnoreLocation string `xml:"ignorelocation"`
}


func main() {
	fmt.Println("start dbGen")
	xmlFile, err := os.Open("data/DB_CHN.xml")
	if err != nil {
		log.Printf("failed to open xml")
	}
	defer xmlFile.Close()

	var freqList db

	data, _ := ioutil.ReadAll(xmlFile)
	xmlFile.Close()
	xml.Unmarshal(data, &freqList)
	fmt.Println(freqList.Name)
	fmt.Println(freqList.Generator)
	fmt.Println(freqList.LastUpdate)
	fmt.Println(freqList.Comment)

	for _, en := range freqList.Entries {
		fmt.Println(en.Transmission)
	}
}
