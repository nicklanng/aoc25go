package internal

import (
	"testing"

	"github.com/nicklanng/aoc25go/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestParseDatabase(t *testing.T) {
	input := `3-5
10-14
11-12
16-20
16-21
12-18

1
5
11
8
17
32`

	database, err := ParseDatabase(input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	assert.Equal(t, []data.Range{
		{Min: 3, Max: 5},
		{Min: 10, Max: 20},
	}, database.Ranges)
	assert.Equal(t, []int{1, 5, 8, 11, 17, 32}, database.Ids)
}

func TestFindId(t *testing.T) {
	database := Database{
		Ranges: []data.Range{
			{Min: 3, Max: 5},
			{Min: 10, Max: 14},
			{Min: 12, Max: 18},
			{Min: 16, Max: 20},
		},
		Ids: []int{1, 5, 8, 11, 17, 32},
	}

	assert.False(t, database.FindId(1))
	assert.True(t, database.FindId(5))
	assert.False(t, database.FindId(8))
	assert.True(t, database.FindId(11))
	assert.True(t, database.FindId(17))
	assert.False(t, database.FindId(32))
}

func TestCountAvailableIds(t *testing.T) {
	database := Database{
		Ranges: []data.Range{
			{Min: 3, Max: 5},
			{Min: 10, Max: 20},
		},
	}

	assert.Equal(t, 18, database.CountAvailableIds())
}
