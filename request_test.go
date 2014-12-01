package microincasso_test

import (
	. "github.com/boxture/microincasso"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	var _ = Describe("Hash", func() {
		It("should compute", func() {
			Secret = "SECRETKEY"
			Username = "Username"
			Password = "Password"
			req := &Request{Merchant: Merchant{Username: Username, Password: Password}, User: &User{Reference: "0612345678"}}
			hash := req.GetHash()
			Expect(hash).To(Equal("30b212ec8c641139839683b621f49c2d9e8d507852911f6b8e2c9b125b717ec5"))
		})
	})
})
