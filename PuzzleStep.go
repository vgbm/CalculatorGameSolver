package main

import "fmt"

type PuzzleStep struct {
	Current   int
	MovesLeft int
	OpString  string
	Prev      *PuzzleStep
	Next      *PuzzleStep
}

func (is *PuzzleStep) Solve(setup PuzzleSetup) *PuzzleStep {
	if is.MovesLeft <= 0 {
		if setup.Goal == is.Current {
			return is.Reverse()
		} else {
			return nil
		}
	}

	for _, op := range setup.Ops {
		newInput := &PuzzleStep{
			op(is.Current),
			is.MovesLeft - 1,
			"",
			is,
			nil,
		}

		if solution := newInput.Solve(setup); solution != nil {
			return solution
		}

	}
	return nil
}

func (is *PuzzleStep) Reverse() *PuzzleStep {
	curr := is
	temp := curr

	for curr.Prev != nil {
		curr = curr.Prev
		curr.Next = temp
		temp = curr
	}

	return curr
}

func (is *PuzzleStep) PrintSolution() {
	fmt.Printf(" -> %d", is.Current)

	if is.Next != nil {
		is.Next.PrintSolution()
	}
}