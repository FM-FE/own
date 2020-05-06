#  VMware虚拟机创建教程

教程可以完成

1. 创建虚拟机
2. 可以在宿主机上通过ssh协议远程登录
3. 可以利用宿主机网络连接外网

## 创建虚拟机

1. 管理员权限打开VMware

2. 新建虚拟机

3. 自定义

4. 硬件兼容性

   ![image-20200122231903741](C:\Users\Friday\AppData\Roaming\Typora\typora-user-images\image-20200122231903741.png)

   如果是新建虚拟机，选择workstation 15.x，下一步即可

   > 如果想要拉起一个别处拷贝来的虚拟机（后缀为.vmx的虚拟机文件）
   >
   > 可能需要选择创建该虚拟机的VMware版本，再继续

5. 选择镜像

6. 选择虚拟机安装位置

7. 配置CPU，内存

8. 配置网络类型

   可以看到有三种网络类型，大概解释如下

   [详细解释在这里](https://blog.csdn.net/Noob_f/article/details/51099040)

   1. 桥接网络（bridge）

      桥接模式最好理解

      桥接模式下，所有的虚拟机网卡和宿主机网卡都连接在同一个虚拟交换机上

      因此，所有虚拟机和宿主机都要在同一个网段

      ---

      **如果需要联网**

      则要保证虚拟机的网关（GATEWAY），DNS，子网掩码都和宿主机保持一致

      所以如果宿主机的网络状态（IP等信息）发生变化，虚拟机的配置也需要随之改变，才能保证虚拟机和宿主机之间依然通畅，虚拟机依然可以访问外网

      

   2. 仅使用主机模式网络（host-only）

      主机模式和桥接模式类似

      差异在于

      主机模式下，虚拟机网卡都接在宿主机的虚拟网卡（VMnet1）上

      ---

      **如果需要联网**

      同上

      

   3. 网络地址转换（NAT）

      NAT模式在桥接模式的基础上，添加了虚拟NAT设备和虚拟DHCP服务器

      > NAT设备允许整个集群以一个公用IP地址出现在外网上
      >
      > 可以缓解公用IP地址不足的问题
      >
      > 桥接模式中的虚拟机，耗费的IP**可能**就是整个网络的公用IP
      >
      > 在虚拟机数量较多的情况下，可能导致IP冲突，实体机没有IP等情况

      ---

      **如果需要联网**

      首先需要配置NAT设备和DHCP服务器

      配置成功后，设置虚拟机网络连接协议为DHCP

      自动分配IP，连接外网

      

      综上，选择NAT模式，该模式对外部网络配置的影响最小

      NAT模式下，需要配置的网络为VMnet8

      > 思考题
   >
      > 1. VMnet8的IP，网关，掩码是否可以随意设置，无需和宿主机保持一致
   >
      >    
      >
      > 2. 最后，是否能通过网络共享，用正在连接的WiFi或以太网（有线）将VMnet8共享到外网
      >
      >    不能，共享后，宿主机无法远程到虚拟机

9. SCSI控制器，虚拟磁盘类型

   按默认推荐

10. 创建新虚拟磁盘

    按照需求分配磁盘大小

    选择：将磁盘拆分成多个文件

11. 安装虚拟机操作系统

    以CentOS为例
    
    导入镜像，按传统步骤安装即可
    
    安装时，耐心一些，不要随便关闭VMware
    
    否则可能会出现**问题一**

## 配置虚拟机网络

​	先用DHCP分配一个段内的IP

​	在将该IP写为固定IP（BOOTPROTO=static）

​	网上资料很多，[还能参考这个](https://blog.csdn.net/Noob_f/article/details/51099040)

+ 改完记得要重启网络

  `systemctl restart network`

+ 还要在 `/etc/resolv.conf` 文件中，添加DNS

  ![image-20200123022137414](C:\Users\Friday\AppData\Roaming\Typora\typora-user-images\image-20200123022137414.png)















# 问题及解决方案

1. VMX进程提前退出

   第一次安装操作系统时，界面黑屏，没有进入CentOS的系统引导

   强制退出了VMware，并重启电脑

   再次开启虚拟机时，报错如下

   ![image-20200123014404837](C:\Users\Friday\AppData\Roaming\Typora\typora-user-images\image-20200123014404837.png)

   **已解决**
   
   禁止VMware开机自启动，再重启电脑
   
   或者
   
   以admin用户登录Windows - 控制面板 - 卸载程序 - 找到VMware，选中 - 选择卸载按钮旁的**更改**
   
   按引导选择修复VMware
   
   打开powershell ( 管理员 ) ，执行
   
   `net user administrator /active:yes`
   
2. 配置IP后，虚拟机ping外网和宿主机都能通，但是宿主机无法ssh到虚拟机

   **已做**

   关闭firewalld，seLinux

   关闭宿主机的域网络防火墙

   重启虚拟机

   **已解决**

   不能给网关配IP为192.168.106.1

   按默认的来
   
   ![image-20200123032312533](C:\Users\Friday\AppData\Roaming\Typora\typora-user-images\image-20200123032312533.png)
   
   虚拟机网关为192.168.106.2
   
   虚拟机的DHCP范围为192.168.106.2-254
   
   VMnet8的IP地址为192.168.106.1，且没有网关