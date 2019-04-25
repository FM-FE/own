### docker 学习笔记

1. image 导出为 tar 文件

   docker save [image:tag] > [image.tar]

   docker save [image:tag] -o [image.tar]

2. 导入image.tar

   docker load < [image.tar]

   docker load -i [image.tar]