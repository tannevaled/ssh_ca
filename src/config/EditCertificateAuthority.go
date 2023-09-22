package Config

import (
	"log"
	ConfigState "ssh-ca/src/config/state"
	"ssh-ca/src/utils"
)

func (self Struct) EditCertificateAuthority(
	uuid *string,
	description *string,
	state *ConfigState.Enum,
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

}
