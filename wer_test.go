package wer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codycollier/wer"
)

type werTest struct {
	werExpected float64
	reference   string
	candidate   string
}

func runWerTests(t *testing.T, testdata []werTest) {
	for _, wt := range testdata {
		reference := strings.Split(wt.reference, " ")
		candidate := strings.Split(wt.candidate, " ")
		wer, wacc := wer.WER(reference, candidate)

		msg := fmt.Sprintf("wer mismatch. reference: %s, candidate: %s", reference, candidate)
		assert.Equal(t, wt.werExpected, wer, msg)

		waccExpected := float64(1.0) - wer
		msg = fmt.Sprintf("wacc mismatch. reference: %s, candidate: %s", reference, candidate)
		assert.Equal(t, waccExpected, wacc, msg)
	}
}

func TestExactMatch(t *testing.T) {
	testData := []werTest{
		{0.0, "This is Green", "This is Green"},
		{0.0, "This is the color Blue", "This is the color Blue"},
		{0.0, "This is another Color", "This is another Color"},
		{0.0, "colors are really great", "colors are really great"},
		{0.0, "the quick brown fox jumps", "the quick brown fox jumps"},
	}
	runWerTests(t, testData)
}

func TestWERVariations(t *testing.T) {
	testData := []werTest{

		// missing one word
		{0.3333333333333333, "This is Green", "This is Blue"},
		{0.2, "This is the color Green", "This is the color Orange"},
		{0.2, "This is the color Green", "is the color Green"},
		{0.2, "This is the color Green", "This is the color"},
		{0.2, "This is the color Green", "This is the Green"},

		// mismatch case
		{0.25, "Blue is a Color", "Blue is a color"},
		{0.5, "Blue is a Color", "blue is a color"},
		{1.0, "Blue is a Color", "BLUE IS A COLOR"},

		// punctuation
		{0.5, "What is the time?", "What's the time?"},
		{0.5, "Oh no!", "no!"},

		// inserts
		{0.0, "token", "token"},
		{1.0, "token", "token foo"},
		{1.0, "token", "foo token"},
		{2.0, "token", "foo token foo"},
		{3.0, "token", "foo token foo foo"},
		{3.0, "token", "foo foo token foo"},
		{4.0, "token", "foo foo token foo foo"},

		// deletes
		{0.5, "token foo", "foo"},
		{0.5, "foo token", "foo"},
		{0.5, "foo token foo foo", "foo foo"},
		{0.25, "foo token foo foo", "foo foo foo"},
		{0.25, "foo foo token foo", "foo foo foo"},
		{0.2, "foo foo token foo foo", "foo foo foo foo"},

		// subs
		{0.25, "one two three four", "one two three five"},
		{0.25, "one two three four", "one five three four"},
		{0.25, "one two three four", "five two three four"},
		{0.5, "one two three four", "five five three four"},
		{0.5, "one two three four", "one five five four"},
	}
	runWerTests(t, testData)
}
