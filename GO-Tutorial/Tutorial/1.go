package main
import "fmt" // 用於做基本輸入輸出
import (
    "Tutorial/ImportTest" // 這裡的路徑取決於你的模組路徑，通常根據你的模組名稱決定，可參考以下『簡易指令教學』
)
/*
簡易指令教學:
- 打開Visual Studio Code -> Topbar看到 View -> Terminal -> cd到所在資料 
-> 初始化設定 可參考 (https://yuminlee2.medium.com/golang-import-local-packages-cbb2201c72e1)
-> go mod init 起始資料夾名稱 (EX:go mod init Tutorial)
-> module 起始資料夾名稱 (EX:module Tutorial)
-> go build 檔名.go (EX: go build 1.go )
-> 執行檔案: 輸入『./1』 -> 『./』=所在路徑 『1』=檔名

- 其他指令 / 衍生指令: 
	1. 『go build 1.go 2.go』:同時compile 1.go 和 2.go
	2. 『go run 1.go 2.go』:同時執行 1.go 和 2.go
*/

func main(){
	fmt.Println("fmt.Println: Hello World"); // Println會自動換行，且只能丟字串、數值
	fmt.Println(123);

	print("print: Hello World")
	print(456)

	fmt.Printf("789"); //Printf 不會自動換行，且只能丟字串
	fmt.Printf("456");
	// fmt.Printf(123); // 執行這行會過不了build，原因是Printf只能丟字串

	fmt.Println("---Import from Other Folder or File---")
	ImportTest.SayHello("Alice") // 這是『/Tutorial/ImportTest/1-1.go』的函數『SayHello』，需要在程式頭部輸入 import("Tutorial/ImportTest") 才能使用
	// ImportTest.sayHello("Alice") // 這是『/Tutorial/ImportTest/1-1.go』的函數『sayHello』，執行這行會build不過，原因是『函數』或『變數名稱』以"大"寫字母開頭，這樣它們才是「公有」的，才可以被其他File訪問
	
	// declare_func() // 這是2.go的程式，要用此函數必須使用『go run 1.go 2.go』指令或是
	
}
