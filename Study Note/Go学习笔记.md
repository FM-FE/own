## Go学习笔记

1. “ ... ” 意味着，可变长参数

   e.g. 

   声明 ：`func (s *Server) InitRouter(routers ...router.Router)`

   调用：`server.InitRouter(routers...)`

2. go run 报错， undefined：[ 方法名 ]

   [reference](<https://blog.csdn.net/pingD/article/details/79143235>)

   同一文件夹下多个go文件

   如果只 go run main.go，会报错 undefined

   因为，此次编译只针对 main.go，导致其他 go 文件的方法，无法找到

   **解决方法：**

   `go run *.go`

   或者

   `go run main.go [go files you wanna compile]`

3. 如何使用 var 定义的变量

   

4. **& 和 * 什么区别，怎么使用**

   & 代表取地址

   \* 代表取值

   ```golang
   func main() {
       b := 6 
   
       var b_ptr *int // *int is used delcare variable
                      // b_ptr to be a pointer to an int
   
       b_ptr = &b     // b_ptr is assigned value that is the address
                          // of where variable b is stored
   
       // Shorthand for the above two lines is:
       // b_ptr := &b
   
       fmt.Printf("address of b_ptr: %p\n", b_ptr)
   
       // We can use *b_ptr get the value that is stored
       // at address b_ptr, or dereference the pointer 
       fmt.Printf("value stored at b_ptr: %d\n", *b_ptr)
   
   }
   ```

   运行结果：

   ```
   address of b_ptr: 0xc82007c1f0
   value stored at b_ptr: 6
   ```