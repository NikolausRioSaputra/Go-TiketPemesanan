package handlergin

import (
	"Go-TiketPemesanan/internal/domain"
	"Go-TiketPemesanan/internal/usecase"
	// "errors"
	"net/http"
	"strconv"
	"time"

	// "github.com/benebobaa/valo"
	"github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserHandlerInterface interface {
	StoreNewUser(c *gin.Context)
	UserFindById(c *gin.Context)
	UserDeleter(c *gin.Context)
	UserUpdater(c *gin.Context)
	GetAllUser(c *gin.Context)
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

// UserFindById implements UserHandlerInterface.
func (h *UserHandler) UserFindById(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "User id is Required",
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("User ID is Required")
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "Invalid User ID",
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("Invalid User ID")
		return
	}

	user, err := h.UserUsecase.UserFindById(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: err.Error(),
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("err.Error()")
		return
	}

	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success get user by ID",
		Data:    user,
	})
	log.Info().
		Int("http-status-code:", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Get User by ID API-COMPLATED")
}

// DeleteUser implements UserHandlerInterface.
func (h *UserHandler) UserDeleter(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "User ID is required",
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("User ID is required")
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "invalid user id",
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("Invalid user id ")
		return
	}

	err = h.UserUsecase.UserDeleter(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: err.Error(),
		})
		log.Error().
			Int("http-status-code", http.StatusBadRequest).
			TimeDiff("waktu proses:", time.Now(), start).
			Msg(err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success delete user",
	})
	log.Info().
		Int("http-status-code:", http.StatusOK).
		TimeDiff("waktu process:", time.Now(), start).
		Msg("Delete User API-Completed")
}

// StoreNewUser implements UserHandlerInterface.
func (h *UserHandler) StoreNewUser(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	var user domain.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "invalid request body",
			Errors:  err.Error(),
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("invalid request body")
		return
	}

	if err := valo.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: "validation error",
			Errors:  err,
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg("validation error")
		return
	}

	user, err := h.UserUsecase.UserSaver(ctx, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: err.Error(),
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg(err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success create user",
		Data:    user,
	})
	log.Info().
		Int("http-status-code:", http.StatusOK).
		TimeDiff("waktu process:", time.Now(), start).
		Msg("Create User API-Completed")

}

// UpdateUser implements UserHandlerInterface.
func (h *UserHandler) UserUpdater(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Message: "User ID is required",
		})
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("User ID is required")
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Message: "Invalid user ID",
		})
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("Invalid user ID")
		return
	}

	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Message: "Invalid request body",
		})
		log.Error().
			Int("http.status.code", http.StatusBadRequest).
			TimeDiff("waktu process", time.Now(), start).
			Msg("Invalid request body")
		return
	}

	user.ID = id
	updatedUser, err := h.UserUsecase.UserUpdater(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Message: err.Error(),
		})
		log.Error().
			Int("http.status.code", http.StatusInternalServerError).
			TimeDiff("waktu process", time.Now(), start).
			Msg(err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{
		Message: "Success update user",
		Data:    updatedUser,
	})
	log.Info().
		Int("http.status.code", http.StatusOK).
		TimeDiff("waktu process", time.Now(), start).
		Msg("Update User API-Completed")
}

// GetAllUser implements UserHandlerInterface.
func (h *UserHandler) GetAllUser(c *gin.Context) {
	ctx := c.Request.Context()
	start := time.Now()

	users, err := h.UserUsecase.GetAllUser(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseMasage{
			Message: err.Error(),
		})
		log.Error().
			Int("http-status-code:", http.StatusBadRequest).
			TimeDiff("waktu process:", time.Now(), start).
			Msg(err.Error())
		return
	}

	c.JSON(http.StatusOK, ResponseMasage{
		Message: "Success get all user",
		Data:    users,
	})
	log.Info().
		Int("http-status-code:", http.StatusOK).
		TimeDiff("waktu process:", time.Now(), start).
		Msg("Get All User API-Completed")
}
