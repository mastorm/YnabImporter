package comdirect

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Credentials struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
}

type comDirectClient struct {
	Credentials Credentials
}

func (c *Credentials) ToFormBody(grantType string) *url.Values {
	body := &url.Values{}
	body.Add("client_id", c.ClientId)
	body.Add("client_secret", c.ClientSecret)
	body.Add("grant_type", grantType)
	body.Add("username", c.Username)
	body.Add("password", c.Password)
	return body
}

const (
	ApiEndpoint = "https://api.comdirect.de/"
	OauthToken  = "oauth/token"
)

func NewClient(creds Credentials) *comDirectClient {
	return &comDirectClient{Credentials: creds}
}

func (c *comDirectClient) RequestToken() error {
	formBody := *c.Credentials.ToFormBody("password")
	res, err := http.PostForm(ApiEndpoint+OauthToken, formBody)
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(resBody))
	return nil
}
