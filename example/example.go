package main

import (
	"fmt"
	"strings"

	"github.com/codycollier/wer"
)

func main() {

	// strings to be evaluated
	knownGoodTranscript := "The quick brown fox jumps over the lazy dog"
	candidateTranscript := "the slow grey snail jumps over the lazy cat"

	// optionally lowercase / pre-process
	knownGoodTranscript = strings.ToLower(knownGoodTranscript)
	candidateTranscript = strings.ToLower(candidateTranscript)

	// convert to a list of words/tokens
	reference := strings.Split(knownGoodTranscript, " ")
	candidate := strings.Split(candidateTranscript, " ")

	// compare
	wer, wacc := wer.WER(reference, candidate)

	fmt.Printf("reference: %s\n", reference)
	fmt.Printf("candidate: %s\n", candidate)
	fmt.Printf("word error rate: %v\n", wer)
	fmt.Printf("word accuracy: %v\n", wacc)

}
