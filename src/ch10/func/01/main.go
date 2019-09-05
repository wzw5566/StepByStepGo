package add

import "fmt"

// Add add方法传入两个int的参数
// 返回一个int的值
func Add(x, y int) int {
	return x + y
}

func main() {

	fmt.Println(Add(2,3))

}
