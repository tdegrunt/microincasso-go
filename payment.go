package microincasso

import "fmt"

type Payment struct {
	Enduser string `xml:"ENDUSER"`
	Amount  int    `xml:"AMOUNT"`
}

func (p *Payment) New() (*Response, error) {
	req := NewPaymentRequest(p)

	be := GetBackend()
	return be.Call("POST", "/RegisteredPayment/", req)
}

func (p *Payment) getHashables() [][]byte {
	amount := fmt.Sprintf("%d", p.Amount)
	return [][]byte{[]byte(p.Enduser), []byte(amount)}
}
