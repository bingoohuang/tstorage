package tstorage

type fakePartition struct {
	minT       int64
	maxT       int64
	numPoints  int
	IsReadOnly bool

	err error
}

func (f *fakePartition) insertRows(_ []Row) error {
	return f.err
}

func (f *fakePartition) selectRows(_ string, _ []Label, _, _ int64) dataPointList {
	return nil
}

func (f *fakePartition) selectAll() []Row {
	return nil
}

func (f *fakePartition) minTimestamp() int64 {
	return f.minT
}

func (f *fakePartition) maxTimestamp() int64 {
	return f.maxT
}

func (f *fakePartition) size() int {
	return f.numPoints
}

func (f *fakePartition) readOnly() bool {
	return f.IsReadOnly
}
