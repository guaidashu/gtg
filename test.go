/**
  create by yy on 2019-10-21
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// 实际中应该用更好的变量名
var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func InitFlag() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
	Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]
	
	Options: 
`)
	flag.PrintDefaults()
}

func TestFlagString() {
	currentPath, _ := os.Getwd()
	// flag.Usage = Usage
	message := flag.String("s", "default s message", "it's user send message[help message]")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("no flag")
	} else {
		fmt.Println("args: ", args)
	}

	log.Println("message:", *message, "current path", currentPath)
}
