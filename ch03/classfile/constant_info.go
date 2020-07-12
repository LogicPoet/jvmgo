package classfile //区分常量池的常量类型，java虚拟机规定了14种类型
// tag常量值定义

const (
	CONSTANT_Class = 7
	CONSTANT_Fieldref = 9
	CONSTANT_Methodref = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String = 8
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_NameAndType = 12
	CONSTANT_Utf8 = 1
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)