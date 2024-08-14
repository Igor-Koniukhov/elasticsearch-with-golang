## Sample of elasticsearch initiation, indexing docs, search query



### Commands for playing around in the Console Elastic Dev Tools (http://localhost:5601/app/dev_tools#/console)

```

PUT movies
{ 
    "settings": {
      "number_of_shards": 1,
      "number_of_replicas": 0
      
    }, 
 
    "mappings": {
      "properties": {
        "title":{
          "type": "text",
          "fields":{
            "en":{
              "type": "text",
              "analyzer": "english"
            },
            "es":{
              "type": "text",
              "analyzer": "spanish"
            },
            "pt":{
              "type": "text",
              "analyzer": "portuguese"
            }
          }
        },
          "year": {
          "type": "long"
        },
         "runningTime": {
          "type": "long"
        },
        "releaseDate": {
          "type": "date",
          "ignore_malformed": true
        },
         "rating": {
          "type": "double"
        },
        "actors": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "directors": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        },
        "genres": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        }
      }
    }
  
}

GET movies/_mapping


GET movies/_count


DELETE movies

GET movies/_search
{
  "size": 1000
}

GET movies/_doc/13

```
# elasticsearch-for-gopher
https://www.youtube.com/watch?v=j0RiWwef8Z8

https://github.com/riferrei/elasticsearch-for-gophers
