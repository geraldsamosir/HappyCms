package Utils

import (
	"gopkg.in/go-playground/validator.v9"
)

type Validation struct {
}

func (v *Validation) ApiValidationResponse(structModel interface{}) []string {
	var validate = validator.New()
	var errList []string
	ErrBodyContent := validate.Struct(structModel)
	if ErrBodyContent != nil {
		for _, err := range ErrBodyContent.(validator.ValidationErrors) {
			errList = append(errList, "Field : "+err.Field()+" is "+err.Tag())
		}
		return errList
	}
	return nil
}
