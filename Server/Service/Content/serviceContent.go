package Content

import (
	"encoding/json"
	"net/http"

	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

var Response Utils.Response
var validate *validator.Validate
var Validation Utils.Validation

type ServiceContent struct {
	Message    string   `json:"message" validate:"required"`
	Status     int      `json:"status" validate:"required,number"`
	Validation []string `json:validation`
}

func (SC *ServiceContent) WelcomeContent(res http.ResponseWriter, req *http.Request) {
	SC.Message = "Welcome to content"
	SC.Status = http.StatusOK
	Response.ResponseJSON(SC, res, http.StatusOK)
}

func (SC *ServiceContent) Create(res http.ResponseWriter, req *http.Request) {
	var content DataContent
	Bodycontent := &ServiceContent{}
	if err := json.NewDecoder(req.Body).Decode(Bodycontent); err != nil {
		SC.Validation = nil
		SC.Message = "Data validation false"
		SC.Status = http.StatusBadRequest
		SC.Validation = append(SC.Validation, err.Error())
		errResp, _ := json.Marshal(SC)
		logrus.WithFields(logrus.Fields{
			"responsetoFrontend": errResp,
		}).Info("validation request in create content false")
		Response.ResponseJSON(SC, res, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	ErrBodyContent := Validation.ApiValidationResponse(Bodycontent)
	if ErrBodyContent != nil {
		SC.Message = "Validaition false"
		SC.Validation = ErrBodyContent
		SC.Status = http.StatusBadRequest
		Response.ResponseJSON(SC, res, http.StatusBadRequest)
		return
	}
	modelContent := content.Models()
	content.Title = "CIE JONES BANGET"
	content.Author = "GERALD"
	content.Categories = "JOKE"
	content.Content = "Contentnya"
	err := modelContent.Insert(content)
	if err != nil {
		panic(err)
	}
	SC.Message = "Content created"
	SC.Status = http.StatusCreated
	Response.ResponseJSON(SC, res, http.StatusCreated)

}
