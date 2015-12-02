# 简介

创建一个新的直播流，可以选择性设置`PublishKey`和`PublishSecurity`信息

# 命令

```
qlive hub create-stream -t <Title> -pbk <PublishKey> -pbs <PublishSecurity>
```

# 参数

|名称|描述|可选|
|--------|---------|---------|
|Title|新直播流的名称，直播流全局唯一性标记，可以不指定，七牛服务器自动生成|Y|
|PublishKey|新直播流的推流密钥，可以不指定，七牛服务器自动生成|Y|
|PublishSecurity|新直播流的推流验证方式，可选为`static`和`dynamic`，默认不设置为`static`|Y|

# 示例

1.不指定任何可选参数创建流

```
qlive hub create-stream
```

结果：

```
{
    "id":"z1.jinxinxin.565f0113fb16df7e0a001361",
    "createdAt":"2015-12-02T14:32:51.608Z",
    "updatedAt":"2015-12-02T14:32:51.608Z",
    "title":"565f0113fb16df7e0a000361",
    "hub":"jinxinxin",
    "disabled":false,
    "publishKey":"701ba5beb522a05b",
    "publishSecurity":"static",
    "hosts":
        {
            "publish":
                {
                    "rtmp":"pili-publish.live.golanghome.com"
                },
            "live":
                {
                    "hdl":"pili-live-hdl.live.golanghome.com",
                    "hls":"pili-live-hls.live.golanghome.com",
                    "http":"pili-live-hls.live.golanghome.com",
                    "rtmp":"pili-live-rtmp.live.golanghome.com"
                },
            "playback":
                {
                    "hls":"pili-playback.live.golanghome.com",
                    "http":"pili-playback.live.golanghome.com"
                }
        }
}
```

2.指定`Title`和`PublishSecurity`

```
qlive hub create-stream -t test_qiniu_live -pbs dynamic
```

结果：

```
{
    "id":"z1.jinxinxin.test_qiniu_live",
    "createdAt":"2015-12-02T14:36:09.881Z",
    "updatedAt":"2015-12-02T14:36:09.881Z",
    "title":"test_qiniu_live",
    "hub":"jinxinxin",
    "disabled":false,
    "publishKey":"b6821d34f879d83a",
    "publishSecurity":"dynamic",
    "hosts":
        {
            "publish":
                {
                    "rtmp":"pili-publish.live.golanghome.com"
                },
            "live":
                {
                    "hdl":"pili-live-hdl.live.golanghome.com",
                    "hls":"pili-live-hls.live.golanghome.com",
                    "http":"pili-live-hls.live.golanghome.com",
                    "rtmp":"pili-live-rtmp.live.golanghome.com"
                },
            "playback":
                {
                    "hls":"pili-playback.live.golanghome.com",
                    "http":"pili-playback.live.golanghome.com"
                }
        }
}
```

# 备注

为了减轻输入子命令的痛苦，上面的`create-stream`也可以缩短为`cs`，比如如下的命令也是一样的效果：

```
qlive hub cs -t test_qiniu_live -pbs static -pbk test_qiniu_key
```