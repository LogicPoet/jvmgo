// 表示目录形式的类路径
package classpath
import "io/ioutil"
import "path/filepath"

// 和java语言不同，GO结构体不需要显示实现接口，只要方法匹配即可
// go没有专门的构造函数，这里统一使用new开头的函数来创建结构体实例，并把这类函数称为构造函数
type DirEntry struct {
	absDir string // 存放目录的绝对路径
}

/**
 * 构造函数
 */
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) // 把参数转换成绝对路径
	if err != nil { // 如果转换过程出现错误，则调用panic()函数终止程序执行
		panic(err)
	}
	return &DirEntry{absDir} // 否则创建DirEntry实例并返回
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className) // 先把目录和class文件名拼成一个完整的路径
	data, err := ioutil.ReadFile(fileName) // 读取class文件内容
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir // 直接返回目录
}