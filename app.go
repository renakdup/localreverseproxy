package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
)

func main() {
	proxyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var proxyUrl string

		services := getMapped()

		fmt.Println(fmt.Sprintf("Requested `%s` host with method `%s` and URI `%s`", r.Host, r.Method, r.RequestURI))

		for servicePattern, pUrl := range services {
			pattern := fmt.Sprintf(`^%s$`, servicePattern)
			match, err := regexp.MatchString(pattern, r.Host)
			if err != nil {
				fmt.Println(fmt.Sprintf("RegExp returned error: %s", err))
				continue
			}

			if match == true {
				proxyUrl = pUrl
				break
			}
		}

		if proxyUrl == "" && r.Host == "localhost" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Server proxy started successfully!")
			return
		}

		if proxyUrl == "" {
			http.Error(w, "No match of service was found in services.json", http.StatusInternalServerError)
			return
		}

		parsedProxyUrl, err := url.Parse(proxyUrl)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error parsing of service URL %s", proxyUrl),
				http.StatusInternalServerError,
			)
			return
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(parsedProxyUrl)
		reverseProxy.ServeHTTP(w, r)
	})

	fmt.Println("Reverse proxy is started.")
	err := http.ListenAndServe(":80", proxyHandler)
	if err != nil {
		panic(err)
	}
}

func getMapped() map[string]string {
	data, err := os.ReadFile("services.json") // For read access.
	if err != nil {
		panic(err)
	}

	var parsedHosts map[string]string
	if err := json.Unmarshal(data, &parsedHosts); err != nil {
		panic(err)
	}

	return parsedHosts
}
