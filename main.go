// httpserver.go
package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Redirection struct {
	Path string
	URL  string
}

type Router struct {
	Host         string
	Port         int
	Redirections []Redirection
}

const DefaultPort = 8888
const DefaultHost = "0.0.0.0"

var logFlag = flag.String("log", "", "Define log file")
var routerFilePath = flag.String("router", "router.json", "Define router file")
var sslKeyFile = flag.String("sslKey", "", "Path to key file")
var sslCertFile = flag.String("sslCert", "", "Path to certificate file")

var logger *log.Logger

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s %s%s", r.RemoteAddr, r.Method, r.Host, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if *logFlag == "" {
		// use default logger
		logger = log.New(os.Stderr, "", log.LstdFlags)
		log.Println("Log to stderr")
	} else {
		logFile, err := os.Create(*logFlag)
		if err != nil {
			log.Println(err)
			return
		}
		defer logFile.Close()
		log.Printf("Log to %s", *logFlag)
		logger = log.New(logFile, "", log.LstdFlags)
	}
	var router Router
	router.Port = DefaultPort
	router.Host = DefaultHost
	routerFile, err := os.Open(*routerFilePath)
	if err != nil {
		log.Println(err)
		router.Redirections = []Redirection{{Path: "/", URL: "."}}
	} else {
		defer routerFile.Close()
		json.NewDecoder(routerFile).Decode(&router)
	}

	routerStr, _ := json.MarshalIndent(router, "", "  ")
	logger.Printf("Router: %s", routerStr)

	for _, redirection := range router.Redirections {
		if strings.HasPrefix(redirection.URL, "http") {
			redirectUrl, err := url.Parse(redirection.URL)
			if err != nil {
				log.Println(err)
				return
			}
			http.Handle(redirection.Path, http.StripPrefix(redirection.Path, httputil.NewSingleHostReverseProxy(redirectUrl)))
		} else {
			http.Handle(redirection.Path, http.StripPrefix(redirection.Path, http.FileServer(http.Dir(redirection.URL))))
		}

	}

	if *sslKeyFile != "" && *sslCertFile != "" {
		logger.Printf("Listen on port https://%s:%d", router.Host, router.Port)
		panic(http.ListenAndServeTLS(router.Host+":"+strconv.Itoa(router.Port), *sslCertFile, *sslKeyFile, Log(http.DefaultServeMux)))
	} else {
		logger.Printf("Listen on port http://%s:%d", router.Host, router.Port)
		panic(http.ListenAndServe(router.Host+":"+strconv.Itoa(router.Port), Log(http.DefaultServeMux)))
	}

}
