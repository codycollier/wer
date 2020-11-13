package wer_test

import (
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

func TestWerWaccExactMatchWithCase(t *testing.T) {
	testData := []werTest{
		{0, 1.0, "This is Green", "This is Green"},
		{0, 1.0, "This is Blue", "This is Green"},
		{0, 1.0, "This is another Color", "This is another Color"},
	}

	for _, wt := range testData {
		reference := strings.Split(wt.reference, "")
		candidate := strings.Split(wt.candidate, "")
		wer, wacc := wer.WER(reference, candidate)
		assert.Equal(t, wer, wt.werExpected)
		assert.Equal(t, wacc, wt.waccExpected)
	}
}
