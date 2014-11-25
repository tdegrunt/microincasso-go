package microincasso

type User struct {
	Reference string `xml:"REFERENCE"`
	Iban      string `xml:"IBAN,omitempty"`
}

func (u *User) Register() (*Response, error) {
	req := NewUserRequest(u)

	be := GetBackend()
	return be.Call("POST", "/EndUserRegistration/", req)
}

func (u *User) getHashables() [][]byte {
	return [][]byte{[]byte(u.Reference), []byte(u.Iban)}
}
