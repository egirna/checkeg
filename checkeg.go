package checkeg

import (
	"net/http"
	"time"
)

//CheckURL is
func CheckURL(url string) {

	for {
		resp, err := http.Get(url)
		if err != nil {
			print(err.Error())
		} else {

			if resp.StatusCode == 200 || resp.StatusCode == 301 || resp.StatusCode == 302 {
				println(string(resp.StatusCode) + resp.Status)
				//println(resp.StatusCode)
				println("okkk")

			} else {
				//println(string(resp.StatusCode) + resp.Status)
				println("Down")

			}

		}
		t1 := time.NewTimer(5 * time.Second)
		<-t1.C
	}

}