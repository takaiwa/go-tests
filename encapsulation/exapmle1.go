package encapsulation

import "encoding/json"

/*
とてもわかりやすいJsonを扱うサンプル
https://ja.coder.work/so/go/1436541
*/
type User interface {
	Name() string
}

// userと頭文字を小文字にすることで他から参照できない
type user struct {
	Username string `json:"name"`
}

func (u *user) Name() string {
	return "Mr. " + u.Username
}

func ParseUserData(data []byte) (User, error) {
	user := &user{}
	if err := json.Unmarshal(data, user); err != nil {
		return nil, err
	}
	return user, nil
}
