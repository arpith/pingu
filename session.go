package main

import (
	"net/http"
	"github.com/garyburd/redigo/redis"
	"bytes"
	"crypto/rand"
	"fmt"
)

func createIdAndToken() (channelId string, token string, err error) {
	c := 2
	b := make([]byte, c)
	_, err = rand.Read(b)
	if err == nil {
		channelId = string(b[0])
		token = string(b[1])
	}
	return 
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
    		// handle error
	}
	defer c.Close()

	switch r.Method {
	case "POST":
		channelId, authToken, err := createIdAndToken()
		if err != nil {
			// handle error in creating id and token
		}
		r, err := c.Do("SETNX", authToken, channelId)
		if err != nil {
			// handle redis error
		} else if r == 0 {
			// token already saved, generate a new token/channelID pair
		} else {
			// return the token and channel id to the user
		}
	}		
}


// NewBroadcastHandler creates a new handler that handles pub sub
func NewSessionHandler() func(w http.ResponseWriter, req *http.Request) {
	return ServeHTTP
}
