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

		hosts := getMapped()

		for domainPattern, pUrl := range hosts {
			pattern := fmt.Sprintf(`^%s$`, domainPattern)
			match, err := regexp.MatchString(pattern, r.Host)
			if err != nil {
				//TODO:: add logging
				continue
			}

			if match == true {
				proxyUrl = pUrl
				break
			}
		}

		// reverseProxy url
		//fmt.Println(proxyUrl)

		parsedProxyUrl, err := url.Parse(proxyUrl)
		if err != nil {
			http.Error(w, "Error parsing backend URL", http.StatusInternalServerError)
			return
		}

		reverseProxy := httputil.NewSingleHostReverseProxy(parsedProxyUrl)
		reverseProxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":80", proxyHandler)
}

func getMapped() map[string]string {
	data, err := os.ReadFile("hosts.json") // For read access.
	if err != nil {
		panic(err)
	}

	var parsedHosts map[string]string
	if err := json.Unmarshal(data, &parsedHosts); err != nil {
		panic(err)
	}

	return parsedHosts
}
