// 定义ClassFile结构体
package classfile
import "fmt"

type classFile struct {
	// magic		uint32
	minorVersion	uint16 			// 次要版本信息
	majorVersion	uint16 			// 主要版本信息
	constantPool	ConstantPool 	// 常量池
	accessFlags		uint16 			// 访问标志
	thisClass		uint16 			// 常量池索引，当前类名，class文件存储的类名类似完全限定名，但把点换成了斜线
	superClass		uint16   		// 常量池索引，超类名，Object.class是0
	interfaces		[]uint16 		// 接口索引表，存放的也是常量池索引，给出该类实现的所有接口的名字
	fields			[]*MemberInfo	// 域（字段）表，存储字段信息
	methods			[]*MemberInfo	// 方法表，存储方法信息，字段和方法的基本结构大致相同
	attributes		[]AttributesInfo// 属性信息
}

/**
 * 把[]byte解析成ClassFile结构体
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func(){
		if r := recover(); r != nil { //go没有异常处理机制，只有一个panic-recover机制
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{
		classData
	}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/**
 * 解析class文件
 */
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader) // 
	self.readAndCheckVersion(reader) // 
	self.constantPool = readConstantPool(reader) // 
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool) // 
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool) //
}

/**
 * 魔数检查
 * 很多文件格式都会规定满足该格式的文件必须以某几个固定字节开头，
 * 这几个字节主要起标识作用，叫作魔数（magic number）。
 * 
 * 例如：PDF文件以4字节“%PDF”（0x25、0x50、0x44、0x46）开头，
 * 		ZIP文件以2字节“PK”（0x50、0x4B）开头。
 * 		class文件的魔数是“0xCAFEBABE”
 *
 * Java虚拟机规范规定，如果加载的class文件不符合要求的格式，Java虚拟机实现就抛出java.lang.ClassFormatError异常。
 * 但是因为我们才刚刚开始编写虚拟机，还无法抛出异常，所以暂时先调用panic（）方法终止程序执行。
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic :=reader.readUint32()
	if magic != 0xCAFEBABE { // 
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
 * 检查版本号 主版本M，次版本m ，一般是M.m格式，
 * 如果版本号不在支持的范围内，Java虚拟机实现就抛出java.lang.UnsupportedClassVersionError异常。
 * Java SE 8 支持版本号为45.0~52.0
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

/**
 * getter
 */
func (self *ClassFile) MinorVersion() uint16 { 
	return self.minorVersion
}

/**
 * getter
 */
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

/**
 * getter
 */
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
} 

/**
 * getter
 */
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
} 

/**
 * getter
 */
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

/**
 * getter
 */
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
} 

/**
 * 从常量池查找类名
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/**
 * 从常量池查找超类名
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""// 只有java.lang.Object没有超类
}

/**
 * 从常量池查找接口名
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
