package e2e

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("e2e testing", func() {
	ginkgo.Context("Forever true", func() {
		ginkgo.It("Test forever true", func() {
			gomega.Eventually(func(g gomega.Gomega) (bool, error) {
				return true, nil
			}, 60, 10).Should(gomega.Equal(true))
		})
	})
})
