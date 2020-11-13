package wer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codycollier/wer"
)

type werTest struct {
	werExpected  int
	waccExpected float64
	reference    string
	candidate    string
}

func runWerTests(t *testing.T, testdata []werTest) {
	for _, wt := range testdata {
		reference := strings.Split(wt.reference, " ")
		candidate := strings.Split(wt.candidate, " ")
		wer, _ := wer.WER(reference, candidate)
		msg := fmt.Sprintf("reference: %s, candidate: %s", reference, candidate)
		assert.Equal(t, wt.werExpected, wer, msg)
		// assert.Equal(t, wt.waccExpected, wacc)
	}
}

func TestExactMatch(t *testing.T) {
	testData := []werTest{
		{0, 1.0, "This is Green", "This is Green"},
		{0, 1.0, "This is Blue", "This is Blue"},
		{0, 1.0, "This is another Color", "This is another Color"},
		{0, 1.0, "colors are really great", "colors are really great"},
		{0, 1.0, "the quick brown fox jumped", "the quick brown fox jumped"},
	}
	runWerTests(t, testData)
}

func TestWERVariations(t *testing.T) {
	testData := []werTest{

		// missing one word
		{1, 0.0, "This is Green", "This is Blue"},
		{1, 0.0, "This is Green", "is Green"},
		{1, 0.0, "This is Green", "This Green"},
		{1, 0.0, "This is Green", "This is"},

		// mismatch case
		{1, 0.0, "Blue is a Color", "Blue is a color"},
		{2, -1.0, "Blue is a Color", "blue is a color"},
		{4, -3.0, "Blue is a Color", "BLUE IS A COLOR"},

		// inserts
		{0, 1.0, "token", "token"},
		{1, 0.0, "token", "token foo"},
		{1, 0.0, "token", "foo token"},
		{2, -1.0, "token", "foo token foo"},
		{3, -2.0, "token", "foo token foo foo"},
		{3, -2.0, "token", "foo foo token foo"},
		{4, -3.0, "token", "foo foo token foo foo"},

		// deletes
		{1, 0.0, "token foo", "foo"},
		{1, 0.0, "foo token", "foo"},
		{1, 0.0, "foo token foo", "foo foo"},
		{1, 0.0, "foo token foo foo", "foo foo foo"},
		{1, 0.0, "foo foo token foo", "foo foo foo"},
		{1, 0.0, "foo foo token foo foo", "foo foo foo foo"},

		// subs
		{1, 0.0, "one two three", "one two four"},
		{1, 0.0, "one two three", "one four three"},
		{1, 0.0, "one two three", "four two three"},
		{2, -1.0, "one two three", "four four three"},
		{2, -1.0, "one two three", "one four four"},
	}
	runWerTests(t, testData)
}
