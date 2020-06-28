// 复合目录路径
package classpath
import "errors"
import "strings"

type CompositeEntry []Entry // CompositeEntry由更小的Entry组成

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) { // 把参数（路径列表）按分隔符分成小路径
		entry := newEntry(path) // 把每个小路径都转换成具体的Entry实例
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self { // 遍历路径
		data, from, err := entry.readClass(className) // 调用每一个子路径的readClass
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className) // 如果遍历完所有的子路径还没有找到class文件，则返回错误
}
	
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self { // 遍历调用子路径的String() 方法
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator) // 把得到的字符串用路径分隔符拼接起来
}