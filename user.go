package microincasso

import "errors"

type UserStatus int

const (
	NoUser                 UserStatus = iota // No Microincasso end-user present
	UnwantedUser                             // Found Microincasso end-user is an unwanted end-user and is blocked for further actions with Microincasso.
	NewUser                                  // Microincasso end-user found but it needs a valid IBAN before any further actions could be done.
	FullUser                                 // Microincasso end-user found and has a valid MSISDN and IBAN, but it is not registered for the current merchant.
	FullRegisteredUser                       // Microincasso end-user found and has a valid MSISDN and IBAN and is registered for the current merchant.
	TemporarilyBlockedUser                   // Microincasso end-user is temporarily blocked due to an unperformed action. (E.g. a one-cent-check is required but not finalized by the end-user self).
)

type User struct {
	//Status    UserStatus
	Reference string `xml:"REFERENCE"`
	Iban      string `xml:"IBAN,omitempty"`
}

// Status of a user
func (u *User) Status() (UserStatus, error) {
	req := newUserRequest(u)

	be := GetBackend()
	resp, err := be.Call("POST", "/EndUserStatus/", req)
	if err != nil {
		return -1, err
	}

	if resp.IsError() {
		return -1, resp.Error()
	}

	return UserStatus(resp.Status), nil
}

// Registers a user
func (u *User) Register() (UserStatus, error) {
	req := newUserRequest(u)

	be := GetBackend()
	resp, err := be.Call("POST", "/EndUserRegistration/", req)
	if err != nil {
		return -1, err
	}

	return UserStatus(resp.Status), nil
}

// Revokes the registration for the user
func (u *User) Revoke() error {
	req := newUserRequest(u)

	be := GetBackend()
	resp, err := be.Call("POST", "/RevokeRegistration/", req)
	if err != nil {
		return err
	}
	if resp.Status != 200 {
		return errors.New("Response status not 200")
	}
	return nil
}

func (u *User) getHashables() [][]byte {
	return [][]byte{[]byte(u.Reference), []byte(u.Iban)}
}

func newUserRequest(u *User) *Request {
	req := &Request{Merchant: Merchant{Username: Username, Password: Password}, User: u}
	req.Merchant.VerificationHash = req.GetHash()
	return req
}
