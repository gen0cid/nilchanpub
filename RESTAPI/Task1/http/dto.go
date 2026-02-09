package http

import (
	"encoding/json"
	"errors"
	"time"
)

type massageDTO struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Pages  int    `json:"Pages"`
}

type ReadDTO struct {
	Read bool `json:"Read"`
}

func (m massageDTO) validateForCreate() error {
	if m.Title == "" {
		return errors.New("Title is empty")

	}
	if m.Author == "" {
		return errors.New("Description is empty")

	}
	if m.Pages == 0 {
		return errors.New("zero pages")
	}
	return nil
}

type errorDTO struct {
	Message string
	Time    time.Time
}

func (e errorDTO) toString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
