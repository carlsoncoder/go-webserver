// Test it by running the following:
// curl -X POST -H 'Content-Type: application/json' -d "{\"name\": \"justin\", \"age\": 34}" http://localhost/test
// or
// curl -X POST -H 'Content-Type: application/json' -d @jsonFile.json http://localhost/test
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	usertypes "github.com/carlsoncoder/webserver/usertypes"
)

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	var t usertypes.User

	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	log.Println(t.Name)
	log.Println(t.Age)
	log.Println(t.Occupation)

	for i, hobby := range t.Hobbies {
		log.Println(fmt.Sprintf("Hobby #%d", i+1))
		log.Println(fmt.Sprintf("%s for %d years", hobby.Name, hobby.Years))
	}
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":80", nil))
}
