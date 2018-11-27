package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/*
{
    "clani": ["jm1234", "mn3322"],
    "opis_projekta": "Nas projekt implementira aplikacijo za oddajo nepremicnin.",
    "mikrostoritve": ["http://35.189.96.118:8081/v1/orders", "http://35.197.209.159:8080/v1/customers/"],
    "github": ["https://github.com/jmezna/rso-customers", "https://github.com/jmezna/rso-orders"],
    "travis": ["https://travis-ci.org/jmezna/rso-customers", "https://travis-ci.org/jmezna/rso-orders"],
    "dockerhub": ["https://hub.docker.com/r/jmezna/rso-customers/", "https://hub.docker.com/r/jmezna/rso-orders/"]
}
*/

type infoResponse struct {
	Clani         []string `json:"clani"`
	OpisProjekta  string   `json:"opis_projekta"`
	Mikrostoritve []string `json:"mikrostoritve"`
	GitHub        []string `json:"github"`
	Travis        []string `json:"travis"`
	DockerHub     []string `json:"dockerhub"`
}

func main() {
	inf := &infoResponse{
		Clani:        []string{"zk6063", "uh8243"},
		OpisProjekta: "Najin projekt implementira aplikacijo za izposojo koles",
		Mikrostoritve: []string{
			"",
		},
		GitHub: []string{

		},
		Travis: []string{

		},
		DockerHub: []string{

		},
	}
	s, _ := json.Marshal(inf)

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(s)
	})

	host := stringOrDefault("0.0.0.0", "SERVER_HOST")
	port := stringOrDefault("80", "SERVER_PORT")
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

func stringOrDefault(def, key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return v
}
