# DBWeb

DBWeb is a web based database admin tool like phpmyadmin. It' written via 
[xorm](http://github.com/go-xorm/xorm), [tango](http://github.com/lunny/tango), [nodb](http://github.com/lunny/nodb).

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
dbweb
```

Then visit http://localhost:8989/