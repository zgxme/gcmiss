- [gcmiss(工程小秘书)](#gcmiss-------)
  * [项目介绍](#----)
    + [FE](#fe)
    + [BACKEND](#backend)
  * [部署](#--)
    + [环境依赖](#----)
    + [配置修改](#----)
    + [线上部署](#----)
      - [编译运行](#----)
      - [Nginx配置](#nginx--)
  * [The End](#the-end)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

# gcmiss(工程小秘书)  
## 项目介绍  
毕业设计校园生活服务网站，集校园交流、失物招领、寻物启事、求人办事、二手市场等功能  
项目前后端分离
用户注册使用邮箱激活方式，项目图片存储使用阿里oss静态资源管理，同时支持ip封禁

### FE  
前端使用Vue框架Vuetify组件    
项目链接  
https://github.com/ctrlcer/gcmiss_fe  

### BACKEND  
后端使用beego框架  
项目链接  
https://github.com/ctrlcer/gcmiss  

## 部署  
### 环境依赖  
安装，配置MySQL  
用户名root、密码qtest、端口3306、host为127.0.0.1  
安装、配置Redis  
Redis的host为127.0.0.1、端口6379  

### 配置修改  
gcmiss/conf/app.conf  
ak、sk为用户阿里OSS静态资源bucket的ak、sk  
password、username为开通smtp服务的邮箱的授权码和邮箱账号  

### 线上部署  

#### 编译运行  
前端项目 gcmiss-fe  
编译产出  
npm run build  
将产出移到views目录下

后端项目  
nohup go run main.go  >/dev/null 2>&1 &  

#### Nginx配置  
恶意请求ip封禁，可参考    
https://blog.csdn.net/weixin_33946020/article/details/91834554?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase  
博客脚本有个错误  
```lua
-- res , err = cache:expire("bind_",ip_bind_time) 应该改为
res , err = cache:expire("bind_"..ngx.var.remote_addr,ip_bind_time)
```
## The End
如果你对项目感兴趣的话，可以提patch，期待你的star⭐  
Thanks all  
Email: purifiedzheng@gmail.com  
Blog：
