package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	URL = "https://hub.docker.com/v2"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Token string `json:"token"`
}

type BuildHistoryItem struct {
	Id        int    `json:"id"`
	Status    int    `json:"status"`
	TagName   string `json:"dockertag_name"`
	BuildCode string `json:"build_code"`
}

type BuildHistory struct {
	Count int                `json:"count"`
	Itens []BuildHistoryItem `json:"results"`
}

func (a Auth) token() (l Login) {
	url := fmt.Sprintf("%s/users/login/", URL)
	bff := new(bytes.Buffer)
	json.NewEncoder(bff).Encode(a)
	req, err := http.NewRequest("POST", url, bff)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&l)
	return
}

func main() {
	a := Auth{
		Username: "",
		Password: "",
	}
	login := a.token()

	url := fmt.Sprintf("%s/repositories/nuveo/runners/buildhistory/", URL)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", login.Token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bh := BuildHistory{}
	json.NewDecoder(resp.Body).Decode(&bh)
	fmt.Println(bh)
}
