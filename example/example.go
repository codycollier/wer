package main

import (
	"fmt"
	"strings"

	"github.com/codycollier/wer"
)

func main() {

	// strings to be evaluated
	knownGoodTranscript := "the quick brown fox jumps over the lazy dog"
	candidateTranscript := "the slow grey snail jumps over the lazy cat"

	// convert to a list of words/tokens
	reference := strings.Split(knownGoodTranscript, " ")
	candidate := strings.Split(candidateTranscript, " ")

	wer, wacc := wer.WER(reference, candidate)

	fmt.Printf("reference: %s\n", reference)
	fmt.Printf("candidate: %s\n", candidate)
	fmt.Printf("word error rate: %v\n", wer)
	fmt.Printf("word accuracy: %v\n", wacc)

}
