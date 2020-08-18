package plugins

import (
	"net"
)


func JavaDebug(ip string, port string, username string, password string) (err error, result bool) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Printf("Poc 扫描错误", err)
			return
		}
	}()

	conn, _ := net.Dial("tcp", ip+":"+port)
	//conn, err := net.Dial("tcp", "10.237.28.8:9091")

	//JDWP-Handshake
	conn.Write([]byte{0x4a, 0x44, 0x57, 0x50, 0x2d, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68,0x61, 0x6b, 0x65})
	defer conn.Close()
	buffer := make([]byte, 32)
	//读取返回数据
	res, _ := conn.Read(buffer)

	if res == 14{
		result = true
	}
	return err, result
}