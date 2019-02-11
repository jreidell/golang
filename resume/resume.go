package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://rdlsvc.azurewebsites.net/api/v1/AuthUser")
	if err != nil {
		log.Fatal(err)
	}
	resumedata, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Printf("%s", resumedata)
}
