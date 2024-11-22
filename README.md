# Overview

This is a quick demo. It uses sqlite as database.

# Running It

The only environment variable required is DB_NAME for the sqlite DB. Once app is running, it will be accessible on port 7000.
You can start app using the following command from the root folder. 

```bash
DB_NAME='EXAMPLE_DB_NAME' go run pkg/main.go
```

You can replace EXAMPLE_DB_NAME with any name appropriate for an sqlite DB, say "example.db".
