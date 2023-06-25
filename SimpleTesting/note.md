# Go Simple Testing

```bash
# run app
go run .

# run test 
go test .

# test cover percentage
go test -cover .

# test cover output file
go test -coverprofile=app.out .

# see cover in html,the file has been existed
go tool cover -html=coverage.out
```
