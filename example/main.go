package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-recaptcha/recaptcha"
)

var html = []byte(`<!DOCTYPE html>
<html>
	<head>
		<title>gopkg.in/recaptcha example</title>
		<script src='https://www.google.com/recaptcha/api.js'></script>
	</head>
	<body>
		<form action="/verify" >
			<div class="g-recaptcha" data-sitekey="6Le9ywoTAAAAAD6fXMq-tJKObM2GVc7y9rOWcDB_"></div>
			<button type="submit" > Check </button>
		</form>
	</body>
</html>`)

var (
	captchaSecret = `6Le9ywoTAAAAAF2DV_ocsI5y8UkAjGQIY2Kr7Gic`
	captcha       = recaptcha.New(captchaSecret)
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/verify", verifyHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	captchaResponse := r.FormValue("g-recaptcha-response")
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)

	valid, err := captcha.Verify(captchaResponse, remoteIP)
	if err != nil {
		fmt.Fprintf(w, "error in verification: %v\n", err)
		return
	}
	if valid {
		fmt.Fprintf(w, "valid captcha validation")
		return
	}
	fmt.Fprintf(w, "invalid catpcha validation")
}
