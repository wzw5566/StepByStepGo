

# Map 元素的访问

与其他语言的差异：

在访问的key不存在时，仍会返回0值，不能通过返回nil，来判断是否存在

# Map 的遍历  range方法

	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k,v:= range m1{
		t.Log(k,v)
	}

  - Map的 value是一个
  
  //定义一个value是一个方法func(op int)，返回值是int的map
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }