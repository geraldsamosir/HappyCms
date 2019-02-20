package Content

import (
	"encoding/json"
	"net/http"

	"github.com/geraldsamosir/geraldsamosir/HappyCms/Server/Utils"
)

var Response Utils.Response

type ServiceContent struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (SC *ServiceContent) WelcomeContent(res http.ResponseWriter, req *http.Request) {
	SC.Message = "Welcome to content"
	SC.Status = http.StatusOK
	Response.ResponseJSON(SC, res, http.StatusOK)
}

func (SC *ServiceContent) Create(res http.ResponseWriter, req *http.Request) {
	var content DataContent
	scv := &ServiceContent{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(scv); err != nil {
		panic(err)
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
