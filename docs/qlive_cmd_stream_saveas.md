# 简介

根据流Id和起始，结束时间将直播流转存为视频格式，并返回视频地址

# 命令

```
qlive stream saveas <StreamId> -n <Name> -f <Format> -s <Start> -e <End> -c <NotifyUrl>
```

# 参数

|名称|描述|
|--------|---------|
|StreamId|直播流的Id|
|Name|另存为的视频名称|
|Format|另存为的视频格式|
|Start|起始时间，可以为Unix时间戳或者如`20151202233001`这样的字符串|
|End|结束时间，可以为Unix时间戳或则如`20151202233201`这样的字符串|
|NotifyUrl|处理结果的通知地址，服务器处理完成之后会发送POST请求将结果发送到该地址|

# 示例

直播转存视频

```
qlive stream saveas z1.jinxinxin.5645a8e1d409d2a0b10011a7 -s 20151202160000 -e 20151202163030 -n hello.mp4 -f mp4
```

输出：

```
M3u8Url:	 http://pili-static.live.golanghome.com/recordings/z1.jinxinxin.5645a8e1d409d2a0b10011a7/hello.m3u8
TargetUrl:	 http://pili-media.live.golanghome.com/recordings/z1.jinxinxin.5645a8e1d409d2a0b10011a7/hello.mp4

See http://api.qiniu.com/status/get/prefop?id=z1.565f0f84f51b8236d007877c
```