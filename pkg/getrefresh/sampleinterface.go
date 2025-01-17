package getrefresh

type SampleInterface interface {
	SampleFunc() int
}

type sampleStruct struct{}

func newSampleStruct() SampleInterface {
	return &sampleStruct{}
}

func (i1 *sampleStruct) SampleFunc() int {
	return 42
}
