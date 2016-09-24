## numbers

### How to run
* `go get githuh.com/abhi11/numbers`
* `cd ~/$GOPATH/src/githuh.com/abhi11/numbers`
* `go build && ./numbers`

### API
* Endpoint: /numbers
* Query Params:
  - u : url
* Usage: `http://localhost:8888/numbers?u=http://example.com/primes&u=http://foobar.com/fibo`
* Output:
```javascript
{
   "numbers" :[1,2,3,5,7,8]
}
```
