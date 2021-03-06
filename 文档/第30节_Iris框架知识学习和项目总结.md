### 项目总结与回顾
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

在前面29节的课程内容中，我们通过整个系列的课程，学习了GoWeb开发中的Iris框架的使用，以及如何使用Iris框架来完成一个实际项目的后台服务器开发。

本系列课程已经基本开发完成和结束，本节内容我们来做一个整体的项目总结和回顾。

### 项目架构
* **前端：**在本项目实战中，前端我们使用Vue框架来进行页面功能实现。Vue框架的特点是适合将前端和后端进行代码分离。
* **后端：**本实战项目，我们采用Iris框架进行实现，使用Iris框架开发服务器后台功能。
* **数据库：**本系列课程中，我们使用的数据库**mysql**和**redis**。关系型数据库使用mysql对表数据进行持久化；对于项目中的一些变动频率不高的数据，选择非关系型数据库**redis：**来进行存储。
* ***xorm：**在本项目实战中，数据库对象映射框架我们采用**xorm**进行实现，**xorm**框架是Goweb项目开发过程中非常常用的数据库对象关系映射框架，是实现后台数据库开发的重点。



### 项目功能开发列表
* 管理员模块
    * 管理员登录
    * 管理员退出
    * 管理员信息
    * 更新管理头像
    * 管理员统计
    
* 用户模块
    * 用户总数
    * 用户列表
    * 用户统计
    
* 商家模块
    * 商铺总数
    * 商铺列表
    * 添加商铺 

* 订单模块
    * 订单列表
    * 订单总数

* 食品模块
    * 食品列表
    * 食品总数
    * 添加食品类别记录
    * 食品类别列表

* 文件操作
    * 静态资源文件
    * 文件上传
    * 文件下载
    * 文件工具

### xorm框架知识点
* 数据库连接
    * mysql驱动
    * mysql配置及连接
    * redis安装及配置
    * redis操作及连接
    * redis数据类型
    
* 实体映射
    * 结构体映射规则
    * 字段映射规则

* 数据库操作
    * 查找
        * get、find
        * count
        * ID
        * where
    * 数据操作
        * 添加数据：insert
        * 删除数据：delete
        * 修改数据：update
    * 事务操作：
        * session.begin
        * 事务回滚
        
    