package main
import ("fmt")
func testcount(x int) int {
	if x == 30 {
	  return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
  }
func main() {
	testcount(1)
	//  array_name := []string{"salom","name","kino"}

	// for i:=0; i < len(array_name); i++ {
	// 	fmt.Println(array_name[i])
	//   }
}	