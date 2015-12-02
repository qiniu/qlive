# 简介

根据流Id获取该流的FLV直播播放地址

# 命令

```
qlive stream flv-live <StreamId>
```

# 参数

|名称|描述|
|--------|---------|
|StreamId|直播流的Id|

# 示例

获取直播流的HLS播放地址

```
qlive stream flv-live z1.jinxinxin.565f0113fb16df7e0a010361
```

输出：

```
ORIGIN http://pili-live-hdl.live.golanghome.com/jinxinxin/565f0113fb16df7e0a010361.flv
```