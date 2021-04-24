# Usage
Help:
```
  -i, --interval string    given files truncate interval in cron notation (default "@every 1h")
  -v, --log-level string   log level (default "info")
```
You can specify one or more files and set truncate interval. Example:
```
$ ./ truncate -i "@every 4h" path/to/file1 path/to/file2
```
