esdsl/essql
============

SQL like syntax sugar for Elasticsearch Query DSL

Examples
============

```go
uri := Select([]Tweet{}).From("idx_twitter").
    Where(Clause{
        "name": Like("otiai"),
        "age":  Range(25, 35),
        "lang": Must("golang"),
        "city": In("Tokyo", "Boston"),
    }).
    OrderBy(Clause{
        "created": Desc(),
    }).
    Limit(10).Offset(20).
    Pretty()

uri.Path()          // "/idx_twitter/tweet/_search"
uri.Query()         // "?size=10&from=20&pretty=1"
uri.Body().String() // json string
uri.Body().Bytes()  // []byte of json string
uri.Body().Map()    // map[string]interface{}
/*
{
    "query": {
        "filtered": {
            "query": {
                "match": {
                    "name": "otiai"
                }
            },
            "filter": {
                "and": [
                    {
                        "range": {
                            "age": {
                                "gt": 25,
                                "lt": 35
                            }
                        },
                        "term: {
                            "lang": "golang"
                        },
                        "terms": {
                            "city": [
                                "Tokyo", "Boston"
                            ],
                            "execution": "or"
                        },
                    }
                ]
            }
        }
    },
    "sort": [
        {
            "created": "desc"
        }
    ]
}
*/
```
