package log_stream_plugin_test

import (
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"github.com/cloudfoundry/log-stream-cli/internal/log_stream_plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RLPRequestFactory", func() {
	It("makes a valid selector for all source ids and metric types when all metric types are valid", func() {
		expected := &loggregator_v2.EgressBatchRequest{
			Selectors: []*loggregator_v2.Selector{
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Gauge{
						Gauge: &loggregator_v2.GaugeSelector{},
					},
				},
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Counter{
						Counter: &loggregator_v2.CounterSelector{},
					},
				},
			},
		}
		actual, err := log_stream_plugin.MakeRequest([]string{"foo"}, []string{"gauge", "counter"})

		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	It("makes one selector with no source id when not given a source id filter", func() {
		expected := &loggregator_v2.EgressBatchRequest{
			Selectors: []*loggregator_v2.Selector{
				{
					Message: &loggregator_v2.Selector_Event{
						Event: &loggregator_v2.EventSelector{},
					},
				},
				{
					Message: &loggregator_v2.Selector_Log{
						Log: &loggregator_v2.LogSelector{},
					},
				},
			},
		}
		actual, err := log_stream_plugin.MakeRequest([]string{}, []string{"event", "log"})

		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	It("makes a selector for all metric types when given no metric type filter", func() {
		expected := &loggregator_v2.EgressBatchRequest{
			Selectors: []*loggregator_v2.Selector{
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Log{
						Log: &loggregator_v2.LogSelector{},
					},
				},
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Counter{
						Counter: &loggregator_v2.CounterSelector{},
					},
				},
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Event{
						Event: &loggregator_v2.EventSelector{},
					},
				},
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Gauge{
						Gauge: &loggregator_v2.GaugeSelector{},
					},
				},
				{
					SourceId: "foo",
					Message: &loggregator_v2.Selector_Timer{
						Timer: &loggregator_v2.TimerSelector{},
					},
				},
			},
		}
		actual, err := log_stream_plugin.MakeRequest([]string{"foo"}, []string{})

		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	It("returns an error when given invalid metric types", func() {
		_, err := log_stream_plugin.MakeRequest([]string{"source-one", "source-two"}, []string{"gauge", "foo", "bar"})

		Expect(err.Error()).To(Equal("invalid metric type(s): foo, bar"))
	})
})
