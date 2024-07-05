
## Soft Delete


If your model includes a gorm.DeletedAt field (which is included in gorm.Model), it will get soft delete ability automatically!

When calling Delete, the record WON’T be removed from the database, but GORM will set the DeletedAt‘s value to the current time, and the data is not findable with normal Query methods anymore.
```go
// user's ID is `111`
db.Delete(&user)
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// Batch Delete
db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Soft deleted records will be ignored when querying
db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
```
If you don’t want to include gorm.Model, you can enable the soft delete feature like:
```go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```
## Find soft deleted records

You can find soft deleted records with Unscoped
```go
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```
## gorm.Model{}
> from :- https://gorm.io/docs/delete.html#Soft-Delete
```go
// GORM provides a predefined struct named gorm.Model, which includes commonly used fields:

// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
// more at :https://gorm.io/docs/models.html#gorm-Model
```
