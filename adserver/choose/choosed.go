
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	//"io"
	"io/ioutil"
	"os"
)


type Config struct {
	XMLName xml.Name `xml:"document"`
	
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

func readConfig(configName string, config *Config) error {
	xmlFile, err := os.Open(configName)
	defer xmlFile.Close()
	if err != nil {
		fmt.Println("Cannot open file: ", err.Error())
		return err
	}

	buf, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("error ocurred while read")
		return err
	}

	err = xml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println("Error occured while unmarshal: ", err.Error())
		return err
	}

	return nil
}

/*func infoHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello, world!\n")

	fmt.Println("remoteaddr = ", request.RemoteAddr)
	header := request.Header.Get("X-Forwarded-For")
	//if len(header) != 0 {
		fmt.Println("X-Forwarded-For = ", header)
	//}
}*/

func main() {
	err, configName := readCommandArgs()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Config =", configName)

	var config Config

	err = readConfig(configName, &config)
		if err != nil {
		fmt.Println("Cannot read config: ", err.Error(), " Conf name = ", configName)
		os.Exit(1)
	}

	fmt.Println("daenon =", config.Run.Daemon)
	fmt.Println("port =", config.Httpserver.Port)

	http.HandleFunc("/info/", infoHandler)
	//http.ListenAndServe(":" + config.Httpserver.Port, nil)
	log.Fatal(http.ListenAndServe(":50080", nil))

	fmt.Println("end")
}