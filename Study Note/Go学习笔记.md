## Go学习笔记

1. 将string差分成单个字符（rune；对应char类型）输出

   [参考](https://stackoverflow.com/questions/18556693/slice-string-into-letters)

   直接将string赋值给 slice of rune

   `str := "abc"`

   `r := []rune(s)`

2. struct的成员只有大写才能被外部文件引用

   

3. 是不是类型定义后面，不跟 json，json.marshal 就不能解析了

   

4. struct 内部的变量有没有必要首字母大写

   如果 struct 类型的首字母不是大写，那内部变量首字母大写，还有意义吗？

   

5. “ ... ” 意味着，可变长参数

   e.g. 

   声明 ：`func (s *Server) InitRouter(routers ...router.Router)`

   调用：`server.InitRouter(routers...)`

6. go run 报错， undefined：[ 方法名 ]

   [reference](<https://blog.csdn.net/pingD/article/details/79143235>)

   同一文件夹下多个go文件

   如果只 go run main.go，会报错 undefined

   因为，此次编译只针对 main.go，导致其他 go 文件的方法，无法找到

   **解决方法：**

   `go run *.go`

   或者

   `go run main.go [go files you wanna compile]`

7. ID  string   \`bson:"_id,omitempty"`

   [参考链接](<https://zhidao.baidu.com/question/459318125920057725.html>)

   其中，omitempty表示，如果 ID 没有赋值，则 json 解析后不输出

   

8. **& 和 * 什么区别，怎么使用**

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