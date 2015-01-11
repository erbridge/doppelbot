package twitter

import (
	"encoding/json"
	"io/ioutil"
)

type (
	ConsumerConfig struct {
		key    string
		secret string
	}

	AccessConfig struct {
		token       string
		tokenSecret string
	}
)

func NewConsumerConfig(key string, secret string) ConsumerConfig {
	return ConsumerConfig{
		key:    key,
		secret: secret,
	}
}

func NewAccessConfig(token string, tokenSecret string) AccessConfig {
	return AccessConfig{
		token:       token,
		tokenSecret: tokenSecret,
	}
}

func LoadConfig(f string) (c ConsumerConfig, a AccessConfig, err error) {
	var conf map[string]string
	bytes, _ := ioutil.ReadFile(f)
	if err = json.Unmarshal(bytes, &conf); err != nil {
		return
	}
	c.key = conf["consumer_key"]
	c.secret = conf["consumer_secret"]
	a.token = conf["access_token"]
	a.tokenSecret = conf["access_token_secret"]
	return
}
