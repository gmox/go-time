package time_test

import (
	"encoding/json"
	"github.com/gmox/go-time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	t "time"
)

type testStruct struct {
	Timestamp *time.Timestamp `json:"date"`
}

var _ = Describe("Timestamp", func() {
	It("Constructs with date time", func () {
		rf := t.Now()

		tm := time.Timestamp{
			rf,
		}

		Expect(tm.Time).To(Equal(rf))
	})

	It("Decodes with date string", func () {
		data := []byte(`{"date":"2018-09-07T16:39:21Z"}`)

		expected := &time.Timestamp{
			Time: t.Date(2018, 9, 7, 16, 39, 21, 0, t.UTC),
		}

		var tm testStruct

		if err := json.Unmarshal(data, &tm); err != nil {
			panic(err)
		}

		Expect(tm.Timestamp).To(Not(BeNil()))
		Expect(tm.Timestamp).To(Equal(expected))
	})

	It("Encodes timestamp", func () {
		data := testStruct {
			Timestamp: &time.Timestamp{
				Time: t.Date(2018, 9, 7, 16, 39, 21, 0, t.UTC),
			},
		}

		expected := []byte(`{"date":"2018-09-07T16:39:21Z"}`)

		encoded, err := json.Marshal(data)

		if err != nil {
			panic(err)
		}

		Expect(encoded).To(Equal(expected))
	})
})
