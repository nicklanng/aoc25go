package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_Horizontal(t *testing.T) {
	tests := []struct {
		name     string
		line     Line
		expected bool
	}{
		{
			name:     "horizontal line left to right",
			line:     Line{Start: Point{X: 0, Y: 1}, End: Point{X: 5, Y: 1}},
			expected: true,
		},
		{
			name:     "horizontal line right to left",
			line:     Line{Start: Point{X: 5, Y: 1}, End: Point{X: 0, Y: 1}},
			expected: true,
		},
		{
			name:     "vertical line",
			line:     Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 5}},
			expected: false,
		},
		{
			name:     "diagonal line",
			line:     Line{Start: Point{X: 0, Y: 0}, End: Point{X: 5, Y: 5}},
			expected: false,
		},
		{
			name:     "single point (same start and end)",
			line:     Line{Start: Point{X: 1, Y: 1}, End: Point{X: 1, Y: 1}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.line.Horizontal()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLine_Vertical(t *testing.T) {
	tests := []struct {
		name     string
		line     Line
		expected bool
	}{
		{
			name:     "vertical line bottom to top",
			line:     Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 5}},
			expected: true,
		},
		{
			name:     "vertical line top to bottom",
			line:     Line{Start: Point{X: 1, Y: 5}, End: Point{X: 1, Y: 0}},
			expected: true,
		},
		{
			name:     "horizontal line",
			line:     Line{Start: Point{X: 0, Y: 1}, End: Point{X: 5, Y: 1}},
			expected: false,
		},
		{
			name:     "diagonal line",
			line:     Line{Start: Point{X: 0, Y: 0}, End: Point{X: 5, Y: 5}},
			expected: false,
		},
		{
			name:     "single point (same start and end)",
			line:     Line{Start: Point{X: 1, Y: 1}, End: Point{X: 1, Y: 1}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.line.Vertical()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLine_Crosses(t *testing.T) {
	tests := []struct {
		name     string
		line1    Line
		line2    Line
		expected bool
	}{
		{
			name:     "horizontal crosses vertical at center",
			line1:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			line2:    Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}},
			expected: true,
		},
		{
			name:     "horizontal crosses vertical at edge",
			line1:    Line{Start: Point{X: 0, Y: 0}, End: Point{X: 2, Y: 0}},
			line2:    Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}},
			expected: true,
		},
		{
			name:     "horizontal crosses vertical at endpoint",
			line1:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			line2:    Line{Start: Point{X: 2, Y: 0}, End: Point{X: 2, Y: 2}},
			expected: true,
		},
		{
			name:     "horizontal does not cross vertical - X mismatch",
			line1:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			line2:    Line{Start: Point{X: 3, Y: 0}, End: Point{X: 3, Y: 2}},
			expected: false,
		},
		{
			name:     "horizontal does not cross vertical - Y mismatch",
			line1:    Line{Start: Point{X: 0, Y: 0}, End: Point{X: 2, Y: 0}},
			line2:    Line{Start: Point{X: 1, Y: 1}, End: Point{X: 1, Y: 2}},
			expected: false,
		},
		{
			name:     "vertical crosses horizontal at center",
			line1:    Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}},
			line2:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			expected: true,
		},
		{
			name:     "vertical crosses horizontal at edge",
			line1:    Line{Start: Point{X: 0, Y: 0}, End: Point{X: 0, Y: 2}},
			line2:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			expected: true,
		},
		{
			name:     "reversed horizontal crosses vertical",
			line1:    Line{Start: Point{X: 2, Y: 1}, End: Point{X: 0, Y: 1}},
			line2:    Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}},
			expected: true,
		},
		{
			name:     "reversed vertical crosses horizontal",
			line1:    Line{Start: Point{X: 1, Y: 2}, End: Point{X: 1, Y: 0}},
			line2:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			expected: true,
		},
		{
			name:     "both horizontal - no cross",
			line1:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			line2:    Line{Start: Point{X: 0, Y: 2}, End: Point{X: 2, Y: 2}},
			expected: false,
		},
		{
			name:     "both vertical - no cross",
			line1:    Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}},
			line2:    Line{Start: Point{X: 2, Y: 0}, End: Point{X: 2, Y: 2}},
			expected: false,
		},
		{
			name:     "diagonal lines - no cross",
			line1:    Line{Start: Point{X: 0, Y: 0}, End: Point{X: 2, Y: 2}},
			line2:    Line{Start: Point{X: 0, Y: 2}, End: Point{X: 2, Y: 0}},
			expected: false,
		},
		{
			name:     "horizontal line touches vertical at corner",
			line1:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}},
			line2:    Line{Start: Point{X: 2, Y: 1}, End: Point{X: 2, Y: 3}},
			expected: true,
		},
		{
			name:     "horizontal line just misses vertical",
			line1:    Line{Start: Point{X: 0, Y: 1}, End: Point{X: 1, Y: 1}},
			line2:    Line{Start: Point{X: 2, Y: 0}, End: Point{X: 2, Y: 2}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.line1.Crosses(tt.line2)
			assert.Equal(t, tt.expected, result, "line1.Crosses(line2)")
			// Crosses should be symmetric
			result2 := tt.line2.Crosses(tt.line1)
			assert.Equal(t, tt.expected, result2, "line2.Crosses(line1)")
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		expected float64
	}{
		{
			name:     "standard rectangle",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 5, MaxY: 3},
			expected: 24.0, // (5-0+1) * (3-0+1) = 6 * 4 = 24
		},
		{
			name:     "square",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 4, MaxY: 4},
			expected: 25.0, // (4-0+1) * (4-0+1) = 5 * 5 = 25
		},
		{
			name:     "rectangle with negative coordinates",
			rect:     Rectangle{MinX: -5, MinY: -3, MaxX: 0, MaxY: 0},
			expected: 24.0, // (0-(-5)+1) * (0-(-3)+1) = 6 * 4 = 24
		},
		{
			name:     "rectangle spanning negative and positive",
			rect:     Rectangle{MinX: -2, MinY: -1, MaxX: 3, MaxY: 4},
			expected: 36.0, // (3-(-2)+1) * (4-(-1)+1) = 6 * 6 = 36
		},
		{
			name:     "unit rectangle",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 1, MaxY: 1},
			expected: 4.0, // (1-0+1) * (1-0+1) = 2 * 2 = 4
		},
		{
			name:     "zero width rectangle",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 0, MaxY: 5},
			expected: 6.0, // (0-0+1) * (5-0+1) = 1 * 6 = 6
		},
		{
			name:     "zero height rectangle",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 5, MaxY: 0},
			expected: 6.0, // (5-0+1) * (0-0+1) = 6 * 1 = 6
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Area()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRectangle_Shrink(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		amount   float64
		expected Rectangle
	}{
		{
			name:     "shrink by 1",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 10, MaxY: 10},
			amount:   1.0,
			expected: Rectangle{MinX: 1, MinY: 1, MaxX: 9, MaxY: 9},
		},
		{
			name:     "shrink by 0.5",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 10, MaxY: 10},
			amount:   0.5,
			expected: Rectangle{MinX: 0.5, MinY: 0.5, MaxX: 9.5, MaxY: 9.5},
		},
		{
			name:     "shrink by 2",
			rect:     Rectangle{MinX: 5, MinY: 5, MaxX: 15, MaxY: 15},
			amount:   2.0,
			expected: Rectangle{MinX: 7, MinY: 7, MaxX: 13, MaxY: 13},
		},
		{
			name:     "shrink by 0 (no change)",
			rect:     Rectangle{MinX: 0, MinY: 0, MaxX: 10, MaxY: 10},
			amount:   0.0,
			expected: Rectangle{MinX: 0, MinY: 0, MaxX: 10, MaxY: 10},
		},
		{
			name:     "shrink rectangle with negative coordinates",
			rect:     Rectangle{MinX: -5, MinY: -5, MaxX: 5, MaxY: 5},
			amount:   1.0,
			expected: Rectangle{MinX: -4, MinY: -4, MaxX: 4, MaxY: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Shrink(tt.amount)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLine_Crosses(b *testing.B) {
	// Test case: horizontal line crossing vertical line
	horizontal := Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}}
	vertical := Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = horizontal.Crosses(vertical)
	}
}

func BenchmarkLine_Crosses_Reversed(b *testing.B) {
	// Test case: reversed horizontal line crossing vertical line
	horizontal := Line{Start: Point{X: 2, Y: 1}, End: Point{X: 0, Y: 1}}
	vertical := Line{Start: Point{X: 1, Y: 2}, End: Point{X: 1, Y: 0}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = horizontal.Crosses(vertical)
	}
}

func BenchmarkLine_Crosses_NoCross(b *testing.B) {
	// Test case: lines that don't cross
	horizontal := Line{Start: Point{X: 0, Y: 0}, End: Point{X: 2, Y: 0}}
	vertical := Line{Start: Point{X: 3, Y: 0}, End: Point{X: 3, Y: 2}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = horizontal.Crosses(vertical)
	}
}

func BenchmarkLine_Crosses_Mixed(b *testing.B) {
	// Test case: mix of crossing and non-crossing scenarios
	lines := []struct {
		l1 Line
		l2 Line
	}{
		{Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}}, Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}}},
		{Line{Start: Point{X: 0, Y: 0}, End: Point{X: 2, Y: 0}}, Line{Start: Point{X: 3, Y: 0}, End: Point{X: 3, Y: 2}}},
		{Line{Start: Point{X: 1, Y: 0}, End: Point{X: 1, Y: 2}}, Line{Start: Point{X: 0, Y: 1}, End: Point{X: 2, Y: 1}}},
		{Line{Start: Point{X: 0, Y: 0}, End: Point{X: 1, Y: 0}}, Line{Start: Point{X: 2, Y: 0}, End: Point{X: 2, Y: 2}}},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, pair := range lines {
			_ = pair.l1.Crosses(pair.l2)
		}
	}
}
