# splinter
A small RAT

## Examples
To use the RAT the C2 server has to be run first:
```
$ go run server/server.go
```
Next you can run the Implant itself:
```
$ go run implant/implant.go
```
Finally you can run the admin client and order commads as arguments:
```
$ go run client/client.go 'ls'
```
