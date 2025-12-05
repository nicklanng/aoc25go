package data

import (
	"strconv"
	"strings"
)

type Range struct {
	Min int
	Max int
}

func ParseRange(s string) (Range, error) {
	parts := strings.Split(s, "-")

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return Range{}, err
	}

	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return Range{}, err
	}

	return Range{Min: min, Max: max}, nil
}

func (r Range) Overlaps(other Range) bool {
	return r.Min <= other.Max && r.Max >= other.Min
}

func (r Range) Length() int {
	return r.Max - r.Min + 1
}
