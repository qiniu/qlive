# 简介

根据流Id和截取图片时间从直播流中截取一张图片，并返回截图地址

# 命令

```
qlive stream snapshot <StreamId> -n <Name> -f <Format> -t <Time> -c <NotifyUrl>
```

# 参数

|名称|描述|
|--------|---------|
|StreamId|直播流的Id|
|Name|截取的图片名称|
|Format|截取的图片格式|
|Time|截取时刻，可以为Unix时间戳或者如`20151202233001`这样的字符串|
|NotifyUrl|处理结果的通知地址，服务器处理完成之后会发送POST请求将结果发送到该地址|

# 示例

直播转存视频

```
qlive stream snapshot z1.jinxinxin.5645a8e1d409d2a0b10011a7 -t 20151202163030 -n hello.jpg -f jpg
```

输出：

```
TargetUrl: http://pili-static.live.golanghome.com/snapshots/z1.jinxinxin.5645a8e1d409d2a0b10011a7/hello.jpg

See http://api.qiniu.com/status/get/prefop?id=z1.565f106cf51b8236d0078877
```