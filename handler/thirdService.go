package handler

import (
	"github.com/SBEKzy/testTask/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) SendRequestForThirdService(c *gin.Context) {
	var req model.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := req.Check(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.Service.SendRequest(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)

}
