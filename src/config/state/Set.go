package ConfigState

import "errors"

func (self *Enum) Set(value string) error {
	switch value {
	case "enable", "disable", "keep":
		*self = Enum(value)
		return nil
	default:
		return errors.New(`must be one of "enable", "disable" or "keep"`)
	}
}
