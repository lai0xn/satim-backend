package utils

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Checkssl(url string) error {
	fmt.Println("Checking SSL certificate for", url)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 40 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v\n", url, err)
		return err
	}
	defer resp.Body.Close()

	if resp.TLS != nil {
		log.Println("SSL certificate exists for", url)
		return nil
	} else {
		log.Println("No SSL certificate found for", url)
		return errors.New("No SSL certificate found")
	}
}
