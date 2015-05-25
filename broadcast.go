package main

import (
	"net/http"
	"github.com/manucorporat/sse"
	"github.com/garyburd/redigo/redis"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	channelName := r.URL.Path
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
    		// handle error
	}
	defer c.Close()
	psc := redis.PubSubConn{c}
	switch r.Method {
	case "GET":
		psc.Subscribe(channelName)
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				sse.Encode(w, sse.Event{
     					Event: "message",
					Data:  v.Data,
				})
			case error:
				// handle error
			}
		}
	
	case "POST":
		r.ParseForm()
		savedChannelName, err = redis.String(c.Do("GET", r.Form["token"]))
		if err != nil {
			// handle error
		} else if savedChannelName != channelName {
			// handle authentication error
		} else {
			c.Do("PUBLISH", channelName, r.Form["message"])
		}
	}

}

// NewBroadcastHandler creates a new handler that handles pub sub
func NewBroadcastHandler() func(w http.ResponseWriter, req *http.Request) {
	return ServeHTTP
}
