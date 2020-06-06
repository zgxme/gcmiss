[TOC]  

# gcmiss(工程小秘书)  
## 项目介绍  
校园生活服务网站，集校园交流、失物招领、寻物启事、求人办事、二手市场等功能  
项目前后端分离  

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
nohup go run main.go  >/dev/null 2>&1 &  

#### Nginx配置  
恶意请求ip封禁，可参考    
https://blog.csdn.net/weixin_33946020/article/details/91834554?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase

