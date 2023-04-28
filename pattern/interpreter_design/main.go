package main

import (
	"fmt"
	"github.com/championlong/design-patterns-go/pattern/interpreter_design/interpreter"
)

/*
解释器模式：
解释器模式为某个语⾔定义它的语法（或者叫⽂法）表⽰，并定义⼀个解释器⽤来处理这个语法
⼀般的做法是，将语法规则拆分成⼀些⼩的独⽴的单元，然后对每个单元进⾏ 解析，最终合并为对整个语法规则的解析。
*/
func main() {
	// 实现一套告警规则只包含“||、&&、>、<、==”这五个运算符，其中，“>、<、==”运算符的优先级高于“||、&&”运算符，“&&”运算符优先级高于“||”
	// key1 > 100 && key2 < 1000 || key3 == 200
	alertRule := interpreter.NewAlertRuleInterpreter("key1 > 100")
	result := alertRule.Interpret(map[string]int{
		"key1": 10,
	})
	fmt.Println(result)
}
