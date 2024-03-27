package main

import (
	"flag"
	"log"
)

type stack struct {
	elements          []rune
	supportedBrackets map[rune]rune
}

func (stack *stack) isOpenBracket(element rune) bool {
	for k, _ := range stack.supportedBrackets {
		if k == element {
			return true
		}
	}
	return false
}

func (stack *stack) isCloseBracket(element rune) bool {
	for _, v := range stack.supportedBrackets {
		if v == element {
			return true
		}
	}
	return false
}

func (stack *stack) push(element rune) {
	stack.elements = append(stack.elements, element)
}

func (stack *stack) pop() *rune {
	index := len(stack.elements) - 1
	if index == -1 {
		return nil
	}

	lastElement := stack.elements[index]
	stack.elements = stack.elements[:index]
	return &lastElement
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	stack := stack{
		supportedBrackets: map[rune]rune{
			'(': ')',
			'[': ']',
			'{': '}',
		},
	}

	for _, e := range expr {
		if stack.isOpenBracket(e) {
			stack.push(e)
		} else if stack.isCloseBracket(e) {
			lastBracket := stack.pop()
			if lastBracket == nil || stack.supportedBrackets[*lastBracket] != e {
				return false
			}
		}
	}

	return len(stack.elements) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
