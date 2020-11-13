package wer_test

import (
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
		wer, wacc := wer.WER(wt.reference, wt.candidate)
		assert.Equal(t, wer, wt.werExpected)
		assert.Equal(t, wacc, wt.waccExpected)
	}
}
