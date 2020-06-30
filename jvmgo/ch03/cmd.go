// 定义一个结构体来表示命令行选项和参数
// 在go语言中，main是一个特殊的包，这个包所在的目录（可以叫作任何名字）会被编译为可执行文件
// go程序的入口也是main()函数，但是不接收任何参数，也不能有返回值
package main 
import "flag"
import "fmt"
import "os"
type Cmd struct {
	helpFlag	bool
	versionFlag	bool
	cpOption	string
	XjreOption	string // 非标准选项-Xjre,指定jre目录的位置来寻找和加载Java标准库中的类。
	class	string
	args []string
}

/**
 * Go语言有函数（Function）和方法（Method）之分，方法调用需要
 * receiver，函数调用则不需要。
 */
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp" ,"", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre" ,"", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}