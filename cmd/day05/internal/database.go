package internal

import (
	"slices"
	"strconv"
	"strings"

	"github.com/nicklanng/aoc25go/internal/data"
)

type Database struct {
	Ranges []data.Range
	Ids    []int
}

func ParseDatabase(database string) (Database, error) {
	parsingRanges := true
	var ranges []data.Range
	var ids []int

	for line := range strings.SplitSeq(database, "\n") {
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			r, err := data.ParseRange(line)
			if err != nil {
				return Database{}, err
			}
			ranges = append(ranges, r)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return Database{}, err
			}
			ids = append(ids, id)
		}
	}

	slices.SortFunc(ranges, func(a, b data.Range) int {
		return a.Min - b.Min
	})
	slices.Sort(ids)

	ranges = compactRanges(ranges)

	return Database{Ranges: ranges, Ids: ids}, nil
}

func compactRanges(ranges []data.Range) []data.Range {
	var compactedRanges = []data.Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		r := compactedRanges[len(compactedRanges)-1]
		if r.Overlaps(ranges[i]) {
			r.Max = max(r.Max, ranges[i].Max)
		} else {
			compactedRanges = append(compactedRanges, ranges[i])
		}
	}

	return compactedRanges
}

func (d Database) FindId(id int) bool {
	for _, r := range d.Ranges {
		if r.Min <= id && r.Max >= id {
			return true
		}
	}
	return false
}

func (d Database) CountAvailableIds() int {
	var availableIds int

	for _, r := range d.Ranges {
		availableIds += r.Length()
	}

	return availableIds
}
