package microincasso

import (
	"encoding/hex"
	"encoding/xml"
	"hash"

	"crypto/sha256"
)

type Merchant struct {
	Username         string `xml:"USERNAME"`
	Password         string `xml:"PASSWORD"`
	VerificationHash string `xml:"VERIFICATIONHASH"`
}

type Request struct {
	XMLName  xml.Name `xml:"MIREQUEST"`
	Merchant Merchant `xml:"MERCHANT"`
	User     *User    `xml:"ENDUSER"`
	Payment  *Payment `xml:"PAYMENT"`
}

func (r *Request) getHash() string {

	hashables := [][]byte{[]byte(r.Merchant.Username), []byte(r.Merchant.Password)}

	if r.User != nil {
		hashables = ConcatenateBytes(hashables, r.User.getHashables())
	}

	hashables = append(hashables, []byte(Secret))

	var h hash.Hash = sha256.New()
	for _, v := range hashables {
		h.Write(v)
	}
	result := hex.EncodeToString(h.Sum([]byte{}))

	return result
}
