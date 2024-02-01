package main

import "fmt"

// Product interface with a GetPrice method
type Product interface {
	GetPrice() float64
}

// Book type implementing the Product interface
type Book struct {
	Title string
	Price float64
}

func (b Book) GetPrice() float64 {
	return b.Price
}

// Laptop type implementing the Product interface
type Laptop struct {
	Model string
	Price float64
}

func (l Laptop) GetPrice() float64 {
	return l.Price
}

// CalculateTotalPrice function remains unchanged
func CalculateTotalPrice(products []Product) float64 {
	total := 0.0
	for _, p := range products {
		total += p.GetPrice()
	}
	return total
}

func main() {
	// Original products
	book := Book{"The Go Programming Language", 29.99}
	laptop := Laptop{"XYZ Laptop", 999.99}

	// Calculate total price without modifying CalculateTotalPrice function
	totalPrice := CalculateTotalPrice([]Product{book, laptop})
	fmt.Printf("Total Price: $%.2f\n", totalPrice)

	// Introduce a new product type without modifying CalculateTotalPrice function
	phone := Phone{"ABC Phone", 499.99}
	totalPriceWithPhone := CalculateTotalPrice([]Product{book, laptop, phone})
	fmt.Printf("Total Price with Phone: $%.2f\n", totalPriceWithPhone)
}

// Phone type implementing the Product interface
type Phone struct {
	Model string
	Price float64
}

func (p Phone) GetPrice() float64 {
	return p.Price
}



/*


Total Price: $1029.98
Total Price with Phone: $1529.97

*/


/* 
Explain
In this example:

We have the Product interface with a GetPrice method.
The Book, Laptop, and Phone types implement the Product interface, providing their specific implementations for calculating the price.
The CalculateTotalPrice function takes a slice of Product interfaces and calculates the total price without knowing the specific types of products.
With this design, you can easily introduce new types of products (e.g., Phone) without modifying the existing CalculateTotalPrice function. This demonstrates adherence to the Open/Closed Principle (OCP), as the system is open for extension (new product types can be added) without modifying existing code.



*/
