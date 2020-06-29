// 表示ZIP或JAR文件形式的类路径
package classpath
import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string // 存放ZIP或JAR文件的绝对路径
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) // 尝试打开ZIP文件
	if err != nil {
		return nil, nil, err
	}
	defer r.Close() // 确保打开的文件得以关闭
	for _, f := range r.File { // 遍历zip压缩包里的文件,寻找class文件
		if f.Name == className { // 当找到以.class结尾的文件时
			rc, err := f.Open() // 尝试打开这个class文件
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc) // 读取class文件里的内容
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil // 返回读取结果
		}
	}
	return nil, nil, errors.New("class not found: " + className) // 未找到class文件
}

func (self *ZipEntry) String() string {
	return self.absPath
}