package wordcounter

import (
	"io"
	"testing"
)

func TestWordLineCounter_Lines(t *testing.T) {
	ts := []struct {
		text     string
		expected int
	}{
		{
			text:     "",
			expected: 0,
		},
		{
			text:     "hello",
			expected: 0,
		},
		{
			text:     "world\n",
			expected: 1,
		},
		{
			text:     "ほげほげ\nhogehoge\nhogehoge\n\n",
			expected: 4,
		},
	}

	for _, tc := range ts {
		wc := &WordLineCounter{}
		io.WriteString(wc, tc.text)

		if got := wc.Lines(); got != tc.expected {
			t.Errorf("unexpected lines. expected: %v, but got: %v", tc.expected, got)
		}
	}
}

func TestWordLineCounter_Words(t *testing.T) {
	ts := []struct {
		text     string
		expected int
	}{
		{
			text:     "",
			expected: 0,
		},
		{
			text:     "hello",
			expected: 1,
		},
		{
			text:     "world\n",
			expected: 1,
		},
		{
			text:     "ほげ ほげ\nhoge hoge\nhogehoge\nほげ\n",
			expected: 6,
		},
	}

	for _, tc := range ts {
		wc := &WordLineCounter{}
		io.WriteString(wc, tc.text)
		if got := wc.Words(); got != tc.expected {
			t.Errorf("unexpected words. expected: %v, but got: %v", tc.expected, got)
		}
	}
}
