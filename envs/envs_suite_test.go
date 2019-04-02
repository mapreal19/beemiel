package envs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEnvs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Envs Suite")
}
