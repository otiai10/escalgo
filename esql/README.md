escalgo/esql
============

SQL like syntax sugar for Elasticsearch Query DSL

Examples
============

```go
req := Select([]Tweet{}).From("idx_twitter").
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

req.Path()          // "/idx_twitter/tweet/_search"
req.Query()         // "?size=10&from=20&pretty=1"
req.Body().String() // json string
req.Body().Bytes()  // []byte of json string
req.Body().Map()    // map[string]interface{}
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
