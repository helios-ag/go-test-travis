package go_test_travis

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestString(t *testing.T) {
	RegisterTestingT(t)

	t.Run("Should error", func(t *testing.T) {
		_, err := triggerError(true)
		Expect(err).To(HaveOccurred())
	})

	t.Run("Without error", func(t *testing.T) {
		res, err := triggerError(false)
		Expect(err).ToNot(HaveOccurred())
		Expect(res).To(ContainSubstring("All is OK"))
	})
}
