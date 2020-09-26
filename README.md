### 介绍

> 一款跨平台小巧的端口爆破工具，支持爆破FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD。 

### 使用说明

> ​	ip.txt中放入需要爆破的ip+端口，比如 10.10.10.10:3306。  如果不是标准端口，比如3307是MYSQL。写成 10.10.10.10:3306|MYSQL。 其他 FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD同理
>
> ​	Windows/MAC/Linux已经打包好了，分别对应(PortBruteWin.exe/PortBruteMac/PortBruteLinux)。无论是内网渗透，还是日常使用，直接下载下来就可以直接用。

![image](https://github.com/awake1t/PortBrute/blob/master/common/example1.png)


### 快速开始

​	先把要爆破的资产放入ip.txt中，直接运行下面命令开始爆破。 结果会在当前目录生成res.txt

`PortBruteWin.exe`


### 已完成
  - [x] 支持多协议 FTP/SSH/SMB/MSSQL/MYSQL/POSTGRESQL/MONGOD
  - [x] 支持非常规端口
  - [x] 支持user:pass 字典模式 [2020-0829]
  - [x] 爆破时看到进度，增加了跑马灯 [2020-0926]

### 待完成
  - [ ] 爆破时更完美的进度展示
  - [ ] 更快！更快！更快！


### 致谢

感谢`netxfly`师傅，学习了很多。感谢感谢～

https://github.com/netxfly/x-crack
