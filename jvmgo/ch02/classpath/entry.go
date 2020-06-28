//可以把类路径想象成一个大的整体，它由启动类路径、扩展类路径和用户类路径三个小路径构成。三个小路径又分别由更小的路径构成。
//是不是很像组合模式（composite pattern）？ 没错，这里就套用组合模式来设计和实现类路径。
package classpath
import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator) // 常量，存放路径分隔符

type Entry interface {
	/**
	 * 寻找和加载class文件
	 * className: class文件的相对路径，路径之间用斜线(/)分隔，文件名有.class后缀
	 * return: []byte：读取到的字节数据
	 *		   Entry：最终定位到class文件的Entry
	 *		   error:错误信息
	 */
	readClass(className string) ([]byte, Entry, error)
 
	/**
	 * 相当java的toString(),用于返回变量的字符串表示
	 */
	String() string
}

/**
 * 根据参数创建不同类型的Entry实例
 * path: class文件的相对路径，路径之间用斜线(/)分隔，文件名有.class后缀
 */
func newEntry(path string) Entry {
	if string.Contains(path, pathListSeparator) {
		return newCompostieEntry(path)
	}
	if stings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	   strings.HasSuffix("path", ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}