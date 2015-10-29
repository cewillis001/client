package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
		//"packet"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

/*
func SendWRQ(filename string, file string){
	//may at least return error
	//pass Conn into function
	_,err := Conn.Write(MakeWRQ(filename))
	buf := make([]byte, 1024)
	n, addr,err := Conn.ReadFromUDP(buf)
}
*/

func SendTest(i int, Conn *net.UDPConn){
	msg := strconv.Itoa(i)
	buf := []byte(msg)
	_,err := Conn.Write(buf)
	if err != nil{
		fmt.Println(msg, err)
	}
} 
 
func main() {
	//shortWriteTest := "Hello world"
	//shortWriteTestName := "hw"

	ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
	CheckError(err)
 
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)
 
	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)
 
	defer Conn.Close()
	i := 0
	for {
		SendTest(i, Conn)
		i++
		time.Sleep(time.Second * 1)
	}	
}
