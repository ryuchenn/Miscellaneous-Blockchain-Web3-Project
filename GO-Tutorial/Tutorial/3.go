


//////////////////////////////////
/// 迴圈、if
//////////////////////////////////

package main
import ( "fmt";)


func main(){
	//// for 
	// declare_for()
	// declare_for2()
	// declare_for3()
	// declare_for4()
	// declare_for5()

	//// if
	// declare_if()
	// declare_if2()
	// declare_if3()
	// declare_if4()
	// declare_if5()
	// declare_if6()
	declare_if7()

}

////////////////////////////// for START //////////////////////////////

func declare_for(){
	for i := 0; i < 5; i++ {
        fmt.Println("迴圈次數:", i)
    }
}

// for 2: 結合 array
func declare_for2(){
	arr := []int{1, 2, 3, 4, 5}

    for index, value := range arr{
		fmt.Printf("索引: %d, 值: %d\n", index, value)
	}
}

// for 3: 結合 dict
func declare_for3(){
	m := map[string]int{
        "Apple":  5,
        "Banana": 2,
        "Orange": 7,
    }

    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}

// for 4: 結合 Array&struct
type Person struct {
    Name string
    Age  int
}
func declare_for4(){
    people := []Person{
        {"Alice", 30}, //不用像dict宣告key
        {"Bob", 25},
        {"Charlie", 35},
    }

    // 使用 for range 遍歷Array
    for index, person := range people { //索引&Value
        fmt.Printf("index: %d, 名字: %s, 年齡: %d\n", index, person.Name, person.Age)
    }
}

// for5: 結合 字典&struct
type Product struct {
    Name  string
    Price int
}
func declare_for5(){
    products := map[string]Product{
        "Laptop":  {"Laptop", 1000}, //要宣告key&value
        "Phone":   {"Phone", 500},
        "Tablet":  {"Tablet", 300},
	}

     // 使用 for range 迴圈遍歷 map
	 for key, product := range products {//Key&Value
        fmt.Printf("產品類型: %s, 名稱: %s, 價格: %d\n", key, product.Name, product.Price)
    }
}
////////////////////////////// for END //////////////////////////////


////////////////////////////// if START //////////////////////////////
// GO的IF判斷是可以不用括號


func declare_if(){
	num := 10

    if num > 5 {
        fmt.Println("num 大於 5")
    } else if num == 5 {
        fmt.Println("num 等於 5")
    } else {
        fmt.Println("num 小於 5")
    }
}

// if2: 使用 for 循環遍歷陣列，並使用 if 判斷是否找到值
func declare_if2(){
	arr := []int{1, 2, 3, 4, 5}
    value := 3

    
    found := false
    for _, v := range arr {
        if v == value {
            found = true
            break
        }
    }

    if found {
        fmt.Printf("找到值: %d\n", value)
    } else {
        fmt.Printf("沒有找到值: %d\n", value)
    }
}

// if3: 判斷索引值2的選項是否是2
func declare_if3(){
    arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(arr): %d \n", len(arr))
	
	if (len(arr) > 0){
		if (arr[2] == 2) {
			fmt.Println("索引 2 的值是 2")
		} else {
			fmt.Println("索引 2 的值不是 2")
		}
	}
    
}


// if4: 搭配struct
type Person2 struct {
    Name string
    Age  int
}
func declare_if4(){
	p := Person2{"Alice", 30}

    // 使用 if 判斷結構體的欄位
    if p.Age >= 18 {
        fmt.Printf("%s 是成年人\n", p.Name)
    } else {
        fmt.Printf("%s 是未成年人\n", p.Name)
    }
}

// if5: dict搭配exists判斷是否存在
func declare_if5(){
	m := map[string]int{
        "apple":  5,
        "banana": 3,
    }
    key := "banana"

    // 使用 if 搭配字典進行鍵是否存在的判斷
	/*
		『if ......; exists {}』：這是一個複合語句，概念同SQL的if exists語句。
			『if value, exists := m[key]; exists』
			- 1. 前半部分初始化了 value 和 exists。 
			- 2. 後半部分則是條件判斷。如果 exists 為 true，就會進入 if 區塊。
-			- 3. m[key]：這是從字典 m 中根據 key 取得對應的值。如果 key 存在，True就返回該值，False就返回字典值類型的零值。
			- 4. value：這是對應於 key 的值，如果 key 存在，它會存放該鍵的值。
			- 5. exists：這是一個布林值，表示鍵是否存在。如果 key 存在於字典中，exists 的值為 true，否則為 false。
			- 總結: 語句就是 => if value, bool := dict[key]; exists{}
	*/
    if value, exists := m[key]; exists {
        fmt.Printf("鍵 %s 存在，值為: %d\n", key, value)
    } else {
        fmt.Printf("鍵 %s 不存在\n", key)
    }
}

// if6: 延伸if5
type Product2 struct {
    Name  string
    Price int
}
func declare_if6(){
    products := map[string]Product2{
        "laptop": {"Laptop", 1000},
        "phone":  {"Phone", 500},
    }

    // 判斷字典中是否存在某個產品
    if product, exists := products["laptop"]; exists {
        // 如果產品價格超過 800，print高價產品
        if product.Price > 800 {
            fmt.Printf("高價產品: %s，價格: %d\n", product.Name, product.Price)
        }
    } else {
        fmt.Println("產品不存在")
    }
}

// if7: GO的if裡可以宣告一個區域變數
func returnValue() int{
	return 11;
}
func declare_if7(){
    if x := returnValue(); x > 10 {
		fmt.Printf("%d is greater than 10", x)
	} else {
		fmt.Printf("%d is less than 10", x)
	}
}

////////////////////////////// if END //////////////////////////////



