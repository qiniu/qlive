# 简介

根据流Id和起始，结束时间获取直播流的HLS回放地址

# 命令

```
qlive stream hls-play <StreamId> -s <Start> -e <End>
```

# 参数

|名称|描述|
|--------|---------|
|StreamId|直播流的Id|
|Start|起始时间，可以为Unix时间戳或者如`20151202233001`这样的字符串|
|End|结束时间，可以为Unix时间戳或则如`20151202233201`这样的字符串|

# 示例

获取直播流的回放HLS地址

```
qlive stream hls-play z1.jinxinxin.5645a8e1d409d2a0b10011a7 -s 20151202160000 -e 20151202163030
```

输出：

```
ORIGIN http://pili-playback.live.golanghome.com/jinxinxin/5645a8e1d409d2a0b10011a7.m3u8?start=1449043200&end=1449045030
```