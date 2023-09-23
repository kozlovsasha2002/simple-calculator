package write_to_stdout

import (
	"fmt"
	"log"
	"simple-calculator/write_to_stdout/exp"
	"sync"
)

func CallCalculator2() {
	var numberOfExpressions int
	fmt.Print("Enter number of expressions: ")
	fmt.Scan(&numberOfExpressions)

	ch := make(chan exp.InputExp)
	results := make(chan OutputResp, numberOfExpressions)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < numberOfExpressions; i++ {
			var value1 int
			var value2 int
			var sign string
			_, err := fmt.Scan(&value1, &sign, &value2)
			if err != nil {
				log.Fatalf("invalid input data: %d %s %d", value1, sign, value2)
			}

			var data exp.InputExp
			data.SetValue1(value1)
			err = data.SetValue2(value2)
			if err != nil {
				log.Fatal(err)
			}
			err = data.SetSign(sign)
			if err != nil {
				log.Fatal(err)
			}

			ch <- data
		}
		close(ch)
	}()

	for i := range ch {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			calculate(results, i)
		}()
	}

	wg.Wait()
	close(results)

	for i := range results {
		fmt.Println(i.Exp, i.Result)
	}
}

type OutputResp struct {
	Exp    string
	Result int
}

func calculate(results chan OutputResp, input exp.InputExp) {
	currResult := 0
	switch input.Sign() {
	case "+":
		currResult = input.Value1() + input.Value2()
	case "-":
		currResult = input.Value1() - input.Value2()
	case "/":
		currResult = input.Value1() / input.Value2()
	case "*":
		currResult = input.Value1() * input.Value2()
	}
	results <- OutputResp{
		Exp:    fmt.Sprintf("%d %s %d =", input.Value1(), input.Sign(), input.Value2()),
		Result: currResult,
	}
}
