# json-selector

json-selector lets you select certain data from json string by string. This enables you to create a dynamic json parser easily.

## install

```
go get github.com/sunho/json-selector
```
## example

```go
text := []byte(`
    {
        "payload":{
            "id":"1234",
            "pw":"1234"
        },
        "addresses":{
            "a":{
                "ip":"127.0.0.1",
                "port":1234
            },
            "b":{
                "ip":"127.0.0.1",
                "pr":4567
            }
        }
    }
`)

id, _ := selector.Select(text, "payload.id")
fmt.Println(string(id)) // 1234

type Address struct {
    Ip string `json:"ip"`
    Pw int    `json:"pw"`
}

adrs := map[string]Address{}
adrs2, _ := selector.Select(text, "addresses")
json.Unmarshal(adrs2, &adrs)

fmt.Println(adrs["a"].Ip) // 127.0.0.1
```
