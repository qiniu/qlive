# 简介

根据流Id获取该流的RTMP推流地址

# 命令

```
qlive stream rtmp-pub <StreamId>
```

# 参数

|名称|描述|
|--------|---------|
|StreamId|直播流的Id|

# 示例

获取流的推流地址

```
qlive stream rtmp-pub z1.jinxinxin.565f0113fb16df7e0a010361
```

输出：

```
rtmp://pili-publish.live.golanghome.com/jinxinxin/565f0113fb16df7e0a010361?key=702ba5beb522a05a
```