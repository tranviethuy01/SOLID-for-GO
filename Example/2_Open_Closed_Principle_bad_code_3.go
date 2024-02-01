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
func CalculateTotalPriceWithoutInterface(products []Book, laptops []Laptop) float64 {
	total := 0.0
	for _, book := range products {
		total += CalculateBookPriceWithoutInterface(book)
	}
	for _, laptop := range laptops {
		total += CalculateLaptopPriceWithoutInterface(laptop)
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
	totalPrice := CalculateTotalPriceWithoutInterface([]Book{book}, []Laptop{laptop})
	fmt.Printf("Total Price: $%.2f\n", totalPrice)

	// Introduce a new product type without modifying CalculateTotalPriceWithoutInterface function
	phone := Phone{"ABC Phone", 499.99}
	fmt.Println("new product type Phone", phone)
	// Adjust the function signature and body for the new type, as needed
	//totalPriceWithPhone := CalculateTotalPriceWithoutInterface([]Book{book}, []Laptop{laptop}, []Phone{phone})
	totalPriceWithPhone := CalculateTotalPriceWithoutInterface([]Book{book}, []Laptop{laptop})
	fmt.Printf("Total Price with Phone: $%.2f\n", totalPriceWithPhone)
}

// Phone type
type Phone struct {
	Model string
	Price float64
}

/*


Total Price: $1029.98
new product type Phone {ABC Phone 499.99}
Total Price with Phone: $1029.98

*/


/*

How bad this code:

The CalculateTotalPriceWithoutInterface function now takes slices of concrete types (Book and Laptop) directly instead of using interface{}.
Separate functions (CalculateBookPriceWithoutInterface and CalculateLaptopPriceWithoutInterface) are provided for each type to calculate the price.
Introducing a new type involves adjusting the function signature and body for the new type.
While this approach directly handles concrete types, it lacks the benefits of abstraction and formal contracts provided by interfaces. Adding new types still involves modifying the existing function, which violates the Open/Closed Principle (OCP). Using interfaces remains a more organized and extensible approach.

*/
