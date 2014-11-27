package microincasso

import (
	"encoding/xml"
	"fmt"
	"time"
)

type MicroIncassoError struct {
	ErrorCode        int
	ErrorDescription string
}

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

func (r *Response) Error() error {
	return MicroIncassoError{ErrorCode: r.ErrorCode, ErrorDescription: r.ErrorDescription}
}

func (e MicroIncassoError) Error() string {
	return fmt.Sprintf("%d: %s", e.ErrorCode, e.ErrorDescription)
}
