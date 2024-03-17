ABC Pharmacy Management System is a web-based application designed to manage invoices and product inventory for a pharmacy. It is an assesment given to me from 21c Care (https://www.21ccare.com/). This system allows users to create, view, edit, and delete invoices, as well as manage product details associated with each invoice.

## Running the System
## Frontend
## To run the frontend:

1. Navigate to the frontend directory.
2. Install dependencies using: `npm install`.
3. Start the development server using : `npm start`.
4. Access the application at http://localhost:3000 in your web browser.

## Backend
## To run the backend:

1. Navigate to the backend directory.
2. Install dependencies using `go mod tidy`.
3. Start the backend server using `go run main.go`.
4. The server will start running at http://localhost:8080.

## connect to a PostgreSQL database
1. Start PostgreSQL Server: 
2. Access PostgreSQL Shell (psql): `psql -U username -d database_name`
3. Create a Database : `CREATE DATABASE your_database_name;`
4. change connection string in main.go  
const (
	host     = "localhost"
	port     = 5432
	user     = "your user"
	password = "yur password"
	dbname   = "abc_pharmacy"
)
5. ## Enter queries
# Create the Product table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    total DECIMAL(10, 2) NOT NULL
);

# Create the Customer table
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mobile_no VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    billing_type VARCHAR(50) NOT NULL
);

 # Create the Invoice table
CREATE TABLE IF NOT EXISTS invoices (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    mobile_no VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    billing_type VARCHAR(50) NOT NULL,
    products JSONB NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL
);


## Sample Image of Backend API
![ABC_Pharmacy Rest Api Call](https://github.com/asfaqahmed/ABC-Pharmacy/blob/main/images/Screenshot%20(4).png)
![ABC_Pharmacy Front Page](https://github.com/asfaqahmed/ABC-Pharmacy/blob/main/images/Screenshot%20(5).png)


Usage
1. Creating Invoices
2. Editing Invoices
3. Deleting Invoices
4. Searching Invoices

Ensure that the backend server is running before accessing the frontend.
Make sure to configure the database connection details in the backend configuration file (config.json or .env file).
The system is designed for managing invoices and product inventory specifically for pharmacies. It may require customization to suit other types of businesses.

## Pull Requests
## welcome to pull requests from contributors to help improve the project. To submit a pull request:

1. Fork the repository to your GitHub account.
2. Clone your forked repository to your local machine.
3. Create a new branch for your changes: git checkout -b feature-name.
4. Make your changes, ensuring to follow the project's coding conventions and guidelines.
5. Commit your changes with descriptive commit messages: git commit -m "Add new feature".
6. Push your changes to your fork: git push origin feature-name.
7. Submit a pull request from your fork's branch to the main repository's main branch.
