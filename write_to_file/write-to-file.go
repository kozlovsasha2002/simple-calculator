package write_to_file

import (
	"fmt"
	"os"
	"runtime"
)

func CallCalculator() {
	f, err := os.Create("./write_to_file/result.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ch := make(chan *Exp)
	ch2 := make(chan *Resp)
	numberOfExpressions := 0
	fmt.Print("number of expressions: ")
	fmt.Scan(&numberOfExpressions)

	go func() {
		for i := 0; i < numberOfExpressions; i++ {
			value1 := 0
			value2 := 0
			sign := ""
			fmt.Scan(&value1, &sign, &value2)

			exp := &Exp{
				value1: value1,
				value2: value2,
				sign:   sign,
			}
			ch <- exp
		}
		close(ch)
	}()

	go func() {
		for i := range ch {
			calculateExpression(ch2, *i)
		}
		close(ch2)
	}()
	fmt.Println("number of goroutines =", runtime.NumGoroutine())

	for i := range ch2 {
		_, err = f.WriteString(fmt.Sprintf("%s = %d\n", i.exp, i.value))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("number of goroutines =", runtime.NumGoroutine())
	fmt.Println("Finish...")
}

func calculateExpression(ch chan *Resp, exp Exp) {
	result := 0
	switch exp.sign {
	case "+":
		result = exp.value1 + exp.value2
	case "-":
		result = exp.value1 - exp.value2
	case "/":
		result = exp.value1 / exp.value2
	case "*":
		result = exp.value1 * exp.value2
	}
	ch <- &Resp{
		exp:   fmt.Sprintf("%d %s %d", exp.value1, exp.sign, exp.value2),
		value: result,
	}
}

type Exp struct {
	value1 int
	value2 int
	sign   string
}

type Resp struct {
	exp   string
	value int
}
