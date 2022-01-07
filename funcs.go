package main
import "fmt"

//function aka func
//func without param
func first_function(){
	fmt.Println("Hurray!! i just wrote my  Go first function")
}
//func with param
func welcome_message(Tname string) {
	fmt.Print("my twitter username is ", Tname,".")
}
//func with multiple param
func lang_experience_year(lang string, years_of_experience float64){
	fmt.Println("I've been",lang," dev for",years_of_experience,"years now.")
}
//func with return value
func compute_age(current_year int, year_born int ) int{
	var my_age = current_year - year_born
	return my_age
}
//func with multiple return value
func birthday_message(age int, message string) (_age int, _message string){
	fmt.Println("\nbirthday message")
	_age = age - 5 
	_message = message + "\n have fun!" 
	return

}
//assigning func to variable names
func graduation_message(set int, student_name string)(_set int, _message string){
	fmt.Println("Congratulations")
	_set = set
	_message = student_name + "Keep soaring!"
	return
}

//Almighty RECURSION function!!!!!
/* function name, arguments,
return type, 
break/run condition to check for,
statement(action to be performed)
statement calling the function itself()
retuen value
*/
//first recursive func
func my_recursive_prog(x int) int{
	if x==10{
		return 0
	} 
	fmt.Println(x)
	return my_recursive_prog(x + 1)
}

func factorial_calculator(x int) (y int){
	if x > 0{
		y = x * factorial_calculator(x-1)
	
	} else {
		y = 1
	}
	return y
}

//structs










func main (){
	first_function()
	welcome_message("ocip_dev")
	lang_experience_year("python", 4)
	lang_experience_year("solidity", 2)
	lang_experience_year("Go", 0.3)
	lang_experience_year("rust", 0)
	fmt.Print(compute_age(2020, 1992))
	fmt.Println(birthday_message(20, "Happy birthday buddy!"))
	a, b := graduation_message(2012, "\nHey ocip,")
	fmt.Println(a, b)
	my_recursive_prog(4)
	fmt.Println(factorial_calculator(8))
	
}