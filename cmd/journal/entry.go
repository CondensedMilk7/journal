package main

import (
	"fmt"
	"time"
)

type entry struct {
	slug string
	path string
	time time.Time
}

func (e *entry) Title() string {
	title := e.time.Format("Monday, January 2, 2006")
	return fmt.Sprintf("# %s \n", title)
}

func newEntry(t time.Time, c *config) *entry {
	e := entry{}
	e.time = t
	e.slug = t.Format("2006-01-02")
	e.path = fmt.Sprintf("%s/%s.%s", c.JournalDir, e.slug, c.FileExtension)

	return &e
}
