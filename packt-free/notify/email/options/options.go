package options

import (
	"encoding/json"
	"io/ioutil"

	"github.com/JSila/packtFree/notify"
)

// Options holds init data for mailgun client
type Options struct {
	Domain    string       `json:"domain"`
	APIKey    string       `json:"apikey"`
	PublicKey string       `json:"publickey"`
	From      *notify.User `json:"from"`
}

// FromFile return options from json file
func FromFile(file string) (*Options, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	opts := new(Options)
	if err := json.Unmarshal(content, opts); err != nil {
		return nil, err
	}
	return opts, nil
}
