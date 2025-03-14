package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func dirExists(filepath string) bool {
	_, err := os.ReadDir(filepath)

	if err != nil {
		return false
	}

	return true
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)

	if err != nil {
		return false
	}

	return true
}
