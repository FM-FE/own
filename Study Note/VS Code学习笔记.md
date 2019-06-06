### VS Code学习笔记

1. VS Code将当前文件导入GOPATH

   设置 -> 工作区设置 -> 搜索 GOPATH

   点击工作区设置下的Go:Gopath -> 在setting.json中编辑

   输入如下代码

   {

   ​	 "go.gopath": "E:\\GitLab\\nas"

   }

   重启VS Code

   ---

2. 安装VS Code的Go插件

   由于 go get 命令不能使用
   
   1. VScode配置代理
   
      [参考链接](https://yeqown.github.io/2018/11/19/go-get遇到墙的解决方法/?utm_source=tuicool&utm_medium=referral)
   
      为VScode配置http代理
   
      http://localhost:41080
   
   2. 指定目录下，使用 git clone 替代 go get **（有效）**
   
      [参考链接](https://zhuanlan.zhihu.com/p/56567884)
   
      [gopm 添加到PATH环境变量中，便于 powershell 直接使用](<https://github.com/gpmgo/gopm>)
   
      golint 插件特殊，除去 `git clone tools` 之外，还需要 `git clone lint`
   
      [golint下载参考链接](https://juejin.im/post/5cb54db6f265da035632278a)
   
      
   
      依赖包下载完成后，重启 VScode，ctrl + shift + P，搜索并选择 go install/update tools
   
      全选所有插件，安装
   
      查看日志是否成功
   
   3. VScode代码没有自动补全
   
      在 setting.json 中添加
   
      `"go.autocompleteUnimportedPackages": true,`
   
      重启VScode，配置文件生效
   
      [配置文件参考](https://maiyang.me/post/2018-09-14-tips-vscode/)
   
   4. 快速打开/关闭调试控制台
   
      `ctrl + shift + Y`
      
   5. 打开终端
   
      `ctrl+`\` 
      
   6. 格式化代码
   
      `alt + shift + f`

