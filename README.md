## sqlite

```bash
$ sqlite3 output/config.db 'select sqlite_version()'
3.39.5
```

```bash
$ sqlite3 output/config.db .tables
ca
```

```bash
$ sqlite3 output/config.db '.schema ca'
CREATE TABLE ca (
        id           INTEGER PRIMARY KEY AUTOINCREMENT,
        name         VARCHAR(32)  NOT NULL,
        private_key  VARCHAR(256) NOT NULL
);
```