package main

import "fmt"

// The difference with the "divide" function declared globally,
// when we declare it up and run the application, we get an error
// because the function "divide" hasn't been declared yet because
// we have declared it here as a variable.
func main() {
	// Declared as a variable and we're working with it
	// exactly the same way as when we declared it as a
	// function
	var divide func(float64, float64) (float64, error)

	// Initializing the variable and setting it equal to
	// an anonymous function that takes "a" and "b"
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}
