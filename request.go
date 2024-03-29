package microincasso

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
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

// Returns the verification hash for this request
func (r *Request) GetHash() string {

	hashables := [][]byte{[]byte(r.Merchant.Username), []byte(r.Merchant.Password)}

	if r.User != nil {
		hashables = ConcatenateBytes(hashables, r.User.getHashables())
	}

	if r.Payment != nil {
		hashables = ConcatenateBytes(hashables, r.Payment.getHashables())
	}

	hashables = append(hashables, []byte(Secret))

	var h hash.Hash = sha256.New()
	for _, v := range hashables {
		if len(v) > 0 {
			if debug {
				fmt.Printf("hashable: %+s\n", v)
			}
			h.Write(v)
		}
	}
	result := hex.EncodeToString(h.Sum([]byte{}))

	return result
}
