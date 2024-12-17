# gurl

This tls-client seems to be trash. Getting

```
go run . -url http://builtwith.com
Error executing HTTP request: Get "https://builtwith.com/": net/http: HTTP/1.x transport connection broken: malformed HTTP response "\x00\x00\x12\x04\x00\x00\x00\x00\x00\x00\x03\x00\x00\x00\x80\x00\x04\x00\x01\x00\x00\x00\x05\x00\xff\xff\xff\x00\x00\x04\b\x00\x00\x00\x00\x00\x7f\xff\x00\x00\x00\x00\b\a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01"

go run . -url http://github.com   
Error executing HTTP request: Get "https://github.com/": EOF
```

So that means this `tls-client` or `fhttp` has weird internal bugs and can't handle upgrades to HTTP, or we have some shitty config. The decompression also didn't work automatically.
