package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

// MongoDB-BSON简介
/*
MongoDB中的JSON文档存储在名为BSON(二进制编码的JSON)的二进制表示中。与其他将JSON数据存储为简单字符串和数字的数据库不同。BSON编码扩展了JSON表示，使其包含额外的类型，如int、long、date、浮点数和decimal128.这使得应用程序更容易可靠地处理、排序和比较数据。

连接MongoDB的Go驱动程序中有两大类型表示BSON数据：D和Raw

类型D家族被用来简洁地构建使用本地Go类型的BSON对象。这对于构造传递给MongoDB的命令特别有用。D家族包括四类：
 D: 一个BSON文档。这种
类型应该在顺序重要的情况下使用，比如MongoDB命令。
 M： 一张无序的map。它和D是一样的，只是它不保持顺序。
 A: 一个BSON数据。
 E: D里面的一个元素。
要使用BSON,需要先导入下面的包：
import `go.mongodb.org/mongo-driver/bson`

下面是一个使用D类型构建的过滤器文档的例子，它可以用来查找name字段与"张三"或"李四"匹配的文档：
bson.D({
  "name",
  bson.D{{
    "$in",
     bson.A("张三","李四")
  }}
})
Raw类型家族用于验证字节切片。你还可以使用Lookup()从原始类型检索单个元素。如果你不想要将BSON反序列化成另一种类型的开销，那么这是非常有用的，
*/

func main() {
	// filter 过滤器 sql where
	d := bson.D{{"name", "tom"}}
	fmt.Printf("d: %v\n", d)
}
