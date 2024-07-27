package handler

import (
	"Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	// "github.com/benebobaa/valo"
	"github.com/benebobaa/valo"
	"github.com/rs/zerolog/log"
)

type UserHandlerInterface interface {
	StoreNewUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetAllUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	UserUsecase usecase.UserUsecaseInterface
}

type ResponseMasage struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) UserHandlerInterface {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

// DeleteUser implements UserHandlerInterface.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Info().
			Int("http.status.code", http.StatusMethodNotAllowed).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	users := r.URL.Query().Get("id")

	if users == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid request payload")
		return
	}

	userID, err := strconv.Atoi(users)
	if err != nil {
		log.Printf("Invalid Id parameter : %v", err)
		http.Error(w, "Invalid id parameter", http.StatusBadGateway)
		return
	}
	_, err = h.UserUsecase.DeleteUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ResponseMasage{
		Message: "Success delete the user",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
}

// GetAllUser implements UserHandlerInterface.
func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Info().
			Int("http.status.code", http.StatusMethodNotAllowed).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	users, err := h.UserUsecase.GetAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ResponseMasage{
		Message: "Success get all user",
		Data:    users,
	})
	if err != nil {
		start := time.Now()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
}

// StoreNewUser implements UserHandlerInterface.
func (h *UserHandler) StoreNewUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Info().
			Int("http.status.code", http.StatusMethodNotAllowed).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var users domain.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	err = valo.Validate(users)
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseMasage{
			Message: "invalid request body",
			Errors: err.Error(),
			Data:    users,
		})
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	users, err = h.UserUsecase.UserSaver(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ResponseMasage{
		Message: "Success add User",
		Data:    users,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		valo.Validate(users)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}
}

// UpdateUser implements UserHandlerInterface.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("Method not allowed")
		return
	}

	var users domain.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("invalid method")
		return
	}

	_, err = h.UserUsecase.UpdateUser(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Info().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ResponseMasage{
		Message: "Success Update User",
		Data:    users,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

}
