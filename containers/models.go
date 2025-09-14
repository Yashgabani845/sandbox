package main

import (
	"sync"
	"time"
)

type Job struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Source   string `json:"source"`
	Stdin    string `json:"stdin"`
	Status   string `json:"status"`
	Stdout   string `json:"stdout,omitempty"`
	Stderr   string `json:"stderr,omitempty"`
	ExitCode int    `json:"exit_code,omitempty"`

	CreatedAt time.Time `json:"createdat"`
}

var jobs = make(map[string]Job)
var mu sync.Mutex
