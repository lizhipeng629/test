package test_study_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestStudy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestStudy Suite")
}
