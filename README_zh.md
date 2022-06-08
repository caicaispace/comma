# goaway

中文分词、违禁词过滤、汉字转拼音

### 功能

- 中文分词
- 繁简转换
- ElasticSearch 中间件
- 违禁词
- 汉字转拼音
- 管理ui
- jsonrpc

### 管理界面

![base ui](.doc/img/base_ui_zh.png)

### TODO

#### 2.0

- [ ] 新词发现
- [ ] grpc
- [ ] 分布式

#### docker

- 启动容器：

```shell
docker-compose up -d
```

- 滚动启动:

```shell
docker-compose pull && docker-compose up -d
```

### 扩展阅读

- [jiagu自然语言处理工具](https://github.com/bububa/jiagu)
- [高性能多语言 NLP 和分词](https://github.com/go-ego/gse)
- [Go efficient multilingual NLP and text segmentation](https://github.com/go-ego/gse)
- [美团搜索中查询改写技术的探索与实践](https://tech.meituan.com/2022/02/17/exploration-and-practice-of-query-rewriting-in-meituan-search.html)
- [新词发现（一）：基于统计](https://www.cnblogs.com/en-heng/p/6699531.html)
- [深入理解 NLP 的中文分词：从原理到实践](https://juejin.cn/book/6844733812102922247)