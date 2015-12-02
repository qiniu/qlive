# 简介

设置工具的`AccessKey`，`SecretKey`和`Hub`信息

# 命令

```
qlive hub reg -ak <AccessKey> -sk <SecretKey> -hub <HubName>
```

# 参数

|名称|描述|
|--------|-------|
|AccessKey|直播账号的`AccessKey`|
|SecretKey|直播账号的`SecretKey`|
|HubName|直播账号的直播通道名称|

# 示例

假设账号的`AccessKey`和`SecretKey`分别为`xxx`和`yyy`，直播通道名称`zzz`

```
qlive hub reg -ak xxx -sk yyy -hub zzz
```