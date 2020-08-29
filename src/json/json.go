package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	loginRequest := LoginRequest{}
	// incoming string
	body := `{"username":"Alex","password":"123"}`

	// parse json - option 1

	/*
		// create Reader stream
		buffer := bytes.NewBufferString(body)
		// map buffer of bytes to struct
		err := json.NewDecoder(buffer).Decode(&loginRequest)
	*/

	// parse json - option 2

	// map body bytes to struct
	err := json.Unmarshal([]byte(body), &loginRequest)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("loginRequest parsed", loginRequest)

	// stringify json

	bytes, err := json.Marshal(&loginRequest)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("loginRequest bytes", string(bytes))
}

// LoginRequest is a body of login request
type LoginRequest struct {
	Username string `json:"username"` // `json:"username,omitempty"` omitempty - exclude uninitialized field from json
	Password string `json:"password"`
}

/*
// MarshalJSON - overwrite standard json.Marshal for LoginRequest struct
func (loginRequest *LoginRequest) MarshalJSON() ([]byte, error) {
	return []byte(`{"hello":"world"}`), nil
}

// UnmarshalJSON - overwrite standard json.Unmarshal for LoginRequest struct
func (loginRequest *LoginRequest) UnmarshalJSON([]byte) error {
	return nil
}
*/
