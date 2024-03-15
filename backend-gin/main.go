package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Root"
	dbname   = "abc_pharmacy"
)

var db *sql.DB

type Item struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	UnitPrice    float64 `json:"unit_price"`
	ItemCategory string  `json:"item_category"`
}

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total"`
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
	Name        string    `json:"name"`
	MobileNo    string    `json:"mobile_no"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	BillingType string    `json:"billing_type"`
	Products    []Product `json:"products"`
	TotalAmount string    `json:"total_amount"`
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
	router.Use(errorLogger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Add your React app's origin
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// API routes for items
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.PUT("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)

	// API routes for customers
	router.GET("/customers", getCustomers)
	router.POST("/customers", addCustomer)
	router.PUT("/customers/:id", updateCustomer)
	router.DELETE("/customers/:id", deleteCustomer)

	// API routes for invoices
	router.GET("/invoices", getInvoices)
	router.POST("/invoices", addInvoice)
	router.PUT("/invoices/:id", updateInvoice)
	router.DELETE("/invoices/:id", deleteInvoice)

	// Run the server
	router.Run(":8080")
}

func errorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Next handler
		c.Next()

		// Check if any errors occurred during handling
		if len(c.Errors) > 0 {
			// Log errors
			log.Println(c.Errors.String())
		}
	}
}

// CRUD operations for items

func getItems(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.UnitPrice, &item.ItemCategory)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

func addItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO items (name, unit_price, item_category) VALUES ($1, $2, $3)", newItem.Name, newItem.UnitPrice, newItem.ItemCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added successfully"})
}

func updateItem(c *gin.Context) {
	id := c.Param("id")

	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE items SET name=$1, unit_price=$2, item_category=$3 WHERE id=$4", updatedItem.Name, updatedItem.UnitPrice, updatedItem.ItemCategory, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
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
		err := rows.Scan(&customer.ID, &customer.Name, &customer.MobileNo, &customer.Email, &customer.Address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
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

// CRUD operations for invoices

func getInvoices(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM invoices")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var invoice Invoice
		var id int
		var name, mobileNo, email, address, billingType string
		if err := rows.Scan(&id, &name, &mobileNo, &email, &address, &billingType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		invoice.ID = id
		invoice.Name = name
		invoice.MobileNo = mobileNo
		invoice.Email = email
		invoice.Address = address
		invoice.BillingType = billingType

		// Fetch products for the invoice
		products, err := getProductsForInvoice(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		invoice.Products = products

		invoices = append(invoices, invoice)
	}

	c.JSON(http.StatusOK, invoices)
}

func addInvoice(c *gin.Context) {
	var newInvoice Invoice
	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert invoice
	var id int
	err := db.QueryRow("INSERT INTO invoices (name, mobile_no, email, address, billing_type) VALUES ($1, $2, $3, $4, $5) RETURNING id", newInvoice.Name, newInvoice.MobileNo, newInvoice.Email, newInvoice.Address, newInvoice.BillingType).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert products for the invoice
	for _, product := range newInvoice.Products {
		_, err := db.Exec("INSERT INTO invoice_products (invoice_id, name, quantity, unit_price, total) VALUES ($1, $2, $3, $4, $5)", id, product.Name, product.Quantity, product.UnitPrice, product.Total)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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

	// Update invoice
	_, err := db.Exec("UPDATE invoices SET name=$1, mobile_no=$2, email=$3, address=$4, billing_type=$5 WHERE id=$6", updatedInvoice.Name, updatedInvoice.MobileNo, updatedInvoice.Email, updatedInvoice.Address, updatedInvoice.BillingType, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update products for the invoice
	for _, product := range updatedInvoice.Products {
		_, err := db.Exec("UPDATE invoice_products SET name=$1, quantity=$2, unit_price=$3, total=$4 WHERE invoice_id=$5 AND id=$6", product.Name, product.Quantity, product.UnitPrice, product.Total, id, product.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully"})
}

func deleteInvoice(c *gin.Context) {
	id := c.Param("id")

	// Delete products for the invoice
	_, err := db.Exec("DELETE FROM invoice_products WHERE invoice_id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete invoice
	_, err = db.Exec("DELETE FROM invoices WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}

func getProductsForInvoice(invoiceID int) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM invoice_products WHERE invoice_id=$1", invoiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.UnitPrice, &product.Total); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
