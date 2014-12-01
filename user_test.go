package microincasso_test

import (
	. "github.com/boxture/microincasso"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	BeforeEach(func() {
		Secret = "3F29A724-00E9-410D-90A5-7B4E391DE266"
		Username = "PostNLPakketten02"
		Password = "QhjlqaajKAf5mtVh"
	})
	Describe("Registration", func() {
		It("should be an new user when not registered", func() {
			u := User{Reference: "31663085410"} // Not my phone number!
			status, err := u.Status()
			Expect(err).NotTo(HaveOccurred())
			Expect(status).To(Equal(UserStatusNewUser))
		})
		It("should register", func() {
			u := User{Reference: "31641085630", Iban: "NL75RABO0103483640"}
			status, err := u.Register()
			Expect(err).NotTo(HaveOccurred())
			Expect(status).To(Equal(UserStatusFullRegisteredUser))
		})
		It("should be an full registered user when registered", func() {
			u := User{Reference: "31641085630"}
			status, err := u.Status()
			Expect(err).NotTo(HaveOccurred())
			Expect(status).To(Equal(UserStatusFullRegisteredUser))
		})
	})
	Describe("Revoke", func() {
		It("should be able to revoke when registered", func() {
			u := User{Reference: "31641085630"}
			err := u.Revoke()
			Expect(err).NotTo(HaveOccurred())
		})
		It("should be an full user (but not registered for this merchant) when revoked", func() {
			u := User{Reference: "31641085630"}
			status, err := u.Status()
			Expect(err).NotTo(HaveOccurred())
			Expect(status).To(Equal(UserStatusFullUser))
		})
	})

})
