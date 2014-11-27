package microincasso_test

import (
	"github.com/boxture/microincasso"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

type TestBackend struct {
	Response microincasso.Response
	Err      error
}

func (s *TestBackend) Call(method, path string, mireq *microincasso.Request) (*microincasso.Response, error) {
	return &s.Response, s.Err
}

func TestMicroincasso(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Microincasso Suite")
}
