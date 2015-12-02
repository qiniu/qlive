# 简介

根据流Id获取一个流的基本信息

# 命令

```
qlive hub get-stream <StreamId>
```

# 参数

|名称|描述|
|--------|-------|
|StreamId|直播流的Id|

# 示例

获取直播流`z1.jinxinxin.565f0113fb16df7e0a001361`的信息

```
qlive hub get-stream z1.jinxinxin.565f0113fb16df7e0a001361
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


# 备注

为了减轻输入子命令的痛苦，上面的`get-stream`也可以缩短为`gs`，比如如下的命令也是一样的效果：

```
qlive hub gs z1.jinxinxin.565f0113fb16df7e0a001361
```