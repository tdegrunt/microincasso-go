package microincasso

import "fmt"

type Payment struct {
	Enduser string `xml:"ENDUSER"`
	Amount  int    `xml:"AMOUNT"`
}

// New payment for a registered end-user
func (p *Payment) New() (*Response, error) {
	req := newPaymentRequest(p)

	be := GetBackend()
	return be.Call("POST", "/RegisteredPayment/", req)
}

func (p *Payment) getHashables() [][]byte {
	amount := fmt.Sprintf("%d", p.Amount)
	return [][]byte{[]byte(p.Enduser), []byte(amount)}
}

func newPaymentRequest(p *Payment) *Request {
	req := &Request{Merchant: Merchant{Username: Username, Password: Password}, Payment: p}
	req.Merchant.VerificationHash = req.GetHash()
	return req
}
