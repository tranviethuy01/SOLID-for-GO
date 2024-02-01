package main

import "fmt"

// Book type
type Book struct {
	Title string
	Price float64
}

// Laptop type
type Laptop struct {
	Model string
	Price float64
}

// Function to calculate total price without modification
func CalculateTotalPriceWithoutInterface(products []interface{}) float64 {
	total := 0.0
	for _, p := range products {
		switch prod := p.(type) {
		case Book:
			total += CalculateBookPriceWithoutInterface(prod)
		case Laptop:
			total += CalculateLaptopPriceWithoutInterface(prod)
		}
	}
	return total
}

// Function to calculate price for Book
func CalculateBookPriceWithoutInterface(book Book) float64 {
	return book.Price
}

// Function to calculate price for Laptop
func CalculateLaptopPriceWithoutInterface(laptop Laptop) float64 {
	return laptop.Price
}

func main() {
	// Original products
	book := Book{"The Go Programming Language", 29.99}
	laptop := Laptop{"XYZ Laptop", 999.99}

	// Calculate total price without modifying CalculateTotalPriceWithoutInterface function
	totalPrice := CalculateTotalPriceWithoutInterface([]interface{}{book, laptop})
	fmt.Printf("Total Price: $%.2f\n", totalPrice)

	// Introduce a new product type without modifying CalculateTotalPriceWithoutInterface function
	phone := PhoneWithoutInterface{"ABC Phone", 499.99}
	// Adjust the switch statement or add a new function for the new type, as needed
	totalPriceWithPhone := CalculateTotalPriceWithoutInterface([]interface{}{book, laptop, phone})
	fmt.Printf("Total Price with Phone: $%.2f\n", totalPriceWithPhone)
}

// PhoneWithoutInterface type
type PhoneWithoutInterface struct {
	Model string
	Price float64
}
/*

Total Price: $1029.98
Total Price with Phone: $1029.98

*/

/*

How bad this code 

The CalculateTotalPriceWithoutInterface function directly handles different types (Book and Laptop) using a type switch.
Separate functions (CalculateBookPriceWithoutInterface and CalculateLaptopPriceWithoutInterface) are provided for each type to calculate the price.
While this approach allows you to handle different types directly, it lacks the formal contract and benefits of interfaces. Adding new types involves modifying the CalculateTotalPriceWithoutInterface function, violating the Open/Closed Principle (OCP). Using interfaces provides a more organized and extensible way to handle varying behavior for different types.

*/
