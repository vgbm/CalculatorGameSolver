package main

import (
	"strconv"
	"strings"
	"math"
	"fmt"
)

type OpGenerator func(string) (Op, error)

type Op func(int) int

type PuzzleSetup struct {
	Moves int `json:"moves"`
	Start int `json:"start"`
	Goal int `json:"goal"`
	OpStrings []string `json:"Ops"`
	Ops []Op
	prev *PuzzleSetup
}

var OpsParseFunctions = map[string]OpGenerator{
	"+" : parseAdd,
	"-" : parseSub,
	"x" : parseMul,
	"<<" : parseShift,
	"^" : parseExp,
}

func (i *PuzzleSetup) ParseOps() error {
	for _, opString := range i.OpStrings {
		for opToken, opGenerator := range OpsParseFunctions {
			if strings.HasPrefix(opString, opToken) {
				f, err := opGenerator(opString)

				if err != nil {
					return err
				}

				i.Ops = append(i.Ops, f)
			}
		}
	}
	return nil
}

func parseAdd(s string) (Op, error) {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, err
	}

	return func(x int) int {
		return x + n
	}, nil
}

func parseSub(s string) (Op, error) {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, err
	}

	return func(x int) int {
		return x - n
	}, nil
}

func parseMul(s string) (Op, error) {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, err
	}

	return func(x int) int {
		return x * n
	}, nil
}

func parseShift(s string) (Op, error) {
	n, err := strconv.Atoi(s[2:])
	if err != nil {
		return nil, err
	}

	return func(x int) int {
		temp := x * 10
		if x > 0 {
			return temp + n
		}
		return temp - n
	}, nil
}

func parseExp(s string) (Op, error) {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return nil, err
	}

	return func(x int) int {
		return int(math.Pow(float64(x), float64(n)))
	}, nil
}

func (i *PuzzleSetup) Solve() {
	is := &PuzzleStep{
	i.Start,
	i.Moves,
	"Start",
	nil,
	nil,
	}

	if solutionRoot := is.Solve(*i); solutionRoot != nil {
		solutionRoot.PrintSolution()
	} else {
		fmt.Println("No solution found.")
	}
}

