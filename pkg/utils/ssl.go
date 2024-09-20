package utils

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func Checkssl(url string) error {
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
		return err
	}
	defer resp.Body.Close()

	if resp.TLS != nil {
		log.Println("SSL certificate exists!")
	}

	return nil
}

