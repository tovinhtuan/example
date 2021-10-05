package main

import (
	// "encoding/json"
	"ex1/handlers"
	"ex1/repositories"
	"log"
	"net/http"
	// "strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type UserInformation struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func main() {
	port := "3333"
	log.Printf("Starting up on http://localhost:%s", port)
	r := chi.NewRouter()
	//create table
	repository, err := repositories.NewDBManager()
	if err != nil {
		log.Printf("Error happened in database server marshal. Err: %s", err)
		return
	}
	h, err1 := handlers.NewUserHandler(repository)
	if err1 != nil {
		log.Printf("Error happened in database server marshal. Err: %s", err)
		return
	}
	//
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Route("/api/me", func(r chi.Router) {
		r.Use(h.AuthenHeader)
		r.Get("/",h.ResInformation)
	})
	http.ListenAndServe(":"+port, r)
}

// func resInformationUser(w http.ResponseWriter, r *http.Request) {
// 	user := UserInformation{"To Vinh Tuan", "tuantv", "male", "11/02/1998"}
// 	userJson, err := json.Marshal(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("Error happened in JSON marshal. Err: %s", err)
// 		return
// 	}
// 	w.Write(userJson)
// }

// func authenHeader(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		md := r.Header.Get("Authorization")
// 		// space := strings.Index(md," ")
// 		// token := md[7:]
// 		if md != "Bearer AKcqHRCTHaBLnznmH3fw6bRSMBSZpa9tAngkKnGydBmST5XFGpxzgsGMuT3z7QsZ" {
// 			failReq := FailedRequest{false, "authentication error"}
// 			fail, err := json.Marshal(failReq)
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 				log.Printf("Error happened in JSON marshal. Err: %s", err)
// 				return
// 			}
// 			// http.Error(w, "authentication error", http.StatusForbidden)
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Write(fail)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
