package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	usertypes "github.com/carlsoncoder/go-webserver/usertypes"
)

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	var user usertypes.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	log.Println(user.Name)
	log.Println(user.Age)
	log.Println(user.Occupation)

	for i, hobby := range user.Hobbies {
		log.Println(fmt.Sprintf("Hobby #%d", i+1))
		log.Println(fmt.Sprintf("%s for %d years", hobby.Name, hobby.Years))
	}
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":80", nil))
}
