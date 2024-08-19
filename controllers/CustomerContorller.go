package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-test/customers"
	"net/http"
)

type customerhandler struct {
	customerService customers.ServiceInterface
}

func NewCustomerHandler(service customers.ServiceInterface) *customerhandler {
	return &customerhandler{service}
}

func (h *customerhandler) CustomerCreate(c *gin.Context) {
	var request customers.CustomerRequest

	_ = c.ShouldBindJSON(&request)

	customer, err := h.customerService.Create(request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(201, gin.H{
		"data": customer,
	})
}
