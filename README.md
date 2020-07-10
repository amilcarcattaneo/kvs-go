## Start Server

`cd kvs-go/server && go run main.go`

## API REST

### Get value from key
- Status Code:
  1. **200**: if the key has a value.
  2. **400**: if the key param is undefined or contains only white spaces.
  3. **404**: if the key doesn’t have a value.

```
curl --location --request GET 'http://localhost:8080/key/test' \
--header 'Content-Type: application/json'
```

### Post new key value pair
- Status Code:
  1. **201**: if the key value pair was accepted.
  2. **400**: if the key or the value are undefined or contains only white spaces.

```
curl --location --request POST 'http://localhost:8080/key/test/value/test' \
--header 'Content-Type: application/json' \
```
