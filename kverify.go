package main

import (
	"fmt"
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

		if r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/unban") {
			formErr := r.ParseForm()

			if formErr != nil {
				log.Printf("BADFO [%s]", ip)
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("<!DOCTYPE html><html><head><meta charset=\"utf-8\"></head><body><p>你的 QQ 号输入有误，请返回重新检查哦~</p></body></html>"))

				return
			}

			qq := r.Form.Get("qq")
			_, atoiErr := strconv.Atoi(qq)

			if atoiErr != nil {
				log.Printf("BADQQ [%s] [%s]", ip, qq)
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("<!DOCTYPE html><html><head><meta charset=\"utf-8\"></head><body><p>你的 QQ 号输入有误，请返回重新检查哦~</p></body></html>"))

				return
			}

			log.Printf("UNBAN [%s] [%s]", ip, qq)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("<!DOCTYPE html><html><head><meta charset=\"utf-8\"></head><body><p>提交成功，你现在可以在用户群参与讨论了~讨论时记得热情、友善哦~</p></body></html>"))

			go func() {
				_, getErr := http.Get(fmt.Sprintf("http://bot-kfm.sh1-legacy/unban?qq=%s", qq))
				if getErr != nil {
					log.Printf("GETER [%s]", getErr)
				}
			}()

			return
		}

		log.Printf("SERVE [%s]", ip)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bytes)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
