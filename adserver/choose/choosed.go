
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)


type Query struct {
	XMLName xml.Name `xml:"document"`
	Test string `xml:"test"`
	
	Httpserver struct {
			Port string `xml:"port"`
		}	`xml:"httpserver"`
	
	Run struct {
		Daemon string `xml:"daemon"`
	} `xml:"run"`
}

func readCommandArgs() (error, string) {
	confPtr := flag.String("c", "", "Config file name")
	flag.Parse()

	return nil, *confPtr
}

func main() {
	err, configName := readCommandArgs()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Config =", configName)

	xmlFile, err := os.Open(configName)
	defer xmlFile.Close()
	if err != nil {
		fmt.Println("Cannot open file: ", err.Error())
		os.Exit(1)
	}

	buf, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("error ocurred while read")
		os.Exit(1)
	}
	fmt.Println("buf =", string(buf))

	query := Query{}
	err = xml.Unmarshal(buf, &query)
	if err != nil {
		fmt.Println("Error occured while unmarshal: ", err.Error())
	}

	fmt.Println("daenon =", query.Run.Daemon)
	fmt.Println("port =", query.Httpserver.Port)
}