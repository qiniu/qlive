# 简介

根据流Id获取该流的当前状态

# 命令

```
qlive stream status <StreamId>
```

# 参数

|名称|描述|
|--------|---------|
|StreamId|直播流的Id|

# 示例

获取流的状态

```
qlive stream status z1.jinxinxin.565f0113fb16df7e0a010361
```

输出：

```
Stream Status: connected

 Address:		 58.34.232.182:22442
 StartFrom:		 2015-12-02T15:22:10.169Z
 BytesPerSecond:	 44402.2
 FramesPerSecond:
    Audio:	 28.2
    Video:	 13.8
    Data:	 0
```