# opreant

```
query {
  GetMetricsByUser(userId:"1"){
    name
  }
  GetMetricsByUserFilterByTreshold(userId:"1",highTreshold:100,lowTreshold:10){
    name
  }
}
```

##### Schema

https://github.com/rotemvolfo/operant/blob/master/graph/schema.graphqls

##### Resolver functions

https://github.com/rotemvolfo/operant/blob/master/graph/schema.resolvers.go

Table schema
https://github.com/rotemvolfo/operant/tree/main/mysql/migrations
https://github.com/rotemvolfo/operant/blob/main/internal/metrics/metrics.go

##### Problem 2.

Let’s take the following schema:
* Entity Service has unique user_id ,name, region
* Entity metric has unique name, high_threshold (int), low_threshold(int), current (int)
* A service can have many metrics

We want to build a graphql backend exposing this data to provide for the following queries
* Get all metric names for a given user
* Get all metric names for a given user that satisfies given value between
high_threshold and low_threshold

Please provide the Graphql queries for the above. You can decide to use a database of your choice for storing the data in the schema. Seed it with 10 or fewer samples.
