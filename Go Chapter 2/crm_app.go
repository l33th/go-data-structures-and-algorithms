package main

import (
	"fmt"
	"net/http"
	"text/template"
	//    "errors"
	"log"
)

var templateHtml = template.Must(template.ParseGlob("templates/*"))

func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	log.Println(templateHtml.ExecuteTemplate(writer, "Home", customers))
}

func Create(writer http.ResponseWriter, request *http.Request) {
	//  var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	log.Println(templateHtml.ExecuteTemplate(writer, "Create", nil))
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	//  var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	log.Println(templateHtml.ExecuteTemplate(writer, "Home", customers))
}
func Alter(writer http.ResponseWriter, request *http.Request) {
	//  var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	var customer Customer
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Scanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	UpdateCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	log.Println(templateHtml.ExecuteTemplate(writer, "Home", customers))
}

func Update(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	//log.Println(customer)
	//var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	log.Println(templateHtml.ExecuteTemplate(writer, "Update", customer))
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	DeleteCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	log.Println(templateHtml.ExecuteTemplate(writer, "Home", customers))
}

func View(writer http.ResponseWriter, request *http.Request) {
	//var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	fmt.Println(customer)
	var customers []Customer
	customers = []Customer{customer}
	//  customers.append(customer)
	log.Println(templateHtml.ExecuteTemplate(writer, "View", customers))
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	//  var template_html *template.Template
	//template_html = template.Must(template.ParseFiles("main.html"))
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	log.Fatalln(http.ListenAndServe(":8000", nil))
