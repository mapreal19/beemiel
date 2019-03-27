package passwords_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPasswords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Passwords Suite")
}
