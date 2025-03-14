package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func help() {
  fmt.Println("usage: journal <yyyy-mm-dd>")
}

var targetTime time.Time

func main() {
	conf := newConfig()
	args := os.Args[1:]
	targetEntry := newEntry(time.Now(), conf)

	if len(args) != 0 {
		targetTime, err := time.Parse(time.DateOnly, args[0])
		if err != nil {
      help()
      os.Exit(1)
    }
		targetEntry = newEntry(targetTime, conf)
	}

	if !dirExists(conf.JournalDir) {
		os.Mkdir(conf.JournalDir, conf.FileMode)
	}

	entryExists := fileExists(targetEntry.path)

	if !entryExists {
		_, err := os.Create(targetEntry.path)
		check(err)
		title := []byte(targetEntry.Title())
		os.WriteFile(targetEntry.path, title, conf.FileMode)
	}

	cmd := exec.Command(conf.Editor, targetEntry.path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
  check(err)
}
