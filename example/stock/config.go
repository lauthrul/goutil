package main

import (
	"fmt"
	"io/ioutil"
)

type Config struct {
	FilePath string   `json:"-"`
	Indexes  []string `json:"indexes"`
	Stocks   []string `json:"stocks"`
}

func (c *Config) Load() error {
	if c.FilePath == "" {
		return fmt.Errorf("no file specificed")
	}
	return c.LoadFile(c.FilePath)
}

func (c *Config) LoadFile(file string) error {
	if file == "" {
		return fmt.Errorf("no file specificed")
	}
	c.FilePath = file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}
