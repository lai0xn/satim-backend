package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/laix0n/satim/pkg/utils"
)

type url struct {
	Url string `json:"url"`
}

var (
	ssl bool 
	captcha bool 
	Logo bool
	GreenNumber bool
	CardsValidation bool 
	reqValidation bool
)

func SendUrl (w http.ResponseWriter, r *http.Request) {
	var link url
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	fmt.Println("Received URL:", link.Url)
	err = utils.Checkssl(link.Url)
	if err != nil {
		ssl = false
	} else {
		ssl = true
	}

	
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("finish"))

}