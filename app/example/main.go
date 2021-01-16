package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/burhon94/goCaptchaMojo/cmd/captcha"
	"github.com/burhon94/goCaptchaMojo/modeles"
)

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	var postParameters modeles.ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	base64blob, captchaId, value := captcha.GenerateCaptcha(postParameters)

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 200, "data": base64blob, "captchaId": captchaId, "msg": "success", "value": value}
	json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	var postParameters modeles.ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	verifyResult := captcha.VerifyCaptcha(postParameters)

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := modeles.Response{
		StatusCode: 404,
		Message:    "captcha failed",
		Payload:    nil,
	}
	if verifyResult {
		body = modeles.Response{
			StatusCode: 200,
			Message:    "captcha verified",
			Payload:    nil,
		}
	}
	json.NewEncoder(w).Encode(body)
}

var Mux = http.NewServeMux()

//start a net/http server
func main() {
	//serve Vuejs+ElementUI+Axios Web Application
	Mux.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	Mux.HandleFunc("/api/get", generateCaptchaHandler)

	//api for verify captcha
	Mux.HandleFunc("/api/verify", captchaVerifyHandle)

	addr := "0.0.0.0:8777"
	fmt.Println(addr)
	if err := http.ListenAndServe(addr, Mux); err != nil {
		log.Fatal(err)
	}
}
