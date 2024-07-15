// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func testcount(x int) int {
// 	if x == 30 {
// 		return 0
// 	}
// 	fmt.Println(x)
// 	return testcount(x + 1)
// }
// func requestSend() {
// 	resp, err := http.Get("https://music-api-inky-xi.vercel.app/artist")
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(body))
// }
// func main() {
// 	requestSend()
// 	// testcount(1)

// }
package main

import (
 "fmt"
 "log"
 "net/http"
)

func main() {
 http.HandleFunc("/main", handler)
 log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Hello, World!")
}