package main

//重写resverse函数，使用数组指针代替slice

func resvers2(ss *[6]int) {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
}

func main() {
	//ss := [6]int{0, 1, 2, 3, 4, 5}
	//resvers2(&ss)
	// 也可以受用以下方法实现数组内元素反转
	//var s = make([]int, 0)
	//copy(s, ss[:])
}

//bytes.Buffer 是GO语言标准库中的一个类型，用于高效的处理字节切片（[]byte）。它提供了许多方便的方法来读写字节数据，适用于需要频繁对字节数据进行拼接、截取、转换等场景的操作

func buffer1() {
	// 创建一个新的bytes.Buffer
	//var buffer byte.Buffer
	// 写入字符串数据
	//buffer.WriteString("Hello,")
	// 写入字节数据

}
