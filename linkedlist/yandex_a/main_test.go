package main

import (
	"bytes"
	"testing"
)

func TestSolve(t *testing.T) {
	testInput(t, `8
1 0 5
1 1 3
1 2 8
1 3 13
2 1
2 2
2 3
2 4
`, `5
3
8
13
`)

	testInput(t, `8
1 0 5
2 1
3 1
1 0 6
1 0 7
2 2
1 1 5
2 2
`, `5
6
5
`)
}

func testInput(t *testing.T, input string, expectedOutput string) {
	var output bytes.Buffer
	solve(bytes.NewBufferString(input), &output)
	if output.String() != expectedOutput {
		t.Errorf("expected:\n%s\ngot:\n%s", expectedOutput, output.String())
	}
}
