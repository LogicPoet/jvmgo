// WildcardEntry(通配符路径)实际上也是CompositeEntry，所以就不再定义新的类型了
package classpath
import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // 把路径末尾的星号去掉，得到baseDir
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir // 返回SkipDir跳过子目录(通配符类路径不能递归匹配子目录下的JAR文件)
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") { // 根据后缀名选出JAR文件
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn) // 遍历baseDir创建ZipEntry，Walk()函数的第二个参数也是一个函数
	return compositeEntry
}