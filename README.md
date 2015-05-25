escalgo
=======

Elasticsearch query builder for Golang.

Less is more.

Sample
======

want

```json
{
  "query": {
    "filtered": {
      "query": {
        "bool": {
          "should": [
            {
              "terms": {
                "lang": [
                  "go",
                  "js"
                ]
              }
            },
            {
              "multi_match": {
                "fields": [
                  "hobby",
                  "description"
                ],
                "query": "japan anime"
              }
            }
          ]
        }
      }
    }
  },
  "sort": [
    {
      "created": "desc"
    }
  ]
}
```

do

```go
search := Search("users_index", "user")

filtered := &Filtered{}
filtered.Query().Bool().Should(Terms{"lang", []string{"go", "js"}})
filtered.Query().Bool().Should(MultiMatch{"japan anime", []string{"hobby", "description"}})

sorts := []Sortable{&Sort{Field: "created", Desc: true}}

search.Set(filtered)
search.Set(sorts)

b, _ := json.MarshalIndent(search, "", "  ")
fmt.Println(string(b))
```
