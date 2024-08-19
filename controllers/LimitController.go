package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-test/limits"
	"net/http"
	"strconv"
)

type limitHandler struct {
	limitServices limits.Service
}

func NewLimitHandler(service limits.Service) *limitHandler {

	return &limitHandler{service}
}

func (h *limitHandler) LimitCreate(c *gin.Context) {
	var request limits.LimitRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_ = c.ShouldBindJSON(&request)

	customerLimit, err := h.limitServices.Create(id, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(201, gin.H{
		"data": customerLimit,
	})
}
