package tracing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// +kubebuilder:scaffold:imports
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

func TestSuite(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t, "Tracing", []Reporter{printer.NewlineReporter{}})
}
