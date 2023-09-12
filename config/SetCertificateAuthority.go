package Config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	SqliteResultCode "ssh-ca/sqlite-result-code"

	sqlite "github.com/glebarez/go-sqlite"
	strftime "github.com/itchyny/timefmt-go"
)

func (self Struct) SetCertificateAuthority(
	name string,
	description string,
	private_key string,
	public_key string,
) {
	res, err := self.db.Exec(`
INSERT INTO certificate_authorities(enabled,name,description,private_key,public_key,creation_date)
VALUES (true,?,?,?,?,?)`,
		name,
		description,
		private_key,
		public_key,
		strftime.Format(time.Now(), "%Y-%m-%d %H:%M:%S"),
	)

	if err != nil {
		var sqliteErr *sqlite.Error
		errors.As(err, &sqliteErr)
		//log.Printf("sqliteErr: %#v\n\n", sqliteErr)

		//r2 := err.(sqlite.Error)
		log.Printf("err: %#v\n\n", err)
		//	log.Printf("sqliteErr: %#v\n\n", sqliteErr)
		//		log.Printf("sqliteErr.code: %#v\n\n", err.(*sqlite.Error).Code())
		//	log.Printf("sqliteErr.code: %#v\n\n", sqliteErr.Code())
		//log.Printf("sqlite.ErrorCodeStrings: %#v\n\n", sqlite.ErrorCodeString)
		//log.Printf("sqliteErr.errorCodeString: %#v\n\n", sqlite.ErrorCodeString[sqliteErr.Code()])
		if int32(sqliteErr.Code()) == SqliteResultCode.SqliteResultCode.Extended().SQLITE_CONSTRAINT_UNIQUE() {
			log.Printf("ca name already exist\n")
			os.Exit(1)
		}
		//sqliteErr, _ := err.(sqlite3.Error)
		//log.Printf("sqliteErr: %#v\n\n", sqliteErr)
		//log.Printf(
		//	"ErrCodes\n\tsqlite3Err.ExtendedCode: %#v\n\tsqlite3.ErrConstraintUnique: %#v\n\tsqlite3.ErrConstraint: %#v\n\tsqlite3.ErrConstraintUnique: %#v\n\n",
		//	sqliteErr.Code,
		//	sqliteErr.ExtendedCode,
		//sqlite3.ErrConstraint,
		//sqlite3.ErrConstraintUnique,
		//)
		//if sqlite3Err, ok := err.(sqlite3.Error); ok {
		//	log.Printf("sqlite3Err.Code: %#v\n\n", sqlite3Err.Code)
		//	if sqlite3Err.Code == sqlite3.ErrConstraint && sqlite3Err.ExtendedCode == sqlite3.ErrConstraintUnique {
		//		log.Printf("ca name already exist\n")
		//		os.Exit(1)
		//	}
		//}
		log.Printf("err %#v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", res)
}
