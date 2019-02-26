package Content

import (
	"encoding/json"
	"net/http"

	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Utils"
	"github.com/gorilla/schema"
	"gopkg.in/go-playground/validator.v9"
)

var Response Utils.Response
var content DataContent
var validate *validator.Validate
var validationType Utils.ValidationType
var Validation Utils.Validation
var Logs Utils.Log
var Schema = schema.NewDecoder()

type ServiceContent struct {
	Utils.Response
}

func (SC *ServiceContent) WelcomeContent(res http.ResponseWriter, req *http.Request) {
	SC.Message = "Welcome to content"
	SC.Status = http.StatusOK
	Response.ResponseJSON(SC, res, http.StatusOK)
}

func (SC *ServiceContent) Create(res http.ResponseWriter, req *http.Request) {
	Bodycontent := &DataContent{}
	SC.Validation = nil
	if err := json.NewDecoder(req.Body).Decode(Bodycontent); err != nil {
		SC.Message = "Data validation false"
		SC.Status = http.StatusBadRequest

		// handle validation
		validationType.ValidationBody = append(validationType.ValidationBody, err.Error())
		SC.Validation = validationType
		Logs.Logging(SC, "validation request in create content false")
		Response.ResponseJSON(SC, res, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	ErrBodyContent := Validation.ApiValidationResponse(Bodycontent)
	if ErrBodyContent != nil {
		SC.Message = "Validaition false"
		validationType.ValidationBody = ErrBodyContent
		SC.Validation = validationType
		SC.Status = http.StatusBadRequest
		Logs.Logging(SC, "validation request in create content false")
		Response.ResponseJSON(SC, res, http.StatusBadRequest)
		return
	}
	modelContent := content.Models()
	err := modelContent.Insert(content)
	if err != nil {
		panic(err)
	}
	SC.Message = "Content created"
	SC.Status = http.StatusCreated
	Response.ResponseJSON(SC, res, http.StatusCreated)

}
func (SC *ServiceContent) Update(res http.ResponseWriter, req *http.Request) {
	Bodycontent := &DataContent{}
	var filterquery QueryUrlForUpdate
	SC.Validation = nil
	if err := Schema.Decode(&filterquery, req.URL.Query()); err != nil {
		SC.Message = "Data validation false"
		SC.Status = http.StatusBadRequest

		// handle validation
		validationType.ValidationQueryUrl = nil
		validationType.ValidationQueryUrl = append(validationType.ValidationQueryUrl, err.Error())
		SC.Validation = validationType
		Logs.Logging(SC, "validation request in update content false")
		Response.ResponseJSON(SC, res, http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(req.Body).Decode(Bodycontent); err != nil {
		SC.Message = "Data Validation false"
		SC.Status = http.StatusBadRequest
		validationType.ValidationBody = nil
		validationType.ValidationBody = append(validationType.ValidationBody, err.Error())
		SC.Validation = validationType
		Response.ResponseJSON(SC, res, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	SC.Status = 200
	SC.Message = "ini untuk update"
	SC.Data = nil
	Response.ResponseJSON(SC, res, http.StatusOK)
	return
}
