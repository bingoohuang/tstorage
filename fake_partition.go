package tstorage

type fakePartition struct {
	err       error
	minT      int64
	maxT      int64
	numPoints int
	IsActive  bool
}

func (f *fakePartition) insertRows(_ []Row) ([]Row, error) { return nil, f.err }

func (f *fakePartition) selectDataPoints(_ string, _ []Label, _, _ int64) ([]*DataPoint, error) {
	return nil, f.err
}

func (f *fakePartition) minTimestamp() int64   { return f.minT }
func (f *fakePartition) maxTimestamp() int64   { return f.maxT }
func (f *fakePartition) size() int             { return f.numPoints }
func (f *fakePartition) active() bool          { return f.IsActive }
func (f *fakePartition) clean() error          { return nil }
func (f *fakePartition) needClean(*int64) bool { return false }
