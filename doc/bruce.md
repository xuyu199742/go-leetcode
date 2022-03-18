## Docker
### 常用命令
列出所有的容器 ID
```shell
 docker ps -aq
```
停止所有容器
```shell
 docker stop $(docker ps -aq)
```

删除所有的容器
```shell
docker rm $(docker ps -aq)
```

删除所有镜像文件
```shell
docker rmi $(docker images -q)
```

### 踩坑

## Linux


## Elasticsearch

### 踩坑
- 挂载目录没有给权限
```shell
'ElasticsearchException[failed to bind service] nested: AccessDeniedException[/usr/share/elasticsearch/data/nodes]'
```

















