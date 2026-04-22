package main

import (
	"errors"
	"fmt"
)

// Type Assertion
func doubler(v interface{}) (string, error) {

	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}

	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("Unsupported type passed")

}

func typeAssertion() {
	res, _ := doubler(5)
	fmt.Println("5 :", res)

	res, _ = doubler(" Hello World")
	fmt.Println("Hello World: ", res)

	_, err := doubler(true)
	fmt.Println("true: ", err)

}

// Type Switch

func anotherDoubler(v interface{}) (string, error) {
	switch t := v.(type) {

	case string:
		return t + t, nil
	case bool:
		if t {
			return "true true", nil
		}
		return "false false", nil
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}

		return fmt.Sprint(t.(float32) * 2), nil
	default:
		return "", errors.New("Unsupported type passed")
	}

}
func typeSwitch() {

	res, _ := anotherDoubler(-5)
	fmt.Println("-5 :", res)
	res, _ = anotherDoubler(5)
	fmt.Println("5 :", res)
	res, _ = anotherDoubler("yum")
	fmt.Println("yum :", res)
	res, _ = anotherDoubler(true)
	fmt.Println("true:", res)
	res, _ = anotherDoubler(float32(3.14))
	fmt.Println("3.14:", res)

}

// type checker
func getData() []interface{} {
	return []interface{}{
		12,
		12.334,
		false,
		"Hello",
		struct{}{},
	}
}

func getTypeName(v interface{}) string {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return "int"

	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "boolean"
	case struct{}:
		return "struct"
	default:
		return "unknown"
	}
}

func main() {

	typeAssertion()
	typeSwitch()

	data := getData()

	for i := 0; i < len(data); i++ {
		fmt.Printf("%v is %v\n", data[i], getTypeName(data[i]))
	}
}
