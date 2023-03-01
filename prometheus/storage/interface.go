package storage

import (
	"github.com/EdgeCloudX/promthheus_client/prometheus/model/labels"
	"github.com/EdgeCloudX/promthheus_client/prometheus/tsdb/chunks"
)

// SeriesSet contains a set of series.
type SeriesSet interface {
	Next() bool
	// At returns full series. Returned series should be iterable even after Next is called.
	At() Series
	// The error that iteration as failed with.
	// When an error occurs, set cannot continue to iterate.
	Err() error
	// A collection of warnings for the whole set.
	// Warnings could be return even iteration has not failed with error.
	Warnings() Warnings
}

// Series exposes a single time series and allows iterating over samples.
type Series interface {
	LabelsInterface
	SampleIterable
}

// Labels represents an item that has labels e.g. time series.
type LabelsInterface interface {
	// Labels returns the complete set of labels. For series it means all labels identifying the series.
	Labels() labels.Labels
}

type SampleIterable interface {
	// Iterator returns a new, independent iterator of the data of the series.
	Iterator() chunks.Iterator
}
type Warnings []error

type testSeriesSet struct {
	series Series
}

func (s testSeriesSet) Next() bool         { return true }
func (s testSeriesSet) At() Series         { return s.series }
func (s testSeriesSet) Err() error         { return nil }
func (s testSeriesSet) Warnings() Warnings { return nil }

// TestSeriesSet returns a mock series set
func TestSeriesSet(series Series) SeriesSet {
	return testSeriesSet{series: series}
}

type errSeriesSet struct {
	err error
}

func (s errSeriesSet) Next() bool         { return false }
func (s errSeriesSet) At() Series         { return nil }
func (s errSeriesSet) Err() error         { return s.err }
func (s errSeriesSet) Warnings() Warnings { return nil }

// ErrSeriesSet returns a series set that wraps an error.
func ErrSeriesSet(err error) SeriesSet {
	return errSeriesSet{err: err}
}
