package main

import("fmt")

func main() {
    fmt.Println("Hello World!")

    // Define a list of integers
    numbers := []int{10, 20, 30, 40, 50}

    // Variable to store the sum
    sum := 0

    // Loop through the numbers to calculate the sum
    for _, number := range numbers {
        sum += number
    }

    // Print the sum
    fmt.Println("The sum is:", sum)
}
