package pages

import (
	"encoding/json"
	"io/ioutil"
)

// Config base type
type Config map[string]interface{}

// Name Config type Name string value
func (c Config) Name() string {
	return c["name"].(string)
}

// Author the author of the post
func (c Config) Author() string {
	return c["author"].(string)
}

// LoadConfig Loads a json config file
func LoadConfig(path string) (Config, error) {
	var m Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(data, &m)
	return m, err
}
