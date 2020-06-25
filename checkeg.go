package checkeg

import (
	"net"
	"net/http"
	"time"
	
)


func CheckUrL(url string, scode []int) (string, string) {

	for {
		resp, err := http.Get(url)
		if err != nil {
			print(err.Error())
		} else {
			for _, num := range scode {

				if num == resp.StatusCode {

					return "check is healthy", resp.Status

				} else {
					
					return "check is un healthy", resp.Status
				}
			}

		}
		/*t1 := time.NewTimer(5 * time.Second)
		<-t1.C*/
	}

}
//TCPcheck is
func TCPcheck(host, port string, timeout time.Duration) (string, string) {
	addr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return "False", err.Error()
		//println(err.Error())
	} else if conn != nil {
		//println(conn)
		return "True", " "
	} else {
		return " ", "nothing"
	}
}
