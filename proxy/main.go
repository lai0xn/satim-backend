package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/elazarl/goproxy"
	"github.com/lai0xn/satim-proxy/redis"
	"github.com/lai0xn/satim-proxy/tester"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	caCert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatal("Failed to load CA certificate:", err)
	}
	redis.NewClient()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if strings.HasPrefix("https://challenges.cloudflare.com", req.URL.String()) {
			log.Info("Captcha Verification Passed")
		}
		return req, nil
	})

	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Warn("Failed to read response body:", err)
			return resp
		}

		// Restore the response body for further processing
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Check for the presence of "actionCode"
		if strings.Contains(string(bodyBytes), `"actionCode":`) {
			var data map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &data); err != nil {
				log.Warn("Failed to unmarshal JSON:", err)
			} else {
				tester.TestFraud("naviguih.com:amount", data)
				log.Info("Parsed Data:", data)
			}
		}

		// Restore the response body again for the client
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		return resp
	})

	// Set the custom CA certificate
	goproxy.GoproxyCa = caCert

	log.Info("Proxy server is running on :8000")
	log.Fatal(http.ListenAndServe(":8000", proxy))
}
