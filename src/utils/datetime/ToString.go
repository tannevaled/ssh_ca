package DateTime

import (
	"time"

	"github.com/ncruces/go-strftime"
)

func ToString(datetime time.Time) string {
	return strftime.Format(Layout(), datetime)
}
