package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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
			"http://api.bicycle.hercog.si/users",
			"http://api.bicycle.hercog.si/billing",
		},
		GitHub: []string{
			"https://github.com/RSO-Bicycle/users",
			"https://github.com/RSO-Bicycle/billing",
		},
		Travis: []string{
			"https://travis-ci.org/RSO-Bicycle/users",
			"https://travis-ci.org/RSO-Bicycle/billing",
		},
		DockerHub: []string{
			"https://console.cloud.google.com/gcr/images/rso-bicycle/EU/users?project=rso-bicycle",
			"https://console.cloud.google.com/gcr/images/rso-bicycle/EU/billing?project=rso-bicycle",
		},
	}
	s, _ := json.Marshal(inf)

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(s)
	})

	host := stringOrDefault("0.0.0.0", "SERVER_HOST")
	port := stringOrDefault("8080", "SERVER_PORT")
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
