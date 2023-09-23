package exp

import "errors"

type InputExp struct {
	value1 int
	value2 int
	sign   string
}

func (i *InputExp) Value1() int {
	return i.value1
}

func (i *InputExp) Value2() int {
	return i.value2
}

func (i *InputExp) Sign() string {
	return i.sign
}

func (i *InputExp) SetValue1(value int) {
	i.value1 = value
}

func (i *InputExp) SetValue2(value int) error {
	if value != 0 && i.sign == "/" {
		return errors.New("you can't divide by zero")
	}
	i.value2 = value
	return nil
}

func (i *InputExp) SetSign(sign string) error {
	switch sign {
	case "+", "-", "/", "*":
		i.sign = sign
	default:
		return errors.New("incorrect operation")
	}
	return nil
}
