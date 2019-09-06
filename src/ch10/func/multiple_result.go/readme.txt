# 函数可以返回多个返回值，在外部调用的时候，也需要设置相应的
变量来接收对应的返回值


// swap输入两个strng 参数，返回 两个字符串参数
func swap(x, y string) (string, string) {
	return y, x
}
