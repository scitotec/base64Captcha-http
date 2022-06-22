package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mojocn/base64Captcha"
	"github.com/scitotec/base64captcha-http/handlers"
)

var store = base64Captcha.DefaultMemStore

func main() {
	listenAddress := os.Getenv("APP_LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":8777"
	}

	http.HandleFunc("/", handlers.HijackNotFound(handlers.NewHelloHandler()))

	//api for create captcha
	http.HandleFunc("/v1/new/audio", handlers.OnlyPost(handlers.NewAudioHandler(store)))
	http.HandleFunc("/v1/new/chinese", handlers.OnlyPost(handlers.NewChineseHandler(store)))
	http.HandleFunc("/v1/new/digit", handlers.OnlyPost(handlers.NewDigitHandler(store)))
	http.HandleFunc("/v1/new/language", handlers.OnlyPost(handlers.NewLanguageHandler(store)))
	http.HandleFunc("/v1/new/math", handlers.OnlyPost(handlers.NewMathHandler(store)))
	http.HandleFunc("/v1/new/string", handlers.OnlyPost(handlers.NewStringHandler(store)))

	//api for verify captcha
	http.HandleFunc("/v1/verify", handlers.OnlyPost(handlers.NewVerifyHandler(store)))

	log.Printf("Server listening on %s\n", listenAddress)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
