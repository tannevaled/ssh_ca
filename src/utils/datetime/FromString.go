package DateTime

import (
	"fmt"
	"os"
	"time"

	"github.com/ncruces/go-strftime"
)

func FromString(datetimeString string) time.Time {
	datetime, err := strftime.Parse(Layout(), datetimeString)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	} // else {
	//	fmt.Printf("%q", datetime)
	//}
	return datetime
}
