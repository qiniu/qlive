# 简介

根据流Id获取该流的RTMP直播播放地址

# 命令

```
qlive stream rtmp-live <StreamId>
```

# 参数

|名称|描述|
|--------|--------|
|StreamId|直播流的Id|

# 示例

获取直播流的RTMP播放地址

```
qlive stream rtmp-live z1.jinxinxin.565f0113fb16df7e0a010361
```

输出：

```
ORIGIN rtmp://pili-live-rtmp.live.golanghome.com/jinxinxin/565f0113fb16df7e0a010361
```