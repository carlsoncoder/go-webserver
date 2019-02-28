package usertypes

import (
	generictypes "github.com/carlsoncoder/webserver/generictypes"
)

type User struct {
	Name       string
	Age        int
	Occupation string
	Hobbies    []generictypes.Hobby
}
