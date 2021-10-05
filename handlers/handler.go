package handlers

import (
	"context"
	"encoding/json"
	"ex1/repositories"
	"log"
	"net/http"
	// "strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FailedRequest struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type UserHandler struct {
	Repository repositories.Repository
}

func NewUserHandler(Repository repositories.Repository) (*UserHandler, error) {
	return &UserHandler{
		Repository: Repository,
	}, nil
}
func (h *UserHandler) ReadTokenByToken(ctx context.Context, token string) (bool, error) {
	_, err := h.Repository.ReadTokenByToken(ctx, token)
	if err != nil {
		return false, status.Error(codes.NotFound, "Token is not found")
	}
	return true, nil
}
func (h *UserHandler) AuthenHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md := r.Header.Get("Authorization")
		// space := strings.Index(md," ")
		token := md[7:]
		checkToken, _ := h.ReadTokenByToken(context.Background(), token)
		// if md != "Bearer AKcqHRCTHaBLnznmH3fw6bRSMBSZpa9tAngkKnGydBmST5XFGpxzgsGMuT3z7QsZ" {
		if !checkToken {
			failReq := FailedRequest{false, "authentication error"}
			fail, err := json.Marshal(failReq)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error happened in JSON marshal. Err: %s", err)
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
func (h *UserHandler) ResInformation(w http.ResponseWriter, r *http.Request){
	md := r.Header.Get("Authorization")
	token := md[7:]
	user, err := h.Repository.ReadUserByToken(context.Background(),token)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}	
	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(userJson)
}