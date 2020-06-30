# checkeg
Server Health check implementation in Golang (Go)
# Installing 
```console
    go get -u  github.com/egirna/checkeg                                 
```

# Usage
**Import the package**
```go
    import "github.com/egirna/checkeg"                                    
 ```


## URL Health Check and  TCP Dial Check

#example
``` go 
 package main
 import (
	"fmt"
	"time"
	
	"github.com/egirna/checkeg"
 )

 func main() {
	
    fmt.Println(checkeg.CheckUrL("https://golangbot.com/go-packages/", []int{200, 300}))
    fmt.Println(checkeg.TCPcheck("golang.org", "80", 6*time.Second))

```
