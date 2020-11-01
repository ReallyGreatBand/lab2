package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type linkedListNode struct {
	value string
	next  *linkedListNode
	prev  *linkedListNode
}

type linkedList struct {
	head *linkedListNode
	tail *linkedListNode
}

func (list *linkedList) AddNode(value string) {
	newNode := &linkedListNode{value, nil, list.tail}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

// TODO: document this function.
// PrefixToPostfix converts
//"- * / 15 - 7 + 1 1 3 + 2 + 1 1"

func CalculatePrefix(input string) (string, error) {
	arr, err := checkInput(input)
	if err != nil {
		return "", err
	}
	list := createLinkedListFromArray(arr)
	curNode := list.head
	// check operators
	for true {
		if curNode == nil {
			return "", fmt.Errorf("Odd characters")
		}
		if curNode.next == nil {
			return "", fmt.Errorf("Odd characters")
		}
		if curNode.next.next == nil {
			return "", fmt.Errorf("Odd characters")
		}
		if isOperator(curNode.value) && isNumber(curNode.next.value) && isNumber(curNode.next.next.value) {
			res := calculate(curNode.value, curNode.next.value, curNode.next.next.value)
			curNode.value = res
			curNode.next = curNode.next.next.next
			if curNode.next != nil {
				curNode.next.prev = curNode
			}
			curNode = curNode.prev
			if curNode == nil {
				if list.head.next != nil {
					return "", fmt.Errorf("Odd characters")
				}
				resNumber, _ := strconv.ParseFloat(res, 64)
				res = strconv.FormatFloat(resNumber, 'f', -1, 64)
				return res, nil
			}
			if isNumber(curNode.value) {
				curNode = curNode.prev
			} else if !isOperator(curNode.value) {
				return "", fmt.Errorf("Odd characters")
			}
		} else {
			curNode = curNode.next
		}
	}
	return "", fmt.Errorf("Should never return")
}

func createLinkedListFromArray(arr []string) *linkedList {
	list := &linkedList{nil, nil}
	for i := 0; i < len(arr); i++ {
		list.AddNode(arr[i])
	}
	return list
}

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}

func isOperator(str string) bool {
	arr := []string{"+", "-", "*", "/", "^"}
	for _, operator := range arr {
		if operator == str {
			return true
		}
	}
	return false
}

func calculate(operator, num1, num2 string) string {
	a, _ := strconv.ParseFloat(num1, 64)
	b, _ := strconv.ParseFloat(num2, 64)
	switch operator {
	case "+":
		return fmt.Sprintf("%f", a+b)
	case "-":
		return fmt.Sprintf("%f", a-b)
	case "*":
		return fmt.Sprintf("%f", a*b)
	case "/":
		return fmt.Sprintf("%f", a/b)
	default:
		return fmt.Sprintf("%f", math.Pow(a, b))
	}
}

func checkInput(input string) ([]string, error) {
	if input == "" {
		return make([]string, 0), fmt.Errorf("Empty string")
	}
	arr := strings.Split(input, " ")
	for _, el := range arr {
		if !isOperator(el) && !isNumber(el) {
			return make([]string, 0), fmt.Errorf("Incorrect characters")
		}
	}
	return arr, nil
}
