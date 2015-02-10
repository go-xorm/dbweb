# DBWeb

DBWeb is a web based database admin tool like phpmyadmin. It' written via 
[xorm](http://github.com/go-xorm/xorm), [tango](http://github.com/lunny/tango), [nodb](http://github.com/lunny/nodb).

# Screenshot

![dbweb](screenshot.png)

# Database Supports

* mysql
* postgres
* sqlite3 : build tag -sqlite3

# Installation

```Go
go get github.com/go-xorm/dbweb
go install github.com/go-xorm/dbweb
```

# Run

```Shell
./dbweb -help

dbweb version 0.1

  -debug=false: enable debug mode
  -help=false: show help
  -https=false: enable https
  -port=8989: listen port
```

Then visit http://localhost:8989/

The default user is admin and password is admin. You can change it after you loged in.