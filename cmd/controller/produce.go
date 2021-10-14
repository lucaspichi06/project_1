package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project_1/cmd/domain"
	"github.com/project_1/cmd/errors"
)

type produceService interface {
	Add(produce *domain.Produce) error
	Fetch() ([]domain.Produce, error)
	Delete(code string) error
}

type produce struct {
	sProduce produceService
}

func NewProduceController(sProduce produceService) produce {
	return produce{
		sProduce: sProduce,
	}
}

func (p *produce) Add(c *gin.Context) {
	body := &domain.Produce{}
	if err := c.BindJSON(body); err != nil {
		log.Println("[controller.produce.add] content has invalid format", err)
		c.JSON(http.StatusBadRequest, errors.NewBadRequestAppError(fmt.Sprintf("content has invalid format: %s", err)))
		return
	}

	err := p.sProduce.Add(body)
	if err != nil {
		log.Println("[controller.produce.add] error while adding the produce", err)
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerAppError("error while adding the produce", err))
		return
	}

	c.JSON(http.StatusCreated, body)
}

func (p *produce) Fetch(c *gin.Context) {
	resp, err := p.sProduce.Fetch()
	if err != nil {
		log.Println("[controller.produce.fetch] error while fetching data", err)
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerAppError("error while fetching data", err))
		return
	}

	if len(resp) == 0 {
		log.Println("[controller.produce.fetch] data not found")
		c.JSON(http.StatusNotFound, errors.NewStatusNotFoundAppError("data not found"))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (p *produce) Delete(c *gin.Context) {
	err := p.sProduce.Delete(c.GetString("code"))
	if err != nil {
		log.Println("[controller.produce.delete] error while deleting the produce", err)
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerAppError("error while deleting the produce", err))
		return
	}

	c.JSON(http.StatusOK, nil)
}
