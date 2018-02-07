# kelly

[![Build Status](https://secure.travis-ci.org/missdeer/kelly.png)](https://travis-ci.org/missdeer/kelly)
[![Total views](https://sourcegraph.com/api/repos/github.com/missdeer/kelly/counters/views.png)](https://sourcegraph.com/github.com/missdeer/kelly)
[![status](https://sourcegraph.com/api/repos/github.com/missdeer/kelly/.badges/status.png)](https://sourcegraph.com/github.com/missdeer/kelly)
[![Gobuild Download](http://gobuild.io/badge/github.com/missdeer/kelly/downloads.svg)](http://gobuild.io/github.com/missdeer/kelly)

Yiili community which is a part of [Kelly project](https://github.com/missdeer/kelly), it is the server side, AKA backend, of the project. The source code is based on WeTalk project, thanks those guys for their great work.

### Usage

```
go get -u github.com/missdeer/kelly
cd $GOPATH/src/github.com/missdeer/kelly
```

I suggest you [update all Dependencies](#dependencies)

Copy `conf/global/app.ini` to `conf/app.ini` and edit it. All configure has comment in it.

The files in `conf/` can overwrite `conf/global/` in runtime.


**Run kelly**

```
bee run watchall
```

### Dependencies

Contrib

* Beego [https://github.com/astaxie/beego](https://github.com/astaxie/beego)
* Social-Auth [https://github.com/beego/social-auth](https://github.com/beego/social-auth)
* Compress [https://github.com/beego/compress](https://github.com/beego/compress)
* i18n [https://github.com/beego/i18n](https://github.com/beego/i18n)
* Mysql [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
* goconfig [https://github.com/Unknwon/goconfig](https://github.com/Unknwon/goconfig)
* fsnotify [https://github.com/fsnotify/fsnotify](https://github.com/fsnotify/fsnotify)
* resize [https://github.com/nfnt/resize](https://github.com/nfnt/resize)
* blackfriday [https://github.com/russross/blackfriday](https://github.com/russross/blackfriday)

Update all Dependencies

```
go get -u github.com/beego/social-auth
go get -u github.com/beego/compress
go get -u github.com/beego/i18n
go get -u github.com/go-sql-driver/mysql
go get -u github.com/Unknwon/goconfig
go get -u github.com/fsnotify/fsnotify
go get -u github.com/nfnt/resize
go get -u github.com/russross/blackfriday
```

### Static Files

kelly use `Google Closure Compile` and `Yui Compressor` compress js and css files.

So you could need Java Runtime. Or close this feature in code by yourself.

### Contact

Maintain by [missdeer](http://minidump.info/)

## License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
