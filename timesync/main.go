package main

import (
	"encoding/binary"
	"flag"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/gogf/gf/os/gproc"
)

const ntpEpochOffset = 2208988800
const timeLayoutStr = "2006-01-02 15:04:05"

type packet struct {
	Settings       uint8
	Stratum        uint8
	Poll           int8
	Precision      int8
	RootDelay      uint32
	RootDispersion uint32
	ReferenceID    uint32
	RefTimeSec     uint32
	RefTimeFrac    uint32
	OrigTimeSec    uint32
	OrigTimeFrac   uint32
	RxTimeSec      uint32
	RxTimeFrac     uint32
	TxTimeSec      uint32
	TxTimeFrac     uint32
}

func main() {

	ntime := GetRemoteTime()
	ts := ntime.Format(timeLayoutStr)
	glog.Infof("接收到ntp服务器时间：%v", ts)
	log2file("接收到ntp服务器时间：" + ts)

	UpdateSystemDate(ts)
}

func GetRemoteTime() time.Time {

	var host string

	//flag.StringVar(&host, "e", "time.windows.com:123", "NTP host")

	// 182.92.12.11:123 是阿里的ntp服务器
	flag.StringVar(&host, "e", "182.92.12.11:123", "NTP host")
	flag.Parse()

	conn, err := net.Dial("udp", host)
	if err != nil {
		glog.Warningf("failed to connect: %v", err.Error())
		log2file("failed to connect: " + err.Error())
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
		glog.Warningf("failed to set deadline: %v", err.Error())
		log2file("failed to set deadline: " + err.Error())
	}

	req := &packet{Settings: 0x1B}

	if err := binary.Write(conn, binary.BigEndian, req); err != nil {
		glog.Warningf("failed to send request: %v", err.Error())
		log2file("failed to send request: " + err.Error())
	}

	rsp := &packet{}
	if err := binary.Read(conn, binary.BigEndian, rsp); err != nil {
		glog.Warningf("failed to read server response: %v", err.Error())
		log2file("failed to read server response: " + err.Error())
	}

	secs := float64(rsp.TxTimeSec) - ntpEpochOffset
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32

	return time.Unix(int64(secs), nanos)
}

func UpdateSystemDate(dateTime string) bool {

	system := runtime.GOOS
	switch system {
	case "windows":
		{
			_, err1 := gproc.ShellExec(`date  ` + gstr.Split(dateTime, " ")[0])
			_, err2 := gproc.ShellExec(`time  ` + gstr.Split(dateTime, " ")[1])
			if err1 != nil && err2 != nil {
				glog.Warningf("更新系统时间错误: 请用管理员身份启动程序!")
				log2file("更新系统时间错误: 请用管理员身份启动程序!")
				return false
			}
			return true
		}
	case "linux":
		{
			_, err1 := gproc.ShellExec(`busybox date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Warningf("更新系统时间错误: %v", err1.Error())
				log2file("更新系统时间错误: " + err1.Error())
				return false
			}
			return true
		}
	case "darwin":
		{
			_, err1 := gproc.ShellExec(`date -s  "` + dateTime + `"`)
			if err1 != nil {
				glog.Warningf("更新系统时间错误: %v", err1.Error())
				log2file("更新系统时间错误: " + err1.Error())
				return false
			}
			return true
		}
	}

	return false
}

func log2file(context string) {
	// appendFile("./timeSync.log", time.Now().Format("2006-01-02 15:04:05") + ": " + context)
}

func appendFile(path string, context string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		file.WriteString(context)
		file.WriteString("\r\n")
		file.Close()
	}
}
