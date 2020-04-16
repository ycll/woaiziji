# http://52ziji.top

# Linux 下安装 go

## 1.设置go环境变量
* 在$HOME/.profile下设置
```go
export GROOT=$HOME/go
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

