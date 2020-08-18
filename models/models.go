package models

type Service struct {
	Ip string
	Port int
	Protocol string
	UserName string
	PassWord string
}

type ScanResult struct {
	Service Service
	Result bool
}

type IpAddr struct {
	Ip string
	Port int
	Protocol string
}
