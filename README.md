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

### Assumptions:
* It is okay to select some invalid URLs after a basic check
  - This assumption makes the selection of URLs faster
  - This makes the whole process faster as the calls to these URLs are parallel
  - The invalid may take more than 3 seconds to reply and throw an error but all the valid URLs are queried by that time

* It is okay to support both **GET** and **POST** methods for the endpoints.
  - Currently it supports both **GET** and **POST**

* 500 ms is the time for all calculations(i.e. GET calls and sorting)
  - So the time to write the response is independent of the time constraint
  - Time constraint only applies to the main business logic of making the GET calls and sorting result

* The output format expected from the dummy servers(like demo/demo.go) is
```javascript
{
   "numbers" :[1,2,3,5,7,8]
}
```

### Note:
* The directory demo has the problem statement and dummy server for numbers
