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
		total += book.Price
	}
	for _, laptop := range laptops {
		total += laptop.Price
	}
	return total
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


How bad this code 

The CalculateTotalPriceWithoutInterface function directly calculates the total price by accessing the Price field of each concrete type.
Introducing a new type (Phone) involves manually adjusting the function signature and body to include the new type.
While this code may be easier to understand, it becomes less maintainable as new types are introduced. This approach violates the Open/Closed Principle (OCP), as adding new types requires modifications to existing functions. Interfaces provide a more modular and organized solution, even though they were omitted in this particular example.

*/

