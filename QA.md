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
