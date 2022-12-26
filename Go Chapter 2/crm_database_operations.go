package main

import (
	//    "fmt"
	"database/sql"
	//  "log"
	//  "net/http"
	//  "text/template"
	//  "errors"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetCustomerById(customerId int) Customer {
	var database *sql.DB
	database = GetConnection()

	var err error
	var rows *sql.Rows
	rows, err = database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerId)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(rows)
	var customer Customer
	customer = Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		err = rows.Scan(&customerId, &customerName, &SSN)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
}

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}

	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}

	defer database.Close()

	return customers
}

func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var err error
	var insert *sql.Stmt
	insert, err = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}

func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}
func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var err error
	var del *sql.Stmt
	del, err = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if err != nil {
		panic(err.Error())
	}
	del.Exec(customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}

/*func main() {

     var customers []Customer
    customers = GetCustomers()
    fmt.Println(customers)

  //  var customer Customer
//    customer.CustomerName = "Thomas Smith"
  //  customer.SSN = "2323343"

  //  InsertCustomer(customer)

  //var customer Customer
  //  customer.CustomerName = "George Thompson"
  //  customer.SSN = "23233432"
  //  customer.CustomerId = 2
  //  UpdateCustomer(customer)

var customer Customer
  //customer.CustomerName = "George Thompson"
  //customer.SSN = "23233432"
 customer.CustomerId = 2

    deleteCustomer(customer)
    customers = GetCustomers()
    fmt.Println(customers)


}
*/
