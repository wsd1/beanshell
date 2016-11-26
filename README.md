#Beanshell
----
This a micro-service node for executing shell cmd writing in Golang.

20161126


## 相关文档

http://www.fzb.me/2015-3-21-beanstalkd-protocol-chinese-version.html

https://golang.org/pkg/os/exec/#Cmd.Run

## 下载运行beanstalkd
参见：https://github.com/kr/beanstalkd


    git clone https://github.com/kr/beanstalkd.git
    cd beanstalkd
    make
    ./beanstalkd -VVV


## 下载运行调试工具beanstool

参见：https://github.com/src-d/beanstool

    git clone https://github.com/src-d/beanstool.git
    cd beanstool

可能缺一堆东西：

    go get github.com/agtorre/gocolorize
    go get github.com/jessevdk/go-flags
    go get github.com/src-d/beanstool/cli
    go get golang.org/x/crypto/ssh/terminal

    go build beanstool.go 

然后就可以：

    ./beanstool stats

往default tube放置一个串：
    ./beanstool put -t default -b chashkd

## 测试golang代码

	go get xxx
	go build beanshell.go

Put job via beanstool:

	./beanstool put -t build-req -b "go build beanshell.go"

check result:
	
	./beanstool peek --tube build-ack --state ready

