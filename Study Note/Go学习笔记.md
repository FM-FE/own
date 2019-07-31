## Go学习笔记

1. Mutex互斥锁

   [参考（概念演示）](https://shockerli.net/post/golang-pkg-mutex/)

   [参考（拓展）](https://cyent.github.io/golang/goroutine/sync_mutex/)

   在某个 go routine（a），mutex.lock() 之后，只要没有 unlock()

   其他 go routine（others）程序中，如果同样调用到 mutex.lock() 就会阻塞住

   直到 go routine（a）解锁 unlock()

   其他 go routine（others）才能再次加锁 lock()

2. Channel

   [参考](https://blog.csdn.net/sureSand/article/details/79633926)

   channel相当于一个锁

   只要channel中的数据，没有被其他go routine读取

   那么程序就会阻塞在channel那行

   直到channel中的数据被其他go routine读取

3. 通过mgo包，连接mongo数据库

   [参考（简单例子）](https://blog.csdn.net/oqqyuan1234567890/article/details/70186134)

   + 创建mongo连接，获取session
   + 调用 `session.DB.c` 创建数据库和表
   + 调用 `c.Insert` 向表中插入数据
   + 调用 `c.Find` 从表中查找数据

4. interface

   

5. mgolock的作用

   

6. go build -o < tag > 

   默认将当前文件夹下的 go 文件项目制作为二进制可执行文件

7. 删除 slice 中的元素

   `slice = append(slice[:i],slice[i+1:]...)`

8. add a slice to an other slice

   [参考](https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go)

   `append(slice1, slice2...)`

   add "..." to make it work

   **figure out what "..." does**

9. 将string差分成单个字符（rune；对应char类型）输出

   [参考](https://stackoverflow.com/questions/18556693/slice-string-into-letters)

   直接将string赋值给 slice of rune

   `str := "abc"`

   `r := []rune(s)`

10. struct的成员只有大写才能被外部文件引用

    

11. 是不是类型定义后面，不跟 json，json.marshal 就不能解析了

    

12. struct 内部的变量有没有必要首字母大写

    如果 struct 类型的首字母不是大写，那内部变量首字母大写，还有意义吗？

    

13. “ ... ” 意味着，可变长参数

    e.g. 

    声明 ：`func (s *Server) InitRouter(routers ...router.Router)`

    调用：`server.InitRouter(routers...)`

14. go run 报错， undefined：[ 方法名 ]

    [reference](<https://blog.csdn.net/pingD/article/details/79143235>)

    同一文件夹下多个go文件

    如果只 go run main.go，会报错 undefined

    因为，此次编译只针对 main.go，导致其他 go 文件的方法，无法找到

    **解决方法：**

    `go run *.go`

    或者

    `go run main.go [and go files name you wanna compile]`

15. ID  string   \`bson:"_id,omitempty"`

   [参考链接](<https://zhidao.baidu.com/question/459318125920057725.html>)

   其中，omitempty表示，如果 ID 没有赋值，则 json 解析后不输出

11. unsafe.Pointer() 的作用

    [参考](https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html)

    强制转化不同类型的指针

    ```go
    func main() {
    	i:= 10 // int类型
    	ip:=&i // int类型指针
        // 强制转化int类型指针为float64类型指针
    	var fp *float64 = (*float64)(unsafe.Pointer(ip))
    	*fp = *fp * 3
    	fmt.Println(i)
    }
    ```
    
12. **& 和 * 什么区别，怎么使用**

    & 代表取地址

    \* 代表取值

    ```go
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