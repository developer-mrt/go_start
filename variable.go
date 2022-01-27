package main

import "fmt"

var job string = "Software Developer" //tanım 1 global

func main() {

	var name string = "Murat"                     //tanım 2
	var lastname = "Öngen"                        //tanım 3
	email := "murat.ongen@hotmail.com"            //tanım 4
	age, address, married := 30, "İstanbul", true //multiple değer atama tanım 5

	fmt.Println(name, lastname, email, age, address, married)
}
