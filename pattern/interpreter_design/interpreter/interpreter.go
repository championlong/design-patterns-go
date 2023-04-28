package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type Expression interface {
	Interpret(map[string]int) bool
}

type GreaterExpression struct {
	Key   string
	Value int
}

func NewGreaterExpression(strExpression string) GreaterExpression {
	strArray := strings.Split(strExpression, " ")
	if len(strArray) != 3 || strArray[1] != ">" {
		fmt.Println("[greater] Expression is invalid: " + strExpression)
		return GreaterExpression{}
	}
	value, err := strconv.Atoi(strArray[2])
	if err != nil {
		fmt.Println("[greater]", err)
		return GreaterExpression{}
	}
	return GreaterExpression{
		Key:   strArray[0],
		Value: value,
	}
}

func (greater GreaterExpression) Interpret(stats map[string]int) bool {
	if value, ok := stats[greater.Key]; ok {
		return value > greater.Value
	}
	return false
}

type AndExpression struct {
	Expressions []Expression
}

func NewAndExpression(strExpression string) AndExpression {
	andExpression := AndExpression{}
	strExpressions := strings.Split(strExpression, "&&")
	for _, strExpr := range strExpressions {
		if strings.Contains(strExpr, ">") {
			andExpression.Expressions = append(andExpression.Expressions, NewGreaterExpression(strExpr))
		} else if strings.Contains(strExpr, "<") {

		} else if strings.Contains(strExpr, "==") {

		} else {

		}
	}
	return andExpression
}

func (and AndExpression) Interpret(stats map[string]int) bool {
	for _, expr := range and.Expressions {
		if !expr.Interpret(stats) {
			return false
		}
	}
	return true
}

type OrExpression struct {
	Expressions []Expression
}

func NewOrExpression(strExpression string) OrExpression {
	orExpression := OrExpression{}
	strExpressions := strings.Split(strExpression, "||")
	for _, strExpr := range strExpressions {
		orExpression.Expressions = append(orExpression.Expressions, NewAndExpression(strExpr))
	}
	return orExpression
}

func (and OrExpression) Interpret(stats map[string]int) bool {
	for _, expr := range and.Expressions {
		if !expr.Interpret(stats) {
			return false
		}
	}
	return true
}

type AlertRuleInterpreter struct {
	expression Expression
}

func NewAlertRuleInterpreter(ruleExpression string) Expression {
	return NewOrExpression(ruleExpression)
}

func (alert AlertRuleInterpreter) Interpret(stats map[string]int) bool {
	return alert.Interpret(stats)
}
