# Golang 的常见问题

### 1. new 和 make 有什么区别？

> **new:** 它是一个分配内存的内置函数，但与其他一些语言中的同名函数不同，它不会初始化内存，它只会将其归零。也就是说， new(T)为 type 的新项目分配零存储 T并返回其地址，即 type 的值*T

```go
// new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针
// 同时 new 函数会把分配的内存置为零，也就是类型的零值
func new(Type) *Type

func main() {
	var number *int
	number = new(int)                    //分配空间
	fmt.Printf("new type: %T\n", number) // new type: *[]int
	*number = 666
	fmt.Printf("number = %d\n", *number)
    /* 
    如果注释掉new(int)，则会报错
        panic: runtime error: invalid memory address or nil pointer dereference
        [signal 0xc0000005 code=0x1 addr=0x0 pc=0xd8c9e7] 
    */
}
```


> **make:** 内置函数make(T, args)的用途不同于new(T). 它只chan、map 以及 slice 的内存创建，它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针。

```go
//  make 函数的 t 参数必须是 chan（通道）、map（字典）、slice（切片）中的一个，并且返回值也是类型本身
func make(t Type, size ...IntegerType) Type

// 返回的值是类型本身
func main() {
	c := make(chan int)
	fmt.Printf("type: %T", c)   // chan int

	c1 := make(chan<- int, 1)
	fmt.Printf("type: %T", c1)  // chan <- int
}

```
**new 和 make 主要区别**
```text
1) make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
2) new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
3) new 分配的空间被清零。make 分配空间后，会进行初始化
```