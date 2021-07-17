package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	hostname, _ = os.Hostname()
)

func main() {
	addr := os.Getenv("TARGET_SERVER")
	mode := os.Getenv("NO_CLIENT")
	client := http.DefaultClient
	client.Timeout = time.Second * 2
	if strings.ToLower(mode) != "true" {
		go func() {
			time.Sleep(time.Second * 2)
			for {
				time.Sleep(time.Millisecond * 500)
				resp, err := client.Get(fmt.Sprintf("http://%s", addr))
				if err != nil {
					fmt.Println(hostname, err)
					continue
				}
				if resp.Body == nil {
					fmt.Println(hostname, "body is nil")
					continue
				}
				bs, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(hostname, err)
					continue
				}
				fmt.Printf("%s call to %s\n", hostname, string(bs))
			}
		}()
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(hostname))
}
