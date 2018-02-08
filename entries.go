package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"encoding/xml"
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

func getEntries(path string) db{
	fmt.Println("start dbGen")
	xmlFile, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open xml")
	}
	defer xmlFile.Close()

	var freqList db

	data, _ := ioutil.ReadAll(xmlFile)
	xmlFile.Close()
	xml.Unmarshal(data, &freqList)

	return freqList
}

