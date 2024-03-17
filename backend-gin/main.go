package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// while send the code empty the values
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Root"
	dbname   = "abc_pharmacy"
)

var db *sql.DB

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total"`
}

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MobileNo    string `json:"mobile_no"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	BillingType string `json:"billing_type"`
}

type Invoice struct {
	ID          int       `json:"id"`
	Name        string    `json:"customer_name"`
	MobileNo    string    `json:"mobile_no"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	BillingType string    `json:"billing_type"`
	Products    []Product `json:"products"`
	TotalAmount float64   `json:"total_amount"`
}

func main() {
	// router
	router := gin.Default()

	// Connect to PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Use error logging middleware

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add your React app's origin
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// API routes for customers
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getSingleCustomer)
	router.POST("/customers", addCustomer)
	router.PUT("/customers/:id", updateCustomer)
	router.DELETE("/customers/:id", deleteCustomer)

	// API routes for products
	router.GET("/products", getProducts)
	router.GET("/products/:id", getSingleProduct)
	router.POST("/products", addProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	// API routes for invoices
	router.GET("/invoices", getInvoice)
	router.GET("/invoices/:id", getSingleInvoice)
	router.POST("/invoices", addInvoice)
	router.PUT("/invoices/:id", updateInvoice)
	router.DELETE("/invoices/:id", deleteInvoice)

	// Run the server
	router.Run(":8080")
}

// CRUD operations for customers
func getCustomers(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.MobileNo, &customer.Email, &customer.Address, &customer.BillingType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

func getSingleCustomer(c *gin.Context) {
	customerID := c.Param("id")

	row := db.QueryRow("SELECT * FROM customers WHERE id = $1", customerID)

	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.MobileNo, &customer.Email, &customer.Address, customer.BillingType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the invoice as JSON response
	c.JSON(http.StatusOK, customer)
}

func addCustomer(c *gin.Context) {
	var newCustomer Customer
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO customers (name, mobile_no, email, address, billing_type) VALUES ($1, $2, $3, $4, $5)", newCustomer.Name, newCustomer.MobileNo, newCustomer.Email, newCustomer.Address, newCustomer.BillingType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer added successfully"})
}

func updateCustomer(c *gin.Context) {
	id := c.Param("id")

	var updatedCustomer Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE customers SET name=$1, mobile_no=$2, email=$3, address=$4, billing_type=$5 WHERE id=$6", updatedCustomer.Name, updatedCustomer.MobileNo, updatedCustomer.Email, updatedCustomer.Address, updatedCustomer.BillingType, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM customers WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// CRUD operations for products
func getProducts(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.UnitPrice, &product.TotalPrice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func getSingleProduct(c *gin.Context) {
	productID := c.Param("id")

	row := db.QueryRow("SELECT * FROM products WHERE id = $1", productID)

	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Quantity, &product.UnitPrice, &product.TotalPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the invoice as JSON response
	c.JSON(http.StatusOK, product)
}

func addProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO products (name, quantity, unit_price, total_price) VALUES ($1, $2, $3, $4)", newProduct.Name, newProduct.Quantity, newProduct.UnitPrice, newProduct.TotalPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")

	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE products SET name=$1, quantity=$2, unit_price=$3, total_price=$4 WHERE id=$5", updatedProduct.Name, updatedProduct.Quantity, updatedProduct.UnitPrice, updatedProduct.TotalPrice, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM products WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// CRUD operations for invoices
func getInvoice(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM invoices")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var invoice Invoice
		var productsJSON string
		if err := rows.Scan(&invoice.ID, &invoice.Name, &invoice.MobileNo, &invoice.Email, &invoice.Address, &invoice.BillingType, &productsJSON, &invoice.TotalAmount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Deserialize productsJSON into the Products slice of Invoice
		if err := json.Unmarshal([]byte(productsJSON), &invoice.Products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		invoices = append(invoices, invoice)
	}

	c.JSON(http.StatusOK, invoices)
}

func getSingleInvoice(c *gin.Context) {
	invoiceID := c.Param("id")

	row := db.QueryRow("SELECT * FROM invoices WHERE id = $1", invoiceID)

	var invoice Invoice
	var productsJSON string

	err := row.Scan(&invoice.ID, &invoice.Name, &invoice.MobileNo, &invoice.Email, &invoice.Address, &invoice.BillingType, &productsJSON, &invoice.TotalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Deserialize productsJSON into the Products slice of Invoice
	if err := json.Unmarshal([]byte(productsJSON), &invoice.Products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the invoice as JSON response
	c.JSON(http.StatusOK, invoice)
}

func addInvoice(c *gin.Context) {
	var newInvoice Invoice
	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productsJSON, err := json.Marshal(newInvoice.Products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize products"})
		return
	}

	_, err = db.Exec("INSERT INTO invoices (customer_name, mobile_no, email, address, billing_type, products, total_amount) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		newInvoice.Name, newInvoice.MobileNo, newInvoice.Email, newInvoice.Address, newInvoice.BillingType, productsJSON, newInvoice.TotalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Invoice created successfully"})
}

func updateInvoice(c *gin.Context) {
	id := c.Param("id")

	var updatedInvoice Invoice
	if err := c.ShouldBindJSON(&updatedInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productsJSON, err := json.Marshal(updatedInvoice.Products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize products"})
		return
	}

	_, err = db.Exec("UPDATE invoices SET customer_name=$1, mobile_no=$2, email=$3, address=$4, billing_type=$5, products=$6, total_amount=$7 WHERE id=$8",
		updatedInvoice.Name, updatedInvoice.MobileNo, updatedInvoice.Email, updatedInvoice.Address, updatedInvoice.BillingType, productsJSON, updatedInvoice.TotalAmount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully"})
}

func deleteInvoice(c *gin.Context) {
	id := c.Param("id")

	// Delete invoice
	_, err := db.Exec("DELETE FROM invoices WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}
