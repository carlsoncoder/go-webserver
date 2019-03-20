package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	kubernetestypes "github.com/carlsoncoder/go-webserver/kubernetestypes"
	usertypes "github.com/carlsoncoder/go-webserver/usertypes"
)

func usertest(rw http.ResponseWriter, req *http.Request) {
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

func kubernetestest(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	var eventList kubernetestypes.EventList

	err = json.Unmarshal(body, &eventList)
	if err != nil {
		panic(err)
	}

	log.Println(len(eventList.Items))

	for i, event := range eventList.Items {
		log.Println(fmt.Sprintf("Event #%d", i+1))
		log.Println(event.StageTimestamp)
		log.Println(event.Level)
		log.Println(event.Stage)
		log.Println(event.RequestURI)
		log.Println(event.Verb)
		log.Println(event.User.UserName)
		log.Println(event.GetSourceIPAddress())
		log.Println(event.ObjectRef.Name)
		log.Println(event.ObjectRef.Resource)
	}
}

func main() {
	http.HandleFunc("/usertest", usertest)
	http.HandleFunc("/kubernetestest", kubernetestest)
	log.Fatal(http.ListenAndServe(":80", nil))
}
