package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/* NOTE: Don't use seperate struct for 'just value' nodes like 'Light'. Directly declare them as string with node name (Light) as Tag. */
/* NOTE: Struct fields can be tagged but not struct name directly. For tagging struct name, you need to use xml.Name field */
type Home struct {
	XMLName xml.Name `xml:"Home"`
	Bedroom struct {
		XMLName xml.Name `xml:"Bedroom"`
		Light   string   `xml:"Light"`
		Fan     Fan      `xml:"Fan"`
	}
	Hall struct {
		XMLName xml.Name `xml:"Hall"`
		Light   string   `xml:"Light"`
		Fan     Fan      `xml:"Fan"`
	}
	Kitchen struct {
		XMLName xml.Name `xml:"Kitchen"`
		Light   []string `xml:"Light"`
	}
}

type Fan struct {
	Value string `xml:",chardata"`
	Speed int    `xml:"speed,attr"`
}

func (db Home) printDb() {
	fmt.Println(db)
}

func main() {
	homeXML, err := os.Open("home.xml")
	if err != nil {
		fmt.Println("Error opening the file")
		return
	}

	defer homeXML.Close()

	homeData, _ := ioutil.ReadAll(homeXML)
	var xHomeDb Home
	xml.Unmarshal(homeData, &xHomeDb)
	xHomeDb.printDb()
}

/* http://jan.newmarch.name/go/xml/chapter-xml.html */
