elql/sdsl
=============

Search DSL for EsQL

Sample
=============

```go
dsl := SearchDSL{
  Query: Query{
  },
  Filtered: Filtered{
    Query: Query{
    },
    Filter: Filter{
      And: Filterables{
        Terms{
          Filed: "language",
          Values: []interface{}{"go","js"},
          Execution: "and",
        },
        Range{
          Filed: "age",
          Gte:   18,
        },
      },
    },
  },
  PostFilter: Filter{
    Or: []Filter{
    },
  },
}
```
