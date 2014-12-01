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

type MIDateTime struct {
	time.Time
}

func (c *MIDateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortForm = "2006-01-02T15:04:05" // 2014-11-28T13:38:45.8 date format
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortForm, v)
	if err != nil {
		return nil
	}
	*c = MIDateTime{parse}
	return nil
}

type Response struct {
	XMLName         xml.Name    `xml:"MIRESPONSE"`
	Status          int         `xml:"STATUS"`
	User            *User       `xml:"ENDUSER"`
	StartDate       *MIDateTime `xml:"STARTDATE"`
	RegistrationUrl *string     `xml:"REGISTRATIONURL"`
	Reference       *string     `xml:"REFERENCE"`

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
