# go-gremlin

Installation
==========
```
go get github.com/go-gremlin/gremlin
```

Usage
======
Import the library eg `import "github.com/go-gremlin/gremlin"`.

```go
	if err := gremlin.NewCluster("ws://dev.local:8182/gremlin", "ws://staging.local:8182/gremlin"); err != nil {
		// handle error
	}
```

To actually run queries against the database, make sure the package is imported and issue a gremlin query like this:-
```go
	data, err := gremlin.Query(`g.V()`).Exec()
	if err != nil  {
		// handle error
	}
```
You can also execute a query with bindings like this:-
```go
	data, err := gremlin.Query(`g.V().has("name", userName).valueMap()`).Bindings(gremlin.Bind{"userName": "john"}).Exec()
```

You can also execute a query with Session, Transaction and Aliases 
```go
	aliases := make(map[string]string)
	aliases["g"] = fmt.Sprintf("%s.g", "demo-graph")
	session := "de131c80-84c0-417f-abdf-29ad781a7d04"  //use UUID generator
	data, err := gremlin.Query(`g.V().has("name", userName).valueMap()`).Bindings(gremlin.Bind{"userName": "john"}).Session(session).ManageTransaction(true).SetProcessor("session").Aliases(aliases).Exec()
```

Authentication
===
```go
	auth := gremlin.OptAuthUserPass("myusername", "mypass")
	client, err := gremlin.NewClient("ws://remote.example.com:443/gremlin", auth)
	data, err = client.ExecQuery(`g.V()`)
	if err != nil {
		panic(err)
	}
	doStuffWith(data)
```
