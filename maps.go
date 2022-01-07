package main
import "fmt"

//maps
var inventory = map[string] string{"product_1" : "shoes", "product_2" : "clothes"} 


//maps using make()
func print_employee_details(){
var employee_details = make(map[string]string)
employee_details["name"] = "ocip"
employee_details["years of experience"] = "3 years"
employee_details["age"]= "35"
employee_details["engineering stack"] = "backend expertise"
fmt.Printf("This is the details of the latest employer\t%v\n", employee_details)

//reference maps
 reference := employee_details
 fmt.Printf(reference["\n age"])

//updategi a map
employee_details["name"] = "dev"
fmt.Printf("\nthe employee list has been updated with a new name \t%v\n", employee_details)
//remove from map
delete(employee_details, "age")
fmt.Printf("\nthe employee list has been updated with a new name \t%v\n", employee_details)

//check for specific values
var val1, state1 = employee_details["age"]
var val2, state2 = employee_details["name"]
var _, state3 = employee_details["engineering stack"]
var val4, _ = employee_details["years of experience"]

fmt.Println(val1, state1)
fmt.Println(val2, state2)
fmt.Println(state3)
fmt.Println(val4)

}




func main(){
	fmt.Printf("List of inventory is \t%v\n", inventory)
	fmt.Printf("\n")
	print_employee_details()

}
