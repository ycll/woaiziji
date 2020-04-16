# http://52ziji.top

# Linux 下安装 go

## 1.设置go环境变量
* 在$HOME/.profile下设置
```go
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/Applications/Go
```
* 编辑完成后保存，source文件内容
## 2.安装C工具
* centos7 下使用yum install  bison ed gawk gcc libc6-dev make
## 3.获取go源码
* [下载地址](https://golang.org/dl/)
```
 wget https://storage.googleapis.com/golang/go<VERSION>.src.tar.gz
 tar -zxvf go<VERSION>.src.tar.gz
 sudo mv go $GOROOT
```
## 4.构建go
```go
 cd $GOROOT/src
 ./all.bash
```

# beego介绍
* conf：项目配置文件所在的目录，项目中有一些全局的配置都可以放在此目录下。默认的app.conf文件中默认指定了三个配置：

    1）appname = BeegoDemo2： 指定项目名称。

    2）httpport = 8080： 指定项目服务监听端口。

    3）runmode = dev： 指定执行模式。
 
 * controllers：该目录是存放控制器文件的目录，所谓控制器就是控制应用调用哪些业务逻辑，由controllers处理完http请求以后，并负责返回给前端调用者。
 
 * models：models层可以解释为实体层或者数据层，在models层中实现和用户和业务数据的处理，主要和数据库表相关的一些操作会在这一目录中实现，然后将执行后的结果数据返回给controller层。比如向数据库中插入新数据，删除数据库表数据，修改某一条数据，从数据库中查询业务数据等都是在models层实现。
 
 * routers：该层是路由层。所谓路由就是分发的意思，当前端浏览器进行一个http请求达到后台web项目时，必须要让程序能够根据浏览器的请求url进行不同的业务处理，从接收到前端请求到判断执行具体的业务逻辑的过程的工作，就由routers来实现。
 
 * static：在static目录下，存放的是web项目的静态资源文件，主要有：css文件，img，js，html这几类文件。html中会存放应用的静态页面文件。
   
 * views：views中存放的就是应用中存放html模版页面的目录。所谓模版，就是页面框架和布局是已经使用html写好了的，只需要在进行访问和展示时，将获取到的数据动态填充到页面中，能够提高渲染效率。因此，模版文件是非常常见的一种方式。





