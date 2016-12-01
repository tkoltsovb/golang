package main

import 
(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(){

	fmt.Println("start")

	resp, err := http.Get("http://ads.sputnik.ru:80/choose?mime-type=application/javascript&adblocks=1&callback=name")//"http://sputnik.ru/")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.ContentLength, " status = ", resp.StatusCode)

	defer resp.Body.Close()

	//_, err = io.Copy(os.Stdout, resp.Body)
	log.Println(resp.StatusCode)
	 
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main(){
	check()

}