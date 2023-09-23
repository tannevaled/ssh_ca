package ConfigField

type Enum string

const (
	Keep Enum = "keep"
	Set  Enum = "set"
)

/*
 *
 * Type is only used in help text
 *
 */
func (self *Enum) Type() string {
	return "enum"
}
