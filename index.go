package main
import (
	"fmt"
	
)

// const data = 2
// func main() {
// // 	var student1 string
// //   student1 = "John"
//     // var a, b int = 10, 3
// 	var arr1 = [...]int{1 ,2,3}
// 	if len(arr1) > 1{
// 		fmt.Println(arr1)
// 	}
// }
type Person struct {
	name string
	age int
	job string
	salary int
  }
  func getArray(){
	var a = map[string]string{"brand":"Ford", "model":"Mustang"}
	
	fmt.Printf("%v\n", a["brand"])

  }
  
func main(){
	getArray()
	// var person1 Person
	// person1.name="Alisher"
	// person1.age=19
	// person1.job = "Developer"
	// person1.salary = 6

	// fmt.Println("Name: ", person1.name)
	// fmt.Println("Age: ", person1.age)
	// fmt.Println("Job: ", person1.job)
	// fmt.Println("Salary: ", person1.salary)

	// var count int ;
	// count = 4;
	// for i := 0; i < count; i++ {
	// 	fmt.Println("salom" + strconv.Itoa(i))
	// }
}
