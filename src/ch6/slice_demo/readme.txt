切片的声明：

var 切片名 []类型
var MySlice []int64

定义切片还不能马上使用，必须使用make或指定长度或容量才可以

1.指定长度
var Myslice [3]float64

2.使用make分配内存空间,5为长度，容量为10
var Myslice2 []int64 = make([]float64,5,10)

切片内的如果没有值，则为类型的默认值，比如 int 默认值是0

