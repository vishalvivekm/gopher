- gorm's preferred type used to represent tables? struct ( though it lets us use map[string]interface{})
- to retrieve all records from a table?: Find()
- Object :  Record (row)
  Model :  Table
  Field :  Column
 
 
 - with db.Where(), we can use Struct,String, Slice and map conditions (https://gorm.io/docs/query.html#Conditions)
 - Gorm config is It is a struct having various fields allowing us to enable or disable configuration options for the GORM connection.
 - To  retrieve a single record from a table: db.Last(), db.First(), db.Take()

- When working with SQLite databases, it is possible to create and connect to an in-memory database named :memory:
// db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

```
SQLite in-memory databases are databases stored entirely in memory, not on disk. 
Use the special data source filename :memory: to create an in-memory database.
 When the connection is closed, the database is deleted.
When using :memory: , each connection creates its own database
```