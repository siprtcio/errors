package errors

import (
	"bytes"
	"text/template"
)

type B2Code uint16

// General Codes.
const (
	ResourceNotFound B2Code = 20404
)

var ErrorTemplates = map[B2Code]ApiError{
	ResourceNotFound: {
		Code:     20404,
		Message:  "The requested resource {{.Resource}} was not found",
		MoreInfo: "https://www.twilio.com/docs/errors/20404",
		Status:   404,
	},
}

func parseErrorField(tmplStr string, errInterface interface{}) (string, error) {
	var err error
	// with name passed as argument
	validationTmpl := template.New("SipRtcValidation")
	// "Parse" parses a string into a template
	validationTmpl, err = validationTmpl.Parse(tmplStr)
	if err != nil {
		return tmplStr, err
	}
	var tmplBytes bytes.Buffer

	if err = validationTmpl.Execute(&tmplBytes, errInterface); err != nil {
		return tmplStr, err
	}

	return tmplBytes.String(), nil
}