package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (*float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse:%v", err)
	}

	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (*float64, error) {
	operands := strings.Fields(expr)

	operandsCount := len(operands)
	if operandsCount != 3 {
		return nil, fmt.Errorf("cannot calculate: need 3 operands, got %d", operandsCount)
	}

	left, err := parseOperand(operands[0])
	if err != nil {
		return nil, err
	}

	right, err := parseOperand(operands[2])
	if err != nil {
		return nil, err
	}

	operator, ok := operators[operands[1]]
	if !ok {
		return nil, fmt.Errorf("cannot calculate: %s is unknown", operands[1])
	}

	result := operator(*left, *right)
	return &result, nil
}

func main() {
	expr := flag.String("expr", "", "The expression to calculate on, separated by spaces.")
	flag.Parse()
	result, err := calculate(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s = %.2f\n", *expr, *result)
}
