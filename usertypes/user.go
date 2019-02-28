package usertypes

import (
	generictypes "github.com/carlsoncoder/go-webserver/generictypes"
)

type User struct {
	Name       string
	Age        int
	Occupation string
	Hobbies    []generictypes.Hobby
}
