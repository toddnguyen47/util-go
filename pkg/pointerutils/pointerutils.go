package pointerutils

import "time"

func PtrString(input string) *string {
	return &input
}

func PtrInt(input int) *int {
	return &input
}

func PtrInt32(input int32) *int32 {
	return &input
}

func PtrInt64(input int64) *int64 {
	return &input
}

func PtrFloat32(input float32) *float32 {
	return &input
}

func PtrFloat64(input float64) *float64 {
	return &input
}

func PtrBool(input bool) *bool {
	return &input
}

func PtrDuration(duration time.Duration) *time.Duration {
	return &duration
}

func PtrUint(input uint) *uint {
	return &input
}

func PtrUint32(input uint32) *uint32 {
	return &input
}

func PtrUint64(input uint64) *uint64 {
	return &input
}
