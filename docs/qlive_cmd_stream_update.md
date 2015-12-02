# 简介

更新流的`PublishKey`或者`PublishSecurity`信息

# 命令

```
qlive stream update <StreamId> -pbk <PublishKey> -pbs <PublishSecurity>
```

# 参数

|名称|描述|可选|
|--------|---------|---------|
|StreamId|直播流的Id|N|
|PublishKey|直播流的新推流密钥|Y|
|PublishSecurity|直播流的新推流验证方式，可选为`static`和`dynamic`|Y|

# 示例

更新流的推流密钥

```
qlive stream update z1.jinxinxin.test_qiniu_live -pbk test_qiniu_key2
```