package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project_1/cmd/errors"
	"github.com/project_1/cmd/utils"
)

func validateId() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("id")

		isValid, err := utils.ValidateCode(code)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.NewInternalServerAppError("internal server error", err))
			c.Abort()
			return
		}

		if isValid {
			c.Set("code", code)
			c.Next()
			return
		}

		c.JSON(http.StatusBadRequest, errors.NewBadRequestAppError("the provided code is invalid"))
		c.Abort()
	}
}
