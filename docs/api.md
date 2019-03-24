简单起见，所有的 response 都没有 error 字段，通过 http status code 来代表其错误内容

## POST /

### request

POST /

```
{
    "url": "https://annatarhe.com"
}
```


### response

```
{
    "url": "https://s.yourdomain.com/1"
}
```

## GET /{url: string}

### response

301 redirect

如果找到了内容会直接 301 跳转到相应的网址

