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

### 2. return 和 defer 的执行顺序？
> 多个defer出现的时候，它是一个“栈”的关系，也就是先进后出。一个函数中，写在前面的defer会比写在后面的defer调用的晚
> return之后的语句先执行，defer后的语句后执行

```go
func main() {
	run()
}

func run() string {
	defer deferFunc()
	return returnFunc()
}

func deferFunc() string {
	fmt.Println("defer...")
	return "defer"
}

func returnFunc() string {
	fmt.Println("return...")
	return "return"
}
/*
out: 
return...
defer...
*/
```

## 3. Golang 的内存泄露

> go中的内存泄露一般都是goroutine泄露，就是goroutine没有被关闭，或者没有添加超时控制，让goroutine一只处于阻塞状态，不能被GC。

#### 暂时性内存泄露

> 获取长字符串中的一段导致长字符串未释放
>
> 获取长slice中的一段导致长slice未释放
>
> 在长slice新建slice导致泄漏

#### 永久性内存泄露

> goroutine泄漏
>
> time.Ticker未关闭导致泄漏
>
> Finalizer导致泄漏
>
> Deferring Function Call导致泄漏

## 4. Golang 的变量逃逸分析

> "逃逸分析" 就是把变量合理地分配到它该去的地方，“找准自己的位置”。即使你是用new申请到的内存，如果我发现你竟然在退出函数后没有用了，那么就把你丢到栈上，毕竟栈上的内存分配比堆上快很多；反之，即使你表面上只是一个普通的变量，但是经过逃逸分析后发现在退出函数之后还有其他地方在引用，那我就把你分配到堆上，真正地做到“按需分配”。

> 如果变量都分配到堆上，堆不像栈可以自动清理。它会引起Go频繁地进行垃圾回收，而垃圾回收会占用比较大的系统开销（占用CPU容量的25%）。

**Go逃逸分析最基本的原则是：如果一个函数返回对一个变量的引用，那么它就会发生逃逸。**

编译器会根据变量是否被外部引用来决定是否逃逸：

> 1. 如果函数外部没有引用，则优先放到栈中；
> 2. 如果函数外部存在引用，则必定放到堆中；
