package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Home struct {
	XMLName xml.Name `xml:"Home"`
	Bedroom struct {
		XMLName xml.Name `xml:"Bedroom"`
		Light   Light
		Fan     Fan
	}
	Hall struct {
		XMLName xml.Name `xml:"Hall"`
		Light   Light
		Fan     Fan
	}
	Kitchen struct {
		XMLName xml.Name `xml:"Kitchen"`
		Light   Light
	}
}

type Light struct {
	XMLName xml.Name `xml:"Light"`
	Value   string   `xml:",chardata"`
}

type Fan struct {
	XMLName xml.Name `xml:"Fan"`
	Value   string   `xml:",chardata"`
	Speed   int      `xml:"speed,attr"`
}

func (db Home) printDb() {
	fmt.Println(db)
}

func main() {
	home, err := os.Open("/Users/ravitejareddy/Desktop/home.xml")
	if err != nil {
		fmt.Println("Error opening the file")
		return
	}

	defer home.Close()

	homeData, _ := ioutil.ReadAll(home)
	var homeDB Home
	xml.Unmarshal(homeData, &homeDB)
	fmt.Println(homeDB)
}
