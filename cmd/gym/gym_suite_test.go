package gym_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGym(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gym Suite")
}
