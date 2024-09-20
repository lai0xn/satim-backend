package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	caCert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatal("Failed to load CA certificate:", err)
	}

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		if resp.Request.URL.String() == "https://api.naviguih.com:443/api/epayment/create_transaction/" {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("Failed to read request body:", err)
			}
			// Print the request body
			fmt.Println("Request Body:", string(bodyBytes))
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			return resp

		}
		return resp
	})
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if req.URL.String() == "https://api.naviguih.com:443/api/epayment/create_transaction/" {
			// Read the request body
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				log.Println("Failed to read request body:", err)
				return req, nil
			}
			// Print the request body
			fmt.Println("Request Body:", string(bodyBytes))

			// Reset the request body to its original state
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		return req, nil
	})

	// Set the custom CA certificate
	goproxy.GoproxyCa = caCert

	log.Println("Proxy server is running on :8000")
	log.Fatal(http.ListenAndServe(":8000", proxy))
}
