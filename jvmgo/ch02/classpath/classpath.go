// classpath结构体
package classpath
import "os"
import "path/filepath"

type Classpath struct {
	bootClasspath Entry // 引导类路径
	extClasspath Entry	// 扩展类路径
	userClasspath Entry // 用户类路径
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption) //使用-Xjre选择解析启动类路径和扩展类路径
	cp.parseUserClasspath(cpOption) //使用-classpath/-cp选择解析用户类路径
	return cp
}

/**
 * 读取class文件
 * className： 接收的类名不包含".class后缀"
 */
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil { // 先从启动类路径搜索class文件
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil { // 再从扩展类路径搜索class文件
		return data, entry, err
	}
	return self.userClasspath.readClass(className) // 最后从用户类路径搜索class文件
}

/**
 * 返回用户类路径的字符串表示
 */
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

/**
 * 解析启动类路径和扩展类路径
 */
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption) // 获取jre目录
	
	// 路径：jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// 路径：jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/**
 * 解析用户类路径
 */
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" { // 如果用户没有提供-classpath/-cp选择，则使用当前目录作为用户类路径
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/**
 * 获取jre目录
 */
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) { // 优先使用用户输入的-Xjre选择作为jre目录
		return jreOption
	}
	if exists("./jre") { // 若没有输出，则在当前目录下寻找jre目录
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" { // 如果当前目录页找不到，尝试使用JAVA_HOME环境变量
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

/**
 * 判断目录是否存在
 */
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
} 