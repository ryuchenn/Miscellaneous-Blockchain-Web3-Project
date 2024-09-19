// mypackage.go
package ImportTest

import "fmt"

// 定義一個公有函數 => 其他程式語言的 public概念
func SayHello(name string) {
    fmt.Println("Here is the function from 「\\ImportTest\\1-1.go」, Name:" + name)
}

// 定義一個公有函數 => 其他程式語言的 private概念
func sayHello(name string) {
    fmt.Println("sayHello: " + name)
}