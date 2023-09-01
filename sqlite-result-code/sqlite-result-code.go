/*
 *
 * as of 2023/08/07 github.com/glebarez/go-sqlite does not define sqlite result codes, so we do
 *
 */

package SqliteResultCode

var SqliteResultCode SqliteResultCodeInterface = SqliteResultCodeStruct{
	primary:  SqlitePrimaryResultCodeStruct{},
	extended: SqliteExtendedResultCodeStruct{},
}

/*
 * https://www.sqlite.org/rescode.html
 *
 * Result codes are signed 32-bit integers.
 * The least significant 8 bits of the result code define a broad category
 * and are called the "primary result code".
 * More significant bits provide more detailed information about the error
 * and are called the "extended result code".
 * Note that the primary result code is always a part of the extended result code.
 * Given a full 32-bit extended result code, the application can always find
 * the corresponding primary result code merely by extracting the least
 * significant 8 bits of the extended result code.
 * All extended result codes are also error codes. Hence the terms
 * "extended result code" and "extended error code" are interchangeable.
 *
 */

type SqliteResultCodeStruct struct {
	primary  SqlitePrimaryResultCodeStruct
	extended SqliteExtendedResultCodeStruct
}
type SqliteResultCodeInterface interface {
	Primary() SqlitePrimaryResultCodeStruct
	Extended() SqliteExtendedResultCodeStruct
}

func (self SqliteResultCodeStruct) Primary() SqlitePrimaryResultCodeStruct {
	return self.primary
}
func (self SqliteResultCodeStruct) Extended() SqliteExtendedResultCodeStruct {
	return self.extended
}

type SqlitePrimaryResultCodeStruct struct{}
type SqlitePrimaryResultCodeInterface interface {
	SQLITE_ABORT() int32
	SQLITE_AUTH() int32
	SQLITE_BUSY() int32
	SQLITE_CANTOPEN() int32
	SQLITE_CONSTRAINT() int32
	//...
	SQLITE_CORRUPT() int32
	SQLITE_DONE() int32
	SQLITE_EMPTY() int32
	SQLITE_ERROR() int32
	SQLITE_FORMAT() int32
	//...
	SQLITE_FULL() int32
	SQLITE_INTERNAL() int32
	SQLITE_INTERRUPT() int32
	SQLITE_IOERR() int32
	SQLITE_LOCKED() int32
	//...
	SQLITE_MISMATCH() int32
	SQLITE_MISUSE() int32
	SQLITE_NOLFS() int32
	SQLITE_NOMEM() int32
	SQLITE_NOTADB() int32
	// ...
	SQLITE_NOTFOUND() int32
	SQLITE_NOTICE() int32
	SQLITE_OK() int32
	SQLITE_PERM() int32
	SQLITE_PROTOCOL() int32
	// ...
	SQLITE_RANGE() int32
	SQLITE_READONLY() int32
	SQLITE_ROW() int32
	SQLITE_SCHEMA() int32
	SQLITE_TOOBIG() int32
	// ...
	SQLITE_WARNING() int32
}

func (SqlitePrimaryResultCodeStruct) SQLITE_ABORT() int32      { return 4 }  // 0b 00000000 00000000 00000000 00000100
func (SqlitePrimaryResultCodeStruct) SQLITE_AUTH() int32       { return 23 } // 0b 00000000 00000000 00000000 00010111
func (SqlitePrimaryResultCodeStruct) SQLITE_BUSY() int32       { return 5 }  // 0b 00000000 00000000 00000000 00000101
func (SqlitePrimaryResultCodeStruct) SQLITE_CANTOPEN() int32   { return 14 } // 0b 00000000 00000000 00000000 00001110
func (SqlitePrimaryResultCodeStruct) SQLITE_CONSTRAINT() int32 { return 19 } // 0b 00000000 00000000 00000000 00010011
// ...
func (SqlitePrimaryResultCodeStruct) SQLITE_CORRUPT() int32 { return 11 }  //
func (SqlitePrimaryResultCodeStruct) SQLITE_DONE() int32    { return 101 } //
func (SqlitePrimaryResultCodeStruct) SQLITE_EMPTY() int32   { return 16 }  //
func (SqlitePrimaryResultCodeStruct) SQLITE_ERROR() int32   { return 1 }   //
func (SqlitePrimaryResultCodeStruct) SQLITE_FORMAT() int32  { return 24 }  //
// ...
func (SqlitePrimaryResultCodeStruct) SQLITE_FULL() int32      { return 13 } //
func (SqlitePrimaryResultCodeStruct) SQLITE_INTERNAL() int32  { return 2 }  // 0b 00000000 00000000 00000000 00000010
func (SqlitePrimaryResultCodeStruct) SQLITE_INTERRUPT() int32 { return 9 }  // 0b 00000000 00000000 00000000 00001001
func (SqlitePrimaryResultCodeStruct) SQLITE_IOERR() int32     { return 10 } // 0b 00000000 00000000 00000000 00001010
func (SqlitePrimaryResultCodeStruct) SQLITE_LOCKED() int32    { return 6 }  // 0b 00000000 00000000 00000000 00000110
// ...
func (SqlitePrimaryResultCodeStruct) SQLITE_MISMATCH() int32 { return 20 } // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_MISUSE() int32   { return 21 } // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_NOLFS() int32    { return 22 } // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_NOMEM() int32    { return 7 }  // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_NOTADB() int32   { return 26 } // 0b
// ...
func (SqlitePrimaryResultCodeStruct) SQLITE_NOTFOUND() int32 { return 12 } // 0b 00000000 00000000 00000000 00001100
func (SqlitePrimaryResultCodeStruct) SQLITE_NOTICE() int32   { return 27 } // 0b 00000000 00000000 00000000
func (SqlitePrimaryResultCodeStruct) SQLITE_OK() int32       { return 0 }  // 0b 00000000 00000000 00000000 00000000
func (SqlitePrimaryResultCodeStruct) SQLITE_PERM() int32     { return 3 }  // 0b 00000000 00000000 00000000 00000011
func (SqlitePrimaryResultCodeStruct) SQLITE_PROTOCOL() int32 { return 15 } // 0b 00000000 00000000 00000000 00001111
// ...
func (SqlitePrimaryResultCodeStruct) SQLITE_RANGE() int32    { return 25 }  // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_READONLY() int32 { return 8 }   // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_ROW() int32      { return 100 } // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_SCHEMA() int32   { return 17 }  // 0b
func (SqlitePrimaryResultCodeStruct) SQLITE_TOOBIG() int32   { return 28 }  // 0b
// ...
func (SqlitePrimaryResultCodeStruct) SQLITE_WARNING() int32 { return 28 } // 0b

type SqliteExtendedResultCodeStruct struct{}
type SqliteExtendedResultCodeInterface interface {
	SQLITE_ABORT_ROLLBACK() int32
	SQLITE_AUTH_USER() int32
	SQLITE_BUSY_RECOVERY() int32
	SQLITE_BUSY_SNAPSHOT() int32
	SQLITE_BUSY_TIMEOUT() int32
	SQLITE_CANTOPEN_CONVPATH() int32
	SQLITE_CANTOPEN_DIRTYWAL() int32
	SQLITE_CANTOPEN_FULLPATH() int32
	SQLITE_CANTOPEN_ISDIR() int32
	SQLITE_CANTOPEN_NOTEMPDIR() int32
	SQLITE_CANTOPEN_SYMLINK() int32
	SQLITE_CONSTRAINT_CHECK() int32
	SQLITE_CONSTRAINT_COMMITHOOK() int32
	SQLITE_CONSTRAINT_DATATYPE() int32
	SQLITE_CONSTRAINT_FOREIGNKEY() int32
	// ...
	SQLITE_CONSTRAINT_FUNCTION() int32
	SQLITE_CONSTRAINT_NOTNULL() int32
	SQLITE_CONSTRAINT_PINNED() int32
	SQLITE_CONSTRAINT_PRIMAREYKEY() int32
	SQLITE_CONSTRAINT_ROWID() int32
	SQLITE_CONSTRAINT_TRIGGER() int32
	SQLITE_CONSTRAINT_UNIQUE() int32
	SQLITE_CONSTRAINT_VTAB() int32
	SQLITE_CORRUPT_INDEX() int32
	SQLITE_CORRUPT_SEQUENCE() int32
	SQLITE_CORRUPT_VTAB() int32
	SQLITE_ERROR_MISSING_COLLSEQ() int32
	SQLITE_ERROR_RETRY() int32
	SQLITE_ERROR_SNAPSHOT() int32
	SQLITE_IOERR_ACCESS() int32
	// ...
	SQLITE_IOERR_AUTH() int32
	SQLITE_IOERR_BEGIN_ATOMIC() int32
	SQLITE_IOERR_BLOCKED() int32
	SQLITE_IOERR_CHECKRESERVEDLOCK() int32
	SQLITE_IOERR_CLOSE() int32
	SQLITE_IOERR_COMMIT_ATOMIC() int32
	SQLITE_IOERR_CONVPATH() int32
	SQLITE_IOERR_CORRUPTFS() int32
	SQLITE_IOERR_DATA() int32
	SQLITE_IOERR_DELETE() int32
	SQLITE_IOERR_DELETE_NOENT() int32
	SQLITE_IOERR_DIR_CLOSE() int32
	SQLITE_IOERR_DIR_FSYNC() int32
	SQLITE_IOERR_FSTAT() int32
	SQLITE_IOERR_FSYNC() int32
	// ...
	SQLITE_IOERR_GETTEMPPATH() int32
	SQLITE_IOERR_LOCK() int32
	SQLITE_IOERR_MMAP() int32
	SQLITE_IOERR_NOMEM() int32
	SQLITE_IOERR_RDLOCK() int32
	SQLITE_IOERR_READ() int32
	SQLITE_IOERR_ROLLBACK_ATOMIC()
	SQLITE_IOERR_SEEK() int32
	SQLITE_IOERR_SHMLOCK() int32
	SQLITE_IOERR_SHMMAP() int32
	SQLITE_IOERR_SHMOPEN() int32
	SQLITE_IOERR_SHMSIZE() int32
	SQLITE_IOERR_SHORT_READ() int32
	SQLITE_IOERR_TRUNCATE() int32
	SQLITE_IOERR_UNLOCK() int32

	// ...
	SQLITE_IOERR_VNODE() int32
	SQLITE_IOERR_WRITE() int32
	SQLITE_LOCKED_SHAREDCACHE() int32
	SQLITE_LOCKED_VTAB() int32
	SQLITE_NOTICE_RECOVER_ROLLBACK() int32
	SQLITE_NOTICE_RECOVER_WAL() int32
	SQLITE_OK_LOAD_PERMANENTLY() int32
	SQLITE_READONLY_CANTINIT() int32
	SQLITE_READONLY_CANTLOCK() int32
	SQLITE_READONLY_DBMOVED() int32
	SQLITE_READONLY_DIRECTORY() int32
	SQLITE_READONLY_RECOVERY() int32
	SQLITE_READONLY_ROLLBACK() int32
	SQLITE_WARNING_AUTOINDEX() int32
}

func (SqliteExtendedResultCodeStruct) SQLITE_ABORT_ROLLBACK() int32        { return 516 }
func (SqliteExtendedResultCodeStruct) SQLITE_AUTH_USER() int32             { return 279 }
func (SqliteExtendedResultCodeStruct) SQLITE_BUSY_RECOVERY() int32         { return 261 }
func (SqliteExtendedResultCodeStruct) SQLITE_BUSY_SNAPSHOT() int32         { return 517 }
func (SqliteExtendedResultCodeStruct) SQLITE_BUSY_TIMEOUT() int32          { return 773 }
func (SqliteExtendedResultCodeStruct) SQLITE_CANTOPEN_CONVPATH() int32     { return 1038 }
func (SqliteExtendedResultCodeStruct) SQLITE_CANTOPEN_DIRTYWAL() int32     { return 1294 }
func (SqliteExtendedResultCodeStruct) SQLITE_CANTOPEN_FULLPATH() int32     { return 782 }
func (SqliteExtendedResultCodeStruct) SQLITE_CANTOPEN_ISDIR() int32        { return 526 }
func (SqliteExtendedResultCodeStruct) SQLITE_CANTOPEN_NOTEMPDIR() int32    { return 270 }
func (SqliteExtendedResultCodeStruct) SQLITE_CANTOPEN_SYMLINK() int32      { return 1550 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_CHECK() int32      { return 275 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_COMMITHOOK() int32 { return 531 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_DATATYPE() int32   { return 3091 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_FOREIGNKEY() int32 { return 787 }

// ...
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_FUNCTION() int32    { return 1043 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_NOTNULL() int32     { return 1299 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_PINNED() int32      { return 2835 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_PRIMAREYKEY() int32 { return 1555 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_ROWID() int32       { return 2579 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_TRIGGER() int32     { return 1811 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_UNIQUE() int32      { return 2067 }
func (SqliteExtendedResultCodeStruct) SQLITE_CONSTRAINT_VTAB() int32        { return 2323 }
func (SqliteExtendedResultCodeStruct) SQLITE_CORRUPT_INDEX() int32          { return 779 }
func (SqliteExtendedResultCodeStruct) SQLITE_CORRUPT_SEQUENCE() int32       { return 523 }
func (SqliteExtendedResultCodeStruct) SQLITE_CORRUPT_VTAB() int32           { return 267 }
func (SqliteExtendedResultCodeStruct) SQLITE_ERROR_MISSING_COLLSEQ() int32  { return 257 }
func (SqliteExtendedResultCodeStruct) SQLITE_ERROR_RETRY() int32            { return 513 }
func (SqliteExtendedResultCodeStruct) SQLITE_ERROR_SNAPSHOT() int32         { return 769 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_ACCESS() int32           { return 3338 }

// ...
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_AUTH() int32              { return 7128 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_BEGIN_ATOMIC() int32      { return 7434 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_BLOCKED() int32           { return 2826 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_CHECKRESERVEDLOCK() int32 { return 3594 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_CLOSE() int32             { return 3594 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_COMMIT_ATOMIC() int32     { return 7690 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_CONVPATH() int32          { return 6666 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_CORRUPTFS() int32         { return 8458 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_DATA() int32              { return 8202 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_DELETE() int32            { return 2570 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_DELETE_NOENT() int32      { return 5898 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_DIR_CLOSE() int32         { return 4362 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_DIR_FSYNC() int32         { return 1290 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_FSTAT() int32             { return 1802 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_FSYNC() int32             { return 1034 }

// ...
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_GETTEMPPATH() int32     { return 6410 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_LOCK() int32            { return 3850 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_MMAP() int32            { return 6154 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_NOMEM() int32           { return 3082 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_RDLOCK() int32          { return 2314 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_READ() int32            { return 266 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_ROLLBACK_ATOMIC() int32 { return 7946 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_SEEK() int32            { return 5642 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_SHMLOCK() int32         { return 5130 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_SHMMAP() int32          { return 5386 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_SHMOPEN() int32         { return 4618 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_SHMSIZE() int32         { return 4874 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_SHORT_READ() int32      { return 522 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_TRUNCATE() int32        { return 1546 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_UNLOCK() int32          { return 2058 }

// ...
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_VNODE() int32             { return 6922 }
func (SqliteExtendedResultCodeStruct) SQLITE_IOERR_WRITE() int32             { return 778 }
func (SqliteExtendedResultCodeStruct) SQLITE_LOCKED_SHAREDCACHE() int32      { return 262 }
func (SqliteExtendedResultCodeStruct) SQLITE_LOCKED_VTAB() int32             { return 518 }
func (SqliteExtendedResultCodeStruct) SQLITE_NOTICE_RECOVER_ROLLBACK() int32 { return 539 }
func (SqliteExtendedResultCodeStruct) SQLITE_NOTICE_RECOVER_WAL() int32      { return 283 }
func (SqliteExtendedResultCodeStruct) SQLITE_OK_LOAD_PERMANENTLY() int32     { return 256 }
func (SqliteExtendedResultCodeStruct) SQLITE_READONLY_CANTINIT() int32       { return 1288 }
func (SqliteExtendedResultCodeStruct) SQLITE_READONLY_CANTLOCK() int32       { return 520 }
func (SqliteExtendedResultCodeStruct) SQLITE_READONLY_DBMOVED() int32        { return 1032 }
func (SqliteExtendedResultCodeStruct) SQLITE_READONLY_DIRECTORY() int32      { return 1544 }
func (SqliteExtendedResultCodeStruct) SQLITE_READONLY_RECOVERY() int32       { return 264 }
func (SqliteExtendedResultCodeStruct) SQLITE_READONLY_ROLLBACK() int32       { return 776 }
func (SqliteExtendedResultCodeStruct) SQLITE_WARNING_AUTOINDEX() int32       { return 284 }
