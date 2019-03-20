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
	var userList usertypes.UserList

	err = json.Unmarshal(body, &userList)
	if err != nil {
		panic(err)
	}

	log.Println(len(userList.Users))

	for i, user := range userList.Users {
		log.Println(fmt.Sprintf("User #%d", i+1))
		log.Println(user.Name)
		log.Println(user.Age)
		log.Println(user.Occupation)

		for i, hobby := range user.Hobbies {
			log.Println(fmt.Sprintf("Hobby #%d", i+1))
			log.Println(fmt.Sprintf("%s for %d years", hobby.Name, hobby.Years))
		}
	}
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":80", nil))
}
