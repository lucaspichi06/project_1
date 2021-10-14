package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/project_1/cmd/controller"
	"github.com/project_1/cmd/domain"
	"github.com/strech/testify/assert"
)

type produceService struct {
	add    func(produce *domain.Produce) error
	fetch  func() ([]domain.Produce, error)
	delete func(code string) error
}

func (p produceService) Add(produce *domain.Produce) error {
	return p.add(produce)
}

func (p produceService) Fetch() ([]domain.Produce, error) {
	return p.fetch()
}

func (p produceService) Delete(code string) error {
	return p.delete(code)
}

func TestPing(t *testing.T) {
	router := Start()
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fail()
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", string(jsonResp))
}

func TestAdd(t *testing.T) {
	serviceMock := &produceService{
		add: func(produce *domain.Produce) error {
			return nil
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	body := []byte(`{"name":"Lettuce","code":"A12T-4GH7-QPL9-3N4M","unit_price":3.46}`)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/produce", bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Lettuce", responseMap["name"])
}

func TestAdd_BadRequestError(t *testing.T) {
	serviceMock := &produceService{
		add: nil,
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	body := []byte(`{"name":123}`)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/produce", bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, responseMap["message"], "content has invalid format")
}

func TestAdd_InternalServerError(t *testing.T) {
	serviceMock := &produceService{
		add: func(produce *domain.Produce) error {
			return errors.New("test-error")
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	body := []byte(`{"name":"Lettuce","code":"A12T-4GH7-QPL9-3N4M","unit_price":3.46}`)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/produce", bytes.NewBuffer(body))
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, responseMap["message"], "error while adding the produce")
}

func TestFetch(t *testing.T) {
	serviceMock := &produceService{
		fetch: func() ([]domain.Produce, error) {
			return []domain.Produce{
				{
					Name:      "Lettuce",
					Code:      "A12T-4GH7-QPL9-3N4M",
					UnitPrice: 3.46,
				},
			}, nil
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/produce", nil)
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	var responseMap []domain.Produce
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Lettuce", responseMap[0].Name)
}

func TestFetch_InternalServerError(t *testing.T) {
	serviceMock := &produceService{
		fetch: func() ([]domain.Produce, error) {
			return []domain.Produce{}, errors.New("test-error")
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/produce", nil)
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, responseMap["message"], "error while fetching data")
}

func TestFetch_NotFound(t *testing.T) {
	serviceMock := &produceService{
		fetch: func() ([]domain.Produce, error) {
			return []domain.Produce{}, nil
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/produce", nil)
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, responseMap["message"], "data not found")
}

func TestDelete(t *testing.T) {
	serviceMock := &produceService{
		delete: func(code string) error {
			return nil
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/produce/A12T-4GH7-QPL9-3N4M", nil)
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDelete_InternalServerError(t *testing.T) {
	serviceMock := &produceService{
		delete: func(code string) error {
			return errors.New("test-error")
		},
	}

	produceController := controller.NewProduceController(serviceMock)
	router := start(&controllers{
		produce: &produceController,
		status:  controller.NewStatusController(),
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/produce/A12T-4GH7-QPL9-3N4M", nil)
	if err != nil {
		t.Fail()
	}
	router.ServeHTTP(w, req)

	responseMap := make(map[string]interface{})
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fail()
	}
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, responseMap["message"], "error while deleting the produce")
}
