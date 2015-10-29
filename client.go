package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func SendWRQ(filename string) []byte{
	packet := []byte{0,2}
	packet = append(packet, []byte(filename)...)
	packet = append(packet, []byte{0}...)
	packet = append(packet, []byte("octet0")...)
	return packet
}
 
func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()
    i := 0
    for {
        msg := strconv.Itoa(i)
        i++
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
        time.Sleep(time.Second * 1)
    }
}
