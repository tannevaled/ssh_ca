package Config

import (
	"log"
	ConfigField "ssh-ca/src/config/field"
	ConfigState "ssh-ca/src/config/state"
	"ssh-ca/src/utils"
)

func (self Struct) EditCertificateAuthority(
	uuid *string,
	state *ConfigState.Enum,
	description *ConfigField.Enum,
	description_value *string,

) {
	var sqlStatement string
	var enabled bool = false

	utils.Trace()
	//if description != nil {
	log.Printf("state: %#v", *state)
	log.Printf("description: %#v", *description)
	//}
	if *state != ConfigState.Keep {
		if *state == ConfigState.Enable {
			enabled = true
		}

		sqlStatement = `
UPDATE certificate_authorities
SET enabled=?
WHERE uuid=?`
		_, err := self.db.Exec(sqlStatement, enabled, uuid)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	}
	if *description != ConfigField.Keep {
		sqlStatement = `
UPDATE certificate_authorities
SET description=?
WHERE uuid=?`
		_, err := self.db.Exec(sqlStatement, description_value, uuid)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
	}

}
