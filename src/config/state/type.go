package ConfigState

type Enum string

const (
	Keep    Enum = "keep"
	Enable  Enum = "enable"
	Disable Enum = "disable"
)

/*
 *
 * Type is only used in help text
 *
 */
func (self *Enum) Type() string {
	return "enum"
}

/*
type Interface interface {
	Set(string) error
	String() string
	Type() string
}
*/
