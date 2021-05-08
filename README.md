# tables

a simple go library to create printable tables


## Use

Tables can be created as easily as
```go
t := tables.NewTable().
    AppendCells("Name", "Firstname", "Birthyear", "Profession").
    AppendCells("Martel", "Nils", "1997", "Developer").
    AppendCells("Doe", "Jane", "1993", "Tourismexpert").
    AppendCells("Parker", "Peter", "1992", "Spiderman")

fmt.Println(t)
```
