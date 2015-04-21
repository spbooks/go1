package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	type Config struct {
		Name     string `json:"SiteName"`
		URL      string `json:"SiteUrl"`
		Database struct {
			Name     string
			Host     string
			Port     int
			Username string
			Password string
		}
	}
	conf := Config{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Site: %s (%s)\n", conf.Name, conf.URL)
	db := conf.Database
	fmt.Printf(
		"DB: mysql://%s:%s@%s:%d/%s\n",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

}
