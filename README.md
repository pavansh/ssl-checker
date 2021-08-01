# ssl-checker
A command-line tool for verify and sent alert on SSL expire

# Usage

```
go get github.com/pavansh/ssl-checker
ssl-checker -h
  -host string
        hostname (ex: -host example.com,sample.com (default "example.com")

example:
ssl-checker -host example.com
ssl-checker -host example.com,abc.com
```