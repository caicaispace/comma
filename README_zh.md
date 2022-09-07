# comma

中文分词、违禁词过滤、汉字转拼音

### 特性

*   中文分词（同义词、停词、上位词、下位词、节日词、违禁词）

*   繁简转换

*   汉字转拼音

*   ElasticSearch 代理中间件（http、jsonrpc、grpc）

*   管理ui

*   监控（grafana）

*   日志收集（loki）

### UI

![base ui](.doc/img/base_ui_zh.png)

### 监控

![grafana](.doc/img/grafana.png)

### 日志

![grafana](.doc/img/loki.png)

### TODO

#### 2.0

*   [ ] 新词发现

*   [ ] 拼写纠错

*   [ ] 精排混排

*   [ ] 分布式

#### docker

*   启动容器：

```shell
docker-compose up -d
```

*   滚动启动:

```shell
docker-compose pull && docker-compose up -d
```
