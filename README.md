# sssssssshort
short url

用来演示的短链服务~

## deploy

```bash
$ docker-compose up api-server
$ docker ps
$ docker exec -it xxx(mysql 实例) mysql
#> (贴入 migration.sql)
```

然后可以访问啦~

### POST /g

request

```
{
    url: "https://some-very-long.domain.name/and/has/very/long/path?and=many&many=query&to=encode"
}
```

response

```
{
    url: "https://your.domain.name/r/1"
}
```


## 配置

可以自己复制出 `src/config.release.go`， 然后改动内部对应的配置呢~
