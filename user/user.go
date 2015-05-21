package user

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("User not found")

// User fields from stream.me
type User struct {
	DisplayName string `json:displayName`
	Username    string `json:"username"`
	Slug        string `json:"slug"`
	PublicId    string `json:"publicId"`
	ChatRoomId  string `json:"chatRoomId"`
}

// Get gets a user from stream.me, client should be pre authed
func Get(client *http.Client) (*User, error) {
	resp, err := client.Get("https://pds.dev.ifi.tv/api-user/v1/me")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, ErrNotFound
	}

	u := &User{}
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, err
	}

	if len(u.PublicId) == 0 {
		return nil, ErrNotFound
	}
	u.ChatRoomId = "user:" + u.PublicId + ":web"

	return u, nil
}
