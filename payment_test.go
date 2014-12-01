package microincasso_test

import (
	. "github.com/boxture/microincasso"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payment", func() {
	BeforeEach(func() {
		Secret = "3F29A724-00E9-410D-90A5-7B4E391DE266"
		Username = "PostNLPakketten02"
		Password = "QhjlqaajKAf5mtVh"

		u := User{Reference: "31641085630", Iban: "NL75RABO0103483640"}
		status, err := u.Register()
		Expect(err).NotTo(HaveOccurred())
		Expect(status).To(Equal(UserStatusFullRegisteredUser))

	})
	AfterEach(func() {
		u := User{Reference: "31641085630"}
		err := u.Revoke()
		Expect(err).NotTo(HaveOccurred())
	})
	It("should allow payment for a registered user", func() {

		p := Payment{Enduser: "31641085630", Amount: 1, Test: 1}
		reference, err := p.New()
		Expect(err).NotTo(HaveOccurred())
		Expect(reference).ToNot(BeNil())

	})

})
