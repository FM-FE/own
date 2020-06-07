### Hexo学习笔记

1. 新建文章（post）

   + 在 博客目录下打开git bash ，输入命令

     `hexo new post "article title"`

   + 打开博客根目录 blog\source\\_posts
   + 编辑已经新建好的 article title.md 
   + `hexo clean` 清空缓存
   + `hexo g`  让修改生效
   + `hexo s`  本地预览
   + `hexo d`  部署到远端 gitpage 

2. 删除文章

   删掉source/_posts目录下的markdown

   > 注意：必须保证目录下至少有一篇文章，否则会报错 cannot find /
   
   
   
3. 设置文章封面

   

4. 让博客带有图片

   将图片放到`source/images`文件夹中，通过相对路径来访问

   [博客图片解决方案]()
   
   1. typora - 文件 - 偏好设置 - 图像
   
      ![image-20200603224747763](Hexo学习笔记/image-20200603224747763.png)
   
   2. hexo根目录 - 编辑`_config.yml` -  设置`post_asset_folder`为`true`
   
   3. 插件准备 - `npm install hexo-image-link --save`
   
   4. `hexo g` - `hexo s` - `hexo d`
   
      

