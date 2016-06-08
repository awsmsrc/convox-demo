package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
)

func main() {
	ConnectRedis()

	// read config from environment
	port := os.Getenv("PORT")

	// start server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello, %q</h1>", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening=" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ConnectRedis() {
	c, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		// handle connection error
		log.Fatal(err.Error())
	}

	defer c.Close()

	p, err := c.Do("PING")
	log.Printf("%#v", p)
}
