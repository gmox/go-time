package time_test

import (
	"github.com/gmox/go-time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	t "time"
)

var _ = Describe("Timestamp", func() {
	It("Constructs with date time", func () {
		rf := t.Now()

		tm := time.Timestamp{
			rf,
		}

		Expect(tm.Time).To(Equal(rf))
	})
})
