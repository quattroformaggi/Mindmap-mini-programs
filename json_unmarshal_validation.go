package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

/*To complete this task, you will need the json.Unmarshal function, which decodes JSON bytes into a structure.

Implement the DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) function, which decodes the JSON request body into a CreateUserRequest structure and validates it.

If invalid JSON is received or the structure is filled incorrectly, the function should return an error.*/

// CreateUserRequest is a request to create a new user.
type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	cur := CreateUserRequest{}

	err := json.Unmarshal(requestBody, &cur)

	if len(cur.Email) < 1 {
		return CreateUserRequest{}, errEmailRequired
	} else if len(cur.Password) < 1 {
		return CreateUserRequest{}, errPasswordRequired
	} else if len(cur.PasswordConfirmation) < 1 {
		return CreateUserRequest{}, errPasswordConfirmationRequired
	} else if cur.Password != cur.PasswordConfirmation {
		return CreateUserRequest{}, errPasswordDoesNotMatch
	}

	return cur, err
}

func main() {
	req, err := DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}"))
	fmt.Println(req, err)

	req, err = DecodeAndValidateRequest([]byte("{\"email\":\"l\",\"password\":\"l\",\"password_confirmation\":\"ll\"}"))
	fmt.Println(req, err)
}
