package microincasso

import "fmt"

type Payment struct {
	Enduser    string `xml:"ENDUSER"`
	Amount     uint64 `xml:"AMOUNT"`
	Test       int    `xml:"TEST"`
	SmsMessage string `xml:"SMSMESSAGE"`
}

// New payment for a registered end-user
func (p *Payment) New() (*string, error) {
	req := newPaymentRequest(p)

	be := GetBackend()
	resp, err := be.Call("POST", "/RegisteredPayment/", req)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, resp.Error()
	}

	return resp.Reference, nil
}

func (p *Payment) getHashables() [][]byte {
	amount := fmt.Sprintf("%d", p.Amount)
	test := fmt.Sprintf("%d", p.Test)
	return [][]byte{[]byte(p.Enduser), []byte(amount), []byte(test), []byte(p.SmsMessage)}
}

func newPaymentRequest(p *Payment) *Request {
	req := &Request{Merchant: Merchant{Username: Username, Password: Password}, Payment: p}
	req.Merchant.VerificationHash = req.GetHash()
	return req
}
