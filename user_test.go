package microincasso_test

import (
	. "github.com/boxture/microincasso"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	Describe("Status", func() {
		It("should be bla", func() {
			SetDebug(true)
			Secret = "3F29A724-00E9-410D-90A5-7B4E391DE266"
			Username = "PostNLPakketten02"
			Password = "QhjlqaajKAf5mtVh"
			// t := &TestBackend{Response: Response{Status: 0}, Err: nil}
			// SetBackend(t)
			u := User{Reference: "31641085630"}
			_, err := u.Status()
			Expect(err).NotTo(HaveOccurred())
			// Expect(status).To(Equal(0))

		})
	})
})
