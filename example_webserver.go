package main

import (
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		println("--->", os.Args[1], req.URL.String())
		println("--->", os.Args[1], "X-GYDRO-ConsumerId:", req.Header.Get("X-GYDRO-ConsumerId"))
		println("--->", os.Args[1], "X-GYDRO-ConsumerCustomId:", req.Header.Get("X-GYDRO-ConsumerCustomId"))
		println("--->", os.Args[1], "X-Forwarded-For:", req.Header.Get("X-Forwarded-For"))

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(map[string]string{"content": "Hello world!"})
		w.Write(response)
	})
	http.ListenAndServe(":"+os.Args[1], nil)
}
