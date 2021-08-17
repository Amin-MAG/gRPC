package handler

import (
	"bale_interview/apis"
	"context"
)

// CalculateListMajority to find most repeated number
// that repeated more than size / 2 times
func (s *TestServer) CalculateListMajority(ctx context.Context, in *apis.ListMajorityRequest) (*apis.ListMajorityReply, error) {
	integerCounter := make(map[int32]int32)
	listSize := int32(len(in.Numbers))

	// fill the integer counter
	fillIntegerCounter(in, integerCounter)

	// find the most repeated item
	exists, maxim := findMostRepeatedItemGreaterHalf(integerCounter, listSize)

	if exists == false {
		return &apis.ListMajorityReply{
			Message:        "Not found any majority",
			MajorityNumber: 0,
		}, nil
	} else {
		return &apis.ListMajorityReply{
			Message:        "majority is Founded",
			MajorityNumber: maxim,
		}, nil
	}
}

// findMostRepeatedItemGreaterHalf
// It takes the integerCounter and find the most repeated item that
// is greater than half size of list.
func findMostRepeatedItemGreaterHalf(integerCounter map[int32]int32, listSize int32) (bool, int32) {
	var exists = false
	var maxim int32
	for key, value := range integerCounter {
		if value > listSize/2 {
			exists = true
			maxim = key
		}
	}
	return exists, maxim
}

// fillIntegerCounter
// Count the repeated number and put them in integerCounter
// integerCounter is int32 map to int32
func fillIntegerCounter(in *apis.ListMajorityRequest, integerCounter map[int32]int32) {
	for _, number := range in.Numbers {
		if _, ok := integerCounter[number]; ok {
			integerCounter[number]++
		} else {
			integerCounter[number] = 1
		}
	}
}
