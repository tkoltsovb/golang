package main

import 
(
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func check(){

	fmt.Println("start")
	t0 := time.Now()
	start := t0.UnixNano()
	

	for i := 0; i < 5000; i++ {
		//resp, err := http.Get("http://ads.sputnik.ru:80/choose?mime-type=application/javascript&adblocks=1&callback=name")//"http://sputnik.ru/")
		resp, err := http.Get("http://127.0.0.1:8080/view/book")//"http://sputnik.ru/")
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

	t1 := time.Now()
	end := t1.UnixNano()
	fmt.Println("WorkTime = ", end-start)
	fmt.Printf("duration = %v\n", t1.Sub(t0))
}

func get() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:50080/info", nil)
	req.Header.Add("X-Forwarded-For", "123.56.5.5")
	_, err := client.Do(req)
	if err != nil{
		fmt.Println(err, err.Error())
		os.Exit(1)
	}

}

func main(){
	//check()
	get()
}