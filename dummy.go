package dummy

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	a := aws.NewConfig()
	fmt.Print(a)
}
