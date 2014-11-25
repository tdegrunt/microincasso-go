package microincasso

import (
	"encoding/xml"
	"time"
)

type Response struct {
	XMLName         xml.Name   `xml:"MIRESPONSE"`
	Status          int        `xml:"STATUS"`
	User            *User      `xml:"ENDUSER"`
	StartDate       *time.Time `xml:"STARTDATE"`
	RegistrationUrl *string    `xml:"REGISTRATIONURL"`
	Reference       *string    `xml:"REFERENCE"`

	ErrorCode        int    `xml:"ERRORCODE"`
	ErrorDescription string `xml:"ERRORDESCRIPTION"`
}

func (r *Response) IsError() bool {
	return r.ErrorCode > 0
}
