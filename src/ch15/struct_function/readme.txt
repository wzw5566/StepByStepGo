# 结构体的方法

type 结构体的名称 struct {
  field1 type
  field2 type
}

## 举例定义一个人的结构体，有名字和年龄的字段
type Person struct {
  Name string
  Aget int
}

## 人有个可以说话的方法speak

给结构体Person，绑定一个Speak的方法
func (person Person) Speak() {
	fmt.Println("hello", person.Name)
}

person为形参，可以随意定义

结构体方法的定义：

func (recevier type) methodName(参数列表) (resl int, resl2 int){
  方法体
  return 返回值

}