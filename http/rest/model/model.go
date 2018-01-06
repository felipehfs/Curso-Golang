package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Protocolo representa isso mesmo
type Protocolo struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Nome    string        `json:"nome"`
	Checado bool          `json:"checado"`
	Criacao *time.Time    `json:"criacao"`
}
