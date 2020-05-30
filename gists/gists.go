package gists

/*
source : https://amzn.to/2ArHZ9y
p40 インタフェースを使ったモック
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gist struct {
	Rawurl string `json:"html_url"`
}

type Doer interface {
	doGistsRequest(user string) (io.Reader, error)
}

type Client struct {
	Gister Doer
}

type Gister struct{}

func (g *Gister) doGistsRequest(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}

	return &buf, nil
}

func (c *Client) ListGists(user string) ([]string, error) {
	r, err := c.Gister.doGistsRequest(user)
	if err != nil {
		return nil, err
	}
	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}
	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}
	return urls, nil
}
