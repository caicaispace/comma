#### 常见问题

- elasticSearch Docker启动报错：max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]

解决办法1：

在宿主主机中执行

```
[root@localhost ~]# sysctl -w vm.max_map_count=262144
```

```
[root@localhost ~]# sysctl -a|grep vm.max_map_count
vm.max_map_count = 262144
```

这种重启会失效

解决办法2：

在/etc/sysctl.conf文件最后添加一行：

```
vm.max_map_count=262144
```
立即生效, 执行：

```
/sbin/sysctl -p
```

### 性能分析 

#### pprof

- [使用pprof分析cpu占用过高问题](https://blog.csdn.net/xmcy001122/article/details/103939923)

```
apt install graphviz
go tool pprof main cpu.prof

top # 查看 top
png ### 输出详细报告
```

### 扩展阅读

- [jiagu自然语言处理工具](https://github.com/bububa/jiagu)
- [高性能多语言 NLP 和分词](https://github.com/go-ego/gse)
- [Go efficient multilingual NLP and text segmentation](https://github.com/go-ego/gse)
- [美团搜索中查询改写技术的探索与实践](https://tech.meituan.com/2022/02/17/exploration-and-practice-of-query-rewriting-in-meituan-search.html)
- [新词发现（一）：基于统计](https://www.cnblogs.com/en-heng/p/6699531.html)
- [深入理解 NLP 的中文分词：从原理到实践](https://juejin.cn/book/6844733812102922247)
- [重新写了之前的新词发现算法：更快更好的新词发现](https://spaces.ac.cn/archives/6920)
- ["新词发现"算法探讨与优化-SmoothNLP](https://zhuanlan.zhihu.com/p/80385615)
- [ElasticSearch性能调优](https://zhuanlan.zhihu.com/p/55092525)

### ide

#### Goland

#### vscode

- Collapse code shortcuts
```
ctrl + k + 0 : Fold all levels (namespace , class , method , block)
ctrl + k + 1 : namspace
ctrl + k + 2 : class
ctrl + k + 3 : methods
ctrl + k + 4 : blocks
ctrl + k + [ or ] : current cursor block
ctrl + k + j : UnFold
```

#### Clear IDE cache

```
1. Press Ctrl + Shift + P
2. type command Clear Editor History
3. Press Enter
```
