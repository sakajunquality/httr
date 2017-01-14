httr (HTTp Response)
=================================

Overview
------------
Command line tool for displaying HTTP response headers.

![image](./doc/image.png)


Installation
------------
#### Using homebrew,
```
$ brew tap sakajunquality/homebrew-quality
$ brew install httr
```

#### Using go get,
```
$ go get github.com/codegangsta/cli
$ go get github.com/fatih/color
$ go get github.com/sakajunquality/httr
```



Usage
------------
#### Normal output,
```
$ httr example.com
$ httr https://example.com
$ httr https://example.com:443
```

#### JSON output,
```
$ httr -j example.com
$ httr --json https://example.com
```

##### combination with jq
```
$ httr -j example.com | jq .
```


## Future Perspectives

- make it cool
- add testing
- methods other than GET
- custom request header / body
- option for displaying response body


Contributing
-----
Always welcome for contributing



License & Authors
-----------------
- Author:: @sakajunquality
- License:: MIT