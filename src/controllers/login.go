package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	redirectURL := "http://127.0.0.1:8000/index"
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print("Error body..................")
		fmt.Printf("error", err.Error())
	}
	str1 := string(bodyRequest)
	str2 := strings.Split(str1, "\n")
	var mstr map[string]string
	mstr = make(map[string]string)

	for _, str3 := range str2 {
		str4 := strings.Split(str3, "=")
		key := strings.ReplaceAll(str4[0], " ", "")
		value := strings.ReplaceAll(str4[1], " ", "")
		mstr[key] = value
	}

	if strings.TrimSpace(mstr["Email"]) == "rafael.tomelin@gmail.com" && strings.TrimSpace(mstr["Tenant"]) == "mothers" && strings.TrimSpace(mstr["Password"]) == "123" {
		fmt.Print("Puts")
		http.Redirect(w, r, redirectURL, 302)
	} else {

		fmt.Print("Ufaa")
		http.Redirect(w, r, redirectURL, 302)
	}
}
