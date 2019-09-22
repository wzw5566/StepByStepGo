## 指针
1.基本数据类型，变量存的就是值，就叫值类型

2.指针类型，存的是内存地址，内存指向的空间才是值
取 内存的地址：&i
fmt.Println("i memory adress is", &i )


3. 定义变量，如果变量前面带了 * 
	var m *int = &i

4.获取指针类型所指向的值，使用*

5.数组和结构体struct 都是值类型

6.指针，slice切片， map，chan，interfce都是引用类型