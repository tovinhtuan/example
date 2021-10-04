package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserInformation struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

type FailedRequest struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func resInformationUser(w http.ResponseWriter, r *http.Request) {
	user := UserInformation{"To Vinh Tuan", "tuantv", "male", "11/02/1998"}
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

func authenHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		md := r.Header.Get("Authorization")
		if md != "Bearer AKcqHRCTHaBLnznmH3fw6bRSMBSZpa9tAngkKnGydBmST5XFGpxzgsGMuT3z7QsZ" {
			failReq := FailedRequest{false, "authentication error"}
			fail, err := json.Marshal(failReq)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
				return
			}
			// http.Error(w, "authentication error", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			w.Write(fail)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := "3333"
	log.Printf("Starting up on http://localhost:%s", port)
	r := chi.NewRouter()
	r.Route("/api/me", func(r chi.Router) {
		r.Use(authenHeader)
		r.Get("/", resInformationUser)
	})
	log.Fatal(http.ListenAndServe(":"+port, r))
}
