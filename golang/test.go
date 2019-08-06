package main


import (
	"fmt"
)

// 函数外部无法使用:=赋值；

var f = [10]string{}
func main() {
	a := 10
	hty, pjw := "lijian", 2
	b := [5]int{1,2,3,4,5}
	b[0] = 20
	fmt.Println("hello word", a, b, f, hty, pjw)
	consttest()
}

func consttest() { // 常量
	const(
		java = iota     // 自增值，初始为0
		php =123
		python = 234
	)
	const filename = "JSON"
	fmt.Println(filename, java, php, python)
	iftest()
}

func iftest() { // if 条件
	const sum = 10
	if sum > 20 {
		fmt.Println("if20")
	} else if sum > 10 {
		fmt.Println("if10")
	} else {
		fmt.Println("if10")
	}
	Switchtest()
}

func Switchtest() { // switch 条件有
	const sum = "xxxpz"
	switch sum{
	case "xxxpz":
		fmt.Println("switch30")
	case "lijan":
		fmt.Println("switch20")
	}
	// switch 有俩种对比的方式
	//const sum = "xxxpz"
	//switch {
	//case sum === "xxxpz":
	//	fmt.Println("switch30")
	//case sum ==="lijan":
	//	fmt.Println("switch20")
	//}
	fortest()
}

func fortest() {
	for i:= 100; i> 0; i--{
		fmt.Println(i)
	}

	//for {// 死循环
	//	fmt.Println('dsfd')
	//}
	forNtest(30)
}

func forNtest(n int) {
	for ; n> 0; n/=2{
		fmt.Println(n)
	}
	test1(10, 10)
}


// 函数

func test1(a, b int)(int,int) { // 返回多个值
	fmt.Println(xxxpz(10, 3))
	return a + b, a* b

}

// 为多个返回值起名字，仅用于简单函数

func xxxpz(a, b int)(q,r int) {
	fmt.Println(errortest(79, 29))
	q = a + b
	r = a - b
	return // 自动对号入座返回对应变量
}

// 返回错误

func errortest(a, b int)(int, error) {
	zheng()
	if a + b > 100 {
		return a + b, fmt.Errorf("%s", "error!")
	} else {
		return a + b, nil
	}

}

// 指针

// 理解 a,b  *int存的是int类型值的地址，
func zheng() {
	a,b:= 1,2
	swap_test(&a, &b)// 取到地址
	fmt.Println("ddd",a)
}

// &是取到地址
// *是指针对象

// & = value

func swap_test(a *int, b *int) {
	fmt.Println(a, *b) // *号只可以和地址（&）配合使用 // *号可以通过地址取到值
	*a=3 // *可以拿到地址中的值，然后赋值
	arraytest()
}



// 理解： a,b *int 存的是 int 类型值的地址，当对指针类型的变量 *a 时，就是取出地址对应的值


//内建容器数组
func arraytest() {
	// var arr [3]int // 初始化会变为[0, 0, 0]
	//arr := [3]int{1,2,3}
	//arr := [2]int{1,3}
	 arr := [2][4]int{ // 二维数据
		 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
		 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
	 }
	fmt.Println(arr)
	forArray()
}

func forArray() {
	arr:= [5]int{1,2,3,4,5}
	for k, v:=range(arr) {
		fmt.Println(k,v)
		//0 1  0为下标， 1为值
		//1 2
		//2 3
		//3 4
		//4 5
	}
	xxpz:= [3]string{"xxpz", "lijian", "lium"}
	printArray(xxpz)
}

 // 按值传递
func printArray(arr [3]string) {
	for k, v := range(arr) {
		fmt.Println(k,v)
		//0 xxpz
		//1 lijian
		//2 lium
	}
	qianpian()
}
// 切片

func qianpian() {
	arr := [...]int{9,8,7,6,5,4,3,2,1} // 留头去尾部
	arr1 := arr[1:2]
	arr2 := arr[:4]
	arr3 := arr[0:2]
	arr4 := arr[:]
	fmt.Println(arr1, arr2, arr3, arr4)
	arrmain()
}

// 切片的扩展

func arrmain() {
	arr := [5]int{0,1,2,3,4}
	arr1 := arr[1:3]
	fmt.Println(arr, arr1)
	updateArr(arr1)
	fmt.Println(arr1)
}

func updateArr(arr []int) { // 可以修改arr []中不写具体大小，表示是切片，引用传递
	arr[1] = 100
}

// 切片的扩展

func updateArr






