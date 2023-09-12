package Config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	utils "ssh-ca/utils"
)

func New(
	ssh_ca_config_file_path *string,
) Interface {
	if utils.FileExists(*ssh_ca_config_file_path) {
		log.Println("Config database found.")
	} else {
		log.Println("Config database does not exists. Creating")
		os.Create(*ssh_ca_config_file_path)
	}

	db, err := sql.Open(
		"sqlite",
		*ssh_ca_config_file_path,
	)
	if err != nil {
		fmt.Println("failed to connect to database")
		os.Exit(1)
	}

	if _, err := db.Exec(`
CREATE TABLE IF NOT EXISTS certificate_authorities (
id            INTEGER PRIMARY KEY AUTOINCREMENT,
enabled       BOOL NOT NULL,
name          VARCHAR(32) NOT NULL UNIQUE,
description   VARCHAR(256),
private_key   VARCHAR(256) NOT NULL,
public_key    VARCHAR(256) NOT NULL,
creation_date TEXT
);
	`); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	__self := &Struct{
		config_file_path: ssh_ca_config_file_path,
		db:               db,
	}
	return __self
}
