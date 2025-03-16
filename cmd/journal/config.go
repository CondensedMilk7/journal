package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	JournalDir    string      `json:"journalDir"`
	Editor        string      `json:"editor"`
	FileMode      os.FileMode `json:"fileMode"`
	FileExtension string      `json:"fileExtension"`
}

func newConfig() *config {
	userConfig, err := newUserConfig()
	if err == nil {
		return userConfig
	}
	return newDefaultConfig()
}

func newUserConfig() (c *config, e error) {
	homeDir, e := os.UserHomeDir()
	conf, e := os.ReadFile(fmt.Sprintf("%s/.config/journal/config.json", homeDir))
	var unmarshaled config
	e = json.Unmarshal(conf, &unmarshaled)

	return &unmarshaled, e
}

func newDefaultConfig() *config {
	c := config{}
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c.JournalDir = fmt.Sprintf("%s/%s/%s", homeDir, "Documents", "journal")
	c.Editor = "nvim"
	c.FileMode = 600
	c.FileExtension = "md"

	return &c
}
