package calculation

import (
	"fmt"
)

func PassSpace(exp string) string {
	a := ""
	for _, i := range exp {
		if string(i) != " " {
			a += string(i)
		}
	}
	return a
}

func ElOper(stack *[]float64, op uint8) error {
	if len(*stack) < 2 {
		return fmt.Errorf("incorrect expression")
	}
	b := (*stack)[len(*stack)-1]
	a := (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]
	switch op {
	case '-':
		(*stack) = append(*stack, a-b)
		return nil
	case '+':
		(*stack) = append(*stack, a+b)
		return nil
	case '*':
		(*stack) = append(*stack, a*b)
		return nil
	case '/':
		if b == 0 {
			return fmt.Errorf("incorrect expression")
		}
		(*stack) = append(*stack, a/b)
		return nil
	case '(':
		return fmt.Errorf("error: '(' without ')'")

	}
	return fmt.Errorf("incorrect")
}

func IsDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}
func precedence(op uint8) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func isOperator(c uint8) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func Calc(expression string) (float64, error) {
	expression = PassSpace(expression)
	var stack []float64

	var stackForOperators []uint8
	var beforeV uint8 = ' '
	if expression == "" {
		return 0, fmt.Errorf("expression is empty")
	}
	for i := 0; i < len(expression); i++ {

		num := 0.0
		v := expression[i]

		if !IsDigit(v) && !isOperator(v) && v != '(' && v != ')' {
			return 0, fmt.Errorf("incorrect symbols")
		}

		if IsDigit(expression[i]) {
			for i < len(expression) && IsDigit(expression[i]) {
				num *= 10
				num += float64(expression[i] - '0')

				i++
			}
			i--
			stack = append(stack, num)
		} else if isOperator(v) {
			for j := len(stackForOperators) - 1; j >= 0; j-- {
				if precedence(v) > precedence(stackForOperators[j]) {
					break
				}

				_ = ElOper(&stack, stackForOperators[j])
				stackForOperators = stackForOperators[:len(stackForOperators)-1]
			}
			stackForOperators = append(stackForOperators, v)
		}
		if v == '(' {
			if IsDigit(beforeV) || isOperator(expression[i+1]) {
				return 0, fmt.Errorf("incorrect expression")
			}
			stackForOperators = append(stackForOperators, v)

		}
		if v == ')' {
			if (i != len(expression)-1) && (IsDigit(expression[i+1]) || isOperator(beforeV)) {
				return 0, fmt.Errorf("incorrect expression")
			}
			for j := len(stackForOperators) - 1; j >= -1; j-- {
				if j == -1 {
					return 0, fmt.Errorf(" error: ')' without '('")
				}
				if stackForOperators[j] == '(' {
					stackForOperators = stackForOperators[:len(stackForOperators)-1]
					break
				}
				er := ElOper(&stack, stackForOperators[j])
				if er != nil {
					return 0, er
				}
				stackForOperators = stackForOperators[:len(stackForOperators)-1]

			}

		}
		beforeV = v

	}
	if isOperator(beforeV) {
		return 0, fmt.Errorf("incorrect expression")
	}
	for j := len(stackForOperators) - 1; j >= 0; j-- {

		err := ElOper(&stack, stackForOperators[j])
		if err != nil {
			return 0, err
		}
		stackForOperators = stackForOperators[:len(stackForOperators)-1]
	}
	return stack[len(stack)-1], nil

}
