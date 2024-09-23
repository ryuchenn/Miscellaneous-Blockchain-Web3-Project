//////////////////////////////////
/// 變數&函數宣告、陣列、字典、Struct、Interface
//////////////////////////////////

package main
import ( "fmt";"sort";)
// import ( "fmt";"sort"; "Tutorial/ImportTest";)

////// 函數宣告
// var=日後可重新指派的變數，const=宣告後就算重新指派也不能改變的恆定值
var x int = 10 // 可宣告資料型態 也可以像是y變數一樣不用宣告資料型態
var xx int = 5 
var y = 20.5
var yy float64 = 20.5  //GO宣告Float只有float64與float32，並不能宣告float
var yyy float32 = 20.5
var z int // 未宣告，給值為0
var a string // 未宣告，給值""
var aa string = "1" // 未宣告，給值""
var b bool // 未宣告，給值false
const Pi = 3.14159 

func main(){
	//// 0. 

	//// 1.變數宣告
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(a)
	// fmt.Println(b)

	//// 2.函數宣告
	// declare_func()

	//// %f = float， GO裡print值不能夠string與float組合之後print出來，必須將float轉為string才能print
	// fmt.Println("declare_func2(x,y):" + fmt.Sprintf("%f", declare_func2(x, y))) 
	// fmt.Printf("declare_func2(x = %d, y = %d): %f\n", x, int(y), declare_func2(x, y)) //%d = 整數

	// sum, difference := declare_func3(x, xx) //接回帶有回傳值的函數
	// fmt.Printf("declare_func3(x = %d, xx = %d, sum = %d, difference = %d)", x, xx, sum, difference) 
	
    // declare_func4()
	// fmt.Println(Person{name: "Alice", age: 30}.declare_func5("Hello"))
    
    // 其他File Function的引用: 這是『/Tutorial/ImportTest/1-1.go』的函數『SayHello』
    // !!! 使用此方法需要將程式最上方的 import ( "fmt";"sort"; ) 改成 import ( "fmt";"sort"; "Tutorial/ImportTest";)
	// ImportTest.SayHello("Alice") // 『函數』或『變數名稱』以"大"寫字母開頭，才能被其他File引用，否則則為priave function

    declare_func6()

	//// 3.ARRAY
	// declare_array()
	// declare_array2()
	// declare_array3()

	//// 4.MAP
	// declare_map()
	// declare_map2()

	//// 5. Struct
	// declare_struct()
    // declare_struct2()

    //// 6. Interface
    // declare_interface()
    // declare_interface2()
    // declare_interface3()
    // declare_interface4()
}

////////////////////////////// 函數宣告 START //////////////////////////////

// Declare_func() 與 declare_func()的差別在於 Declare_func為物件導向public概念，而declare_func()為private
func Declare_func(){
	b := 5 // 宣告b變數(此宣告方式無法用於函式外部的全域變數宣告)
	fmt.Println(b)
}

// 函數宣告1
func declare_func(){
	b := 6 
	fmt.Println(b)
}

// 函數宣告2: 帶有回傳值的function
func declare_func2(x int, y float64) float64 {
    return float64(x) + y  //將x從int型態轉型為float64傳回，且回傳型態為float64
}

// 函數宣告3: 帶有多個回傳值的function
func declare_func3(x int, y int) (int, int) {
    sum := x + y
    difference := x - y
    return sum, difference
}

////// 函數宣告4 (概念會結合Struct，先看完除了函數宣告4、5之外的所有內容後，再回來看這個)
// GO的宣告函數順序是 func (變數名稱 TYPE型別) 函數名稱(傳入值 傳入值的資料型態) 回傳的資料型態{}
// 函數宣告4: 型別函數宣告方式 (型別函數宣告 + 使用其他型別函數)
type Person struct {
    name string
    age  int
}
func (p Person) Print(){
    fmt.Printf("---%s, %d --- ", p.name, p.age)
}
func declare_func4(){  // 
    tmp := Person{name: "Alice", age: 30}
    tmp.Print()
}

// 函數宣告5: 型別函數宣告 + 回傳值
// 帶有Person屬性的p，把greeting參數傳入declare_func5後，回傳string變數
func (p Person) declare_func5(greeting string) string { 
    return greeting + ", my name is " + p.name
}

// 函數宣告6: 帶function的型別宣告
type testInt func(int) bool // 宣告了一個函式型別，此testInt能用於帶入值是int 且返回值是bool的function

func isOdd(integer int) bool {
    if integer%2 == 0 {
        return false
    }
    return true
}
func isEven(integer int) bool {
    if integer%2 == 0 {
        return true
    }
    return false
}

// 宣告的函式型別在這個地方當做了一個參數
func filter(slice []int, f testInt) []int {
    var result []int
    for _, value := range slice {
        if f(value) {
            result = append(result, value)
        }
    }
    return result
}
func declare_func6(){
    slice := []int {1, 2, 3, 4, 5, 7}
    fmt.Println("slice = ", slice)
    odd := filter(slice, isOdd)    // 函式當做值來傳遞了
    fmt.Println("Odd elements of slice are: ", odd)
    even := filter(slice, isEven)  // 函式當做值來傳遞了
    fmt.Println("Even elements of slice are: ", even)
}

////////////////////////////// 函數宣告 END //////////////////////////////

////////////////////////////// 陣列 START //////////////////////////////
//陣列 1
func declare_array(){
	var numbers [5]int     // 宣告一個整數陣列，大小為 5
    numbers[0] = 1         // 設定第一個元素的值為 1
    numbers[1] = 2         // 設定第二個元素的值為 2

    fmt.Println("陣列:", numbers) // 輸出整個陣列

	var numbers2 []int             // 不宣告陣列大小 ("大小可隨後續指派而動態改變")
    numbers2 = append(numbers2, 1)  // 使用 append 向切片中添加元素
    numbers2 = append(numbers2, 2, 3, 4)

    fmt.Println("切片:", numbers2) // 輸出切片內容
}

// 陣列 2-排序、找值
func declare_array2(){
	numbers := [5]int{4, 2, 3, 1, 5}
    sort.Ints(numbers[:]) // 將陣列轉換為切片並排序

    fmt.Println("排序後的陣列:", numbers)
}

// 陣列 3-找值
func declare_array3(){
	numbers := [5]int{4, 2, 3, 1, 5}
	target := 3 //要尋找的目標
    found := false

	for i, v := range numbers {
        if v == target {
            fmt.Printf("找到目標值 %d 在索引 %d\n", target, i)
            found = true
            break // 一旦找到，退出迴圈
        }
    }

    if !found {
        fmt.Printf("目標值 %d 未找到\n", target)
    }
}

////////////////////////////// 陣列 END //////////////////////////////

////////////////////////////// 字典 START //////////////////////////////

// 字典1
func declare_map(){
	// 宣告一個字典，鍵是字串，值是整數 => [Key: Value]
	var dict map[string]int = make(map[string]int) // make概念有點像是C#的new

    dict["age"] = 25
    dict["score"] = 90

    fmt.Println("字典:", dict)           // 輸出整個字典
    fmt.Println("年齡:", dict["age"])     // 輸出特定鍵的值

	dict = make(map[string]int) // 清空
    dict["age"] = 25
    dict["score"] = 90
	
    fmt.Println("字典的鍵：")
    for key := range dict {
        fmt.Println(key) // 輸出每個鍵
    }
}

// 字典2
func declare_map2(){
	var dict map[string]int = make(map[string]int)
    dict["age"] = 25
    dict["score"] = 90

    //// 1. 刪除特定KEY "age"
    delete(dict, "age")
    fmt.Println("刪除後的字典:", dict) // 輸出：刪除後的字典: map[score:90]

	//// 2. 刪除特定 VALUE 90 
	dict = make(map[string]int) // 清空
	dict["age"] = 25
    dict["score"] = 90
    dict["height"] = 180

    targetValue := 90
    for key, value := range dict {
        if value == targetValue {
            delete(dict, key) // 刪除找到的VALUE
        }
    }
	fmt.Println("刪除值為 90 後的字典:", dict)

	//// 3. 新增值至字典
	newKey := "weight"
    newValue := 63
    dict[newKey] = newValue

	fmt.Println("新增鍵值後的字典:", dict)
}

////////////////////////////// 字典 END //////////////////////////////

////////////////////////////// STRUCT START //////////////////////////////

// 宣告一個 Person 的 Struct => 類似於別的語言的 Class，可將多個不同類型的數據組合在一起
type Personn struct {
    name string
    age  int
}

// Struct
func declare_struct(){
	p := Personn{name: "Alice", age: 30} // 初始化一個 Person 結構體

    fmt.Println("姓名:", p.name) // 訪問結構體屬性
    fmt.Println("年齡:", p.age)
}

// Struct2: Struct裡包Struct
type Address struct {
    City    string
    Country string
}
type Person2 struct {
    Name    string
    Age     int
    Address Address  // 將 Address struct 作為欄位
}
func declare_struct2(){
	p := Person2{
        Name: "Alice",
        Age:  30,
        Address: Address{
            City:    "Toronto",
            Country: "Canada",
        },
    }
    fmt.Printf("%s lives in %s, %s.\n", p.Name, p.Address.City, p.Address.Country)
}

// Struct3: Struct裡包Struct + 型別宣告
type Skills []string
type Human struct {
    name string
    age int
    weight int
}
type Student struct {
    Human  // 匿名欄位，struct
    Skills // 匿名欄位，自訂的型別 string slice
    int    // 內建型別作為匿名欄位
    speciality string
}

func declare_struct3(){
	 // 初始化學生 Jane
     jane := Student{Human:Human{"Jane", 35, 100}, speciality:"Biology"}
     // 現在我們來存取相應的欄位
     fmt.Println("Her name is ", jane.name)
     fmt.Println("Her age is ", jane.age)
     fmt.Println("Her weight is ", jane.weight)
     fmt.Println("Her speciality is ", jane.speciality)
     // 我們來修改他的 skill 技能欄位
     jane.Skills = []string{"anatomy"}
     fmt.Println("Her skills are ", jane.Skills)
     fmt.Println("She acquired two new ones ")
     jane.Skills = append(jane.Skills, "physics", "golang")
     fmt.Println("Her skills now are ", jane.Skills)
     // 修改匿名內建型別欄位
     jane.int = 3
     fmt.Println("Her preferred number is", jane.int)
}
////////////////////////////// STRUCT END //////////////////////////////

////////////////////////////// INTERFACE START //////////////////////////////

//// Interface1 : 宣告了interface，"一定"要實作interface裡的方法。通常會結合struct。
type Speaker interface {
    Talk() //不帶回傳值
    Talk2() string //帶回傳值
}
type Dog struct {Name string}
type Person3 struct {Name string}

//// 不能用預有的資料型態，要重新指派type
// func (i int)Talk(){
//     fmt.Println("---%d---", i)
// }
// func (s string)Talk(){
//     fmt.Println("---%s---", s)
// }

// 修訂後
type MyInt int
type MyString string

// Talk()的『多型』
func (i MyInt) Talk() {
    fmt.Printf("--- %d --- \n", i)
}
func (s MyString) Talk() {
    fmt.Printf("--- %s --- \n", s)
}
func (d Dog) Talk() {
    fmt.Printf("--- %s --- \n", d.Name)
}
func (p Person3) Talk() {
    fmt.Printf("--- %s --- \n", p.Name)
}

func (d Dog)Talk2() string{
    output := "Hello, I am " + d.Name
    return output
}
func (p Person3)Talk2() string{
    output := "Hello, I am " + p.Name
    return output
}
// 此function 尚未實作interface，僅為後續作鋪墊(概念同 declare_func4())
func declare_interface(){
    tmp := MyInt(100)
    tmp.Talk()

    tmp2 := MyString("HELLO WORLD")
    tmp2.Talk()
}

// 實作interface2
func declare_interface2(){
    var s Speaker
    d := Dog{Name: "Buddy"}
    p := Person3{Name: "Alice"}

    s = d
    s.Talk()
    //// fmt.Println(s.Talk())  //不能這樣寫的原因是因為這樣會變成 fmt.Println( fmt.Printf("--- %s --- \n", d.Name) )
    fmt.Println(s.Talk2())  // 輸出: Woof! I am Buddy

    s = p
    s.Talk()
    //// fmt.Println(s.Talk())  //問題同上
    fmt.Println(s.Talk2())  // 輸出: Hello, I am Alice
}

//// Interface3 : 同時實作多個回傳值方法
type Animal interface {
    Speak() string
    Move() string
}
type Dog2 struct {Name string}
func (d Dog2) Speak() string {return "Woof!"}
func (d Dog2) Move() string {return "Runs"}

func declare_interface3(){
	var a Animal
    d := Dog2{Name: "DDOOGG"}
    a = d
    fmt.Println(a.Speak()) // 輸出: Woof!
    fmt.Println(a.Move())  // 輸出: Runs
}

// Interface4: 空interface(類似於 泛型)
func declare_interface4(){
	var x interface{} = "hello"

    switch v := x.(type) {
    case string:
        fmt.Printf("x 是 string 類型，值為 %s\n", v)
    case int:
        fmt.Printf("x 是 int 類型，值為 %d\n", v)
    default:
        fmt.Printf("未知類型 %T\n", v)  // %T 顯示實際的類型
    }

    // 延伸 (interface 無法像array、dict、struct存取多個變數，interface只能存單一值)
    x = "hello"        // 指派字串
    fmt.Printf("x 現在的值是: %v, 類型是: %T\n", x, x)

    x = 42             // 指派整數
    fmt.Printf("x 現在的值是: %v, 類型是: %T\n", x, x)

    x = 3.14           // 指派浮點數
    fmt.Printf("x 現在的值是: %v, 類型是: %T\n", x, x)

    x = true           // 指派布林值
    fmt.Printf("x 現在的值是: %v, 類型是: %T\n", x, x)
}

////////////////////////////// INTERFACE END //////////////////////////////
