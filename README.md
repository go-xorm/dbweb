# Xorm Admin

Admin is a web based database admin tool like phpmyadmin. It' written via 
[xorm](github.com/go-xorm/xorm), [xweb](github.com/go-xweb/xweb), [ql](github.com/lunny/ql/driver).

# Database Supports

* mysql
* postgres
* sqlite3

# Installation

```Go
go get github.com/go-xorm/admin
go install github.com/go-xorm/admin
```

# Run

```Shell
admin
```

Then visit http://localhost:8989/