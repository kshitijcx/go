package main //compiler looks for main function in main package to run
import "fmt"
func main(){
	fmt.Println("hello world")
}
//6 points to remember
//statically typed language
//strongly typed lang -> operations depend upon datatype
//GO is compiled
//Fast compile time
//Built in concurrency -> goroutines
//simplicity (garbage collection)
 
//packages -> collection of go files
//modules -> collection of packages

//go mod init mod_name -> creates go.mod file (like package.json)
//go build path -> creates bin file that we can run
//go run path -> compile and run in single command