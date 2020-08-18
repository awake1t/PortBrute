/*
date:2020-08-15
author:awake1t
*/
package main

import (
	"PortBrute/brute"
	"PortBrute/common"
	"flag"
	"github.com/fatih/color"
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, `支持协议FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD.
`)
	flag.PrintDefaults()
}

func main() {

	h := flag.Bool("h",false,"帮助")
	ips := flag.String("f","ip.txt","要爆破的ip列表")
	thread := flag.Int("t",100,"扫描线程")
	user := flag.String("u","user.txt","用户名字典")
	pass := flag.String("p","pass.txt","密码字典")

	flag.Parse()

	if *h{
		usage()
		return
	}

	startTime := time.Now()

	userDict, uErr := common.ReadUserDict(*user)
	passDict, pErr := common.ReadUserDict(*pass)

	ipList := common.ReadIpList(*ips)

	color.Cyan("Thread: %d", *thread)
	color.Cyan("Number of ip list : %d", len(ipList))
	color.Cyan("Number of username dict : %d",len(userDict))
	color.Cyan("Number of password dict : %d", len(passDict))

	if uErr == nil && pErr == nil {
		scanTasks := brute.GenerateTask(ipList, userDict,passDict)
		color.Cyan("Number of all task : %d", len(scanTasks))
		brute.RunTask(scanTasks,*thread)
	} else {
		fmt.Println("Read File Err!")
	}

	endTime := time.Now()
	color.Red("Run Time is : %s\n", endTime.Sub(startTime))

}

