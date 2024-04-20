package main

import "fmt"

func main() {
	fmt.Println("hello")
	var sums []int = []int{100, 200, 400, 300}
	fmt.Println(sum(sums...))
}

func multiplereturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

func multiplenamedreturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("some error happend")
		return 0, err
	}
	return
}

func sum(in ...int) (result int) {
	for _, val := range in {
		result += val
	}
	return result
}
