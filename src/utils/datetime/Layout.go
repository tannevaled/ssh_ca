package DateTime

import (
	"fmt"
	"os"

	"github.com/ncruces/go-strftime"
)

func Layout() string {
	layout, err := strftime.Layout("%Y-%m-%d %H:%M:%S")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	} else {
		fmt.Printf("%q", layout)
	}
	return layout
}
