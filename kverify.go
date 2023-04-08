package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, fileErr := os.ReadFile("/app/index.html")
	if fileErr != nil {
		log.Fatalf("Failed to load index.html: %s", fileErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("CF-Connecting-IP")

		if strings.HasPrefix(r.URL.Path, "/unban") {
			qq := r.URL.Query().Get("qq")
			_, atoiErr := strconv.Atoi(qq)

			if atoiErr != nil {
				log.Printf("BADRQ [%s]", ip)
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte{})

				return
			}

			log.Printf("UNBAN [%s] [%s]", ip, qq)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte{})

			return
		}

		log.Printf("SERVE [%s]", ip)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bytes)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
