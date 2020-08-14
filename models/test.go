package models

import (
	"log"
	"time"
)

type Test struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (t *Test) C() int {
	log.Printf(" === C started for (%d, %d) at %d\n", t.A, t.B, time.Now().UnixNano())
	<-time.After(5 * time.Second)
	log.Printf("C ended for (%d, %d) at %d ===\n", t.A, t.B, time.Now().UnixNano())
	return t.A * t.B
}

func (t *Test) D() int {
	log.Printf(" === D started for (%d, %d) at %d\n", t.A, t.B, time.Now().UnixNano())
	<-time.After(2 * time.Second)
	log.Printf("D ended for (%d, %d) at %d ===\n", t.A, t.B, time.Now().UnixNano())
	return t.A + t.B
}

func (Test) IsTestResult() {}
