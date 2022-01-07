package main
import ("fmt")


//structs

/* 
create the Struct,
create an instance of the struct
a function that takes a structures(struct) & prints the instance of the struct
 */

type software_eng struct{
	name string
	stack string
	years_of_expertise float64
	age int
	salary int
}

type Student struct {
	name string 
	student_id uint
	age int
	major string
	level int
  }

func main(){

	var eng_details software_eng
	//eng details
	eng_details.name = "ocip"
	eng_details.stack = "backend"
	eng_details.years_of_expertise = 2.5
	eng_details.age = 30
	eng_details.salary = 100000
	fmt.Print("what's ocip salary>\n", "$", eng_details.salary)
	
	print_details(eng_details)

	var student_1 Student 
	student_1.name = "kay"
	student_1.student_id = 001
	student_1.age = 32
	student_1.major = "electronics engineering"
	student_1.level = 300

	print_students_details(student_1)
}

func print_students_details(student_id Student){
	fmt.Println("name :" , student_id.name)
	fmt.Println("id :", student_id.student_id)
	fmt.Println("age :", student_id.age)
	fmt.Println("major :", student_id.major)
	fmt.Println("level :", student_id.level)


}



func print_details(skills software_eng) {
	fmt.Println("name :", skills.name)
	fmt.Println("stack :", skills.stack)
	fmt.Println("years of experience :", skills.years_of_expertise)
	fmt.Println("age :", skills.age)
	fmt.Print("salary :","$", skills.salary)
	}
	

/*
var pers1 Person
var pers2 Person

// Pers1 specification
pers1.name = "Hege"
pers1.age = 45
pers1.job = "Teacher"
pers1.salary = 6000

// Pers2 specification
pers2.name = "Cecilie"
pers2.age = 24
pers2.job = "Marketing"

pers2.salary = 4500

// Print Pers1 info by calling a function
printPerson(pers1)

// Print Pers2 info by calling a function
printPerson(pers2)

func printPerson(pers Person) {
fmt.Println("Name: ", pers.name)
fmt.Println("Age: ", pers.age)
fmt.Println("Job: ", pers.job)
fmt.Println("Salary: ", pers.salary)

}
*/