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

func resInformationUser(w http.ResponseWriter, r *http.Request){
		user := UserInformation{"To Vinh Tuan", "tuantv", "male", "11/02/1998"}
		userJson, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(userJson)
}
func authenHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		md := r.Header.Get("Authorization")
		log.Println(md)
		if md != "Bearer AKcqHRCTHaBLnznmH3fw6bRSMBSZpa9tAngkKnGydBmST5XFGpxzgsGMuT3z7QsZ"{
			log.Println("nhay vao day")
			return
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	port := "3333"

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()
	r.Use(authenHeader)
	r.Get("/api/me", resInformationUser)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

