package pkg

import (
	"log"
	"net/http"
	"time"
	"crypto/tls"
)

func Checkssl(url string ) bool {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.TLS != nil {
		log.Println("SSL certificate exists!")
		return true
	} else {
		log.Println("No SSL certificate found.")
		return false
	}
}