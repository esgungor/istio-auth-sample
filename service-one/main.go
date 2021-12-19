package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserInfo struct {
	Username string `json:"preferred_username"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	//Get parsed user token
	h := r.Header.Get("jwt-parsed")
	decoded, err := base64.StdEncoding.DecodeString(h)
	if err != nil {
		panic("Base64 decode operation failed!")
	}
	user := UserInfo{}
	err = json.Unmarshal(decoded, &user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Username: %v", user.Username)
}

func main() {
	http.HandleFunc("/user", greet) // service endpoint
	http.ListenAndServe(":3000", nil)
}
