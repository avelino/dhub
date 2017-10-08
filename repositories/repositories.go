package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avelino/dhub/helpers"
	"github.com/avelino/dhub/users"
)

func BuildHistory(a users.Auth, user, repo string) (bh buildHistory) {
	uri := fmt.Sprintf("/repositories/%s/%s/buildhistory/", user, repo)
	req, err := http.NewRequest("GET", helpers.URL+uri, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", a.Token()))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&bh)
	return
}
