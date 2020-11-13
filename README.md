
## Word Error Rate (wer)

`wer` is a golang package which provides a function for calculating word error rate and word accuracy. It expects a pair of pre-tokenized and optionally pre-processed strings.


References:

* https://martin-thoma.com/word-error-rate-calculation/
* https://en.wikipedia.org/wiki/Word_error_rate


## Example Usage

```
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
```


```
[host]$ go run example/example.go
reference: [the quick brown fox jumps over the lazy dog]
candidate: [the slow grey snail jumps over the lazy cat]
word error rate: 4
word accuracy: -3
```
