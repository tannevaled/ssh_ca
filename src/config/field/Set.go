package ConfigField

import "errors"

func (self *Enum) Set(value string) error {
	switch value {
	case "set", "keep":
		*self = Enum(value)
		return nil
	default:
		return errors.New(`must be one of "set" or "keep"`)
	}
}
