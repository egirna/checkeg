import (
	"fmt"

	"github.com/egirna/checkeg"
)

func main() {
	
	fmt.Println(checkeg.CheckUrL("https://golangbot.com/go-packages/", []int{200, 300}))
	//fmt.Println(checkeg.CheckUrL("https://www.yahoo.com/games", []int{400, 500}))
	
}
