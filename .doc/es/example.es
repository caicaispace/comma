POST search_test/_search
{
  "query": {
    "function_score": {
        "query": {
            "bool" : {
                "must"   : {
                    "multi_match" : {
                        "query" : "安阳",
                        "fields" : [
                            "name^33",
                            "nickname",
                            "job"
                        ],
                        "type" : "most_fields",
                        // "operator" : "and",
                        "minimum_should_match" : "100%"
                    }
                }
                // "must_not" : {},
                // "filter"   : [
                //     {"term" : {"is_display": false}}
                // ]
            }
        }
    }
  },
  "size": 10,
  "from": 0,
  "sort": []
}


// # 创建索引
PUT search_test
{
  "mappings": {
    "all": {
      "dynamic": "strict",
      "include_in_all": false,
      "_all": {
        "enabled": false
      },
      "properties": {
        "address": {
            "type": "nested",
            "properties": {
                "area_code": {
                    "type": "keyword",
                    "index": "no"
                },
                "area_name": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart",
                    "fielddata": true
                },
                "city_code": {
                    "type": "keyword",
                    "index": "no"
                },
                "lat": {
                    "type": "float"
                },
                "lng": {
                    "type": "float"
                },
                "name":{
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart",
                    "fielddata": true
                },
                "short_name": {
                    "type": "text",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart",
                    "fielddata": true
                },
                "zip_code": {
                    "type": "integer"
                }
             }
        },
        "age": {
          "type": "integer"
        },
        "airlineinfo": {
            "type": "nested",
            "properties": {
                "code": {
                    "type": "keyword"
                },
                "name": {
                    "type": "keyword"
                }
            }
        },
        "airport": {
            "type": "nested",
            "properties": {
                "city":{
                    "type": "keyword"
                },
                "iata_code":{
                    "type": "keyword"
                },
                "icao_code":{
                    "type": "keyword"
                },
                "name":{
                    "type": "keyword"
                },
                "pinyin":{
                    "type": "keyword"
                }
            }
        },
        "capturetime": {
          "type": "keyword"
        },
        "citycode": {
          "type": "keyword"
        },
        "color":{
          "type": "keyword"
        },
        "date": {
          "type": "keyword"
        },
        "deviceid": {
          "type": "keyword"
        },
        "email": {
          "type": "keyword"
        },
        "flightseat": {
          "type": "keyword"
        },
        "idcard": {
          "type": "keyword"
        },
        "imei": {
          "type": "keyword"
        },
        "imid": {
          "type": "keyword"
        },
        "imsi": {
          "type": "keyword"
        },
        "ipv4": {
          "type": "keyword"
        },
        "ipv6": {
          "type": "keyword"
        },
        "job": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart",
          "fielddata": true
        },
        "mac": {
          "type": "keyword"
        },
        "meid": {
          "type": "keyword"
        },
        "mobilephone": {
          "type": "keyword"
        },
        "name": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart",
          "fielddata": true
        },
        "nickname": {
          "type": "text",
          "analyzer": "ik_max_word",
          "search_analyzer": "ik_smart",
          "fielddata": true
        },
        "password": {
          "type": "keyword"
        },
        "sex": {
          "type": "keyword"
        },
        "specialphone": {
          "type": "keyword"
        },
        "telphone": {
          "type": "keyword"
        },
        "trainseat": {
          "type": "keyword"
        },
        "traintrips": {
          "type": "keyword"
        },
        "url": {
          "type": "keyword"
        },
        "useragent": {
          "type": "keyword"
        },
        "username": {
          "type": "keyword"
        },
        "voyage": {
          "type": "keyword"
        },
        "website": {
          "type": "keyword"
        }
      }
    }
  },
  "settings": {
    "index": {
      "number_of_shards": "2",
      "number_of_replicas": "1",
      "max_result_window": 10000
    },
    "analysis": {
      "analyzer": {
        "payload_analyzer": {
          "type": "custom",
          "tokenizer": "whitespace",
          "filter": "delimited_payload_filter"
        }
      }
    }
  }
}