# bluray
GO bluray information parser library

## Example usage
### Listing PlayLists in (more or less) human readable format
```go
b, _ := bluray.OpenBDMV(http.Dir("/mnt/Movies/Awesome BluRay Movie/BDMV"))
bytes, _ := json.Marshal(b.PlayLists)
fmt.Printf("%s\n", bytes)
```
