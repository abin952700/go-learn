package calc

//go mod init 初始化项目

func Add(x, y int) int { //首字母大小写表示共私有方法(私有当前包)
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
