# 简介

根据列举选项，获取当前`Hub`中的流列表

# 命令

```
qlive hub list-stream -s <Status> -l <Limit> -p <Prefix>
```

# 参数

|名称|描述|可选
|--------|-------|
|Status|直播流的状态，可选参数，可选值为`connected`和`disconnected`|Y|
|Limit|限定返回的记录最大数量，可选参数，如果不设置，会返回所有的流|Y|
|Prefix|直播流Title的前缀，可以根据该前缀过滤流|Y|

# 示例

获取当前直播通道下的所有流

```
qlive hub list-stream
```

结果：

```
z1.jinxinxin.56580f1cfb16dfad7e0000a7	56580f1cfb16dfad7e1000a7	92b1538205bc7127	dynamic
z1.jinxinxin.56580d75eb6f92072a000084	56580d75eb6f92072a100084	1449a422e3c3c317	dynamic
z1.jinxinxin.5656d4cad409d2984a0003b5	5656d4cad409d2984a1003b5	289f7c52d527df30	dynamic
z1.jinxinxin.56566d31eb6f922e0f000052	56566d31eb6f922e0f100052	3b1a88b2dc56313e	dynamic
z1.jinxinxin.565662b2fb16df6fed000038	565662b2fb16df6fed100038	4c68edf2bdbbc81e	dynamic
z1.jinxinxin.5655424bd409d213a600010b	5655424bd409d213a610010b	c15bf1824ea3c136	dynamic
z1.jinxinxin.5653d240d409d27a2400005e	5653d240d409d27a2410005e	154de302d45c1853	dynamic
z1.jinxinxin.56531c13eb6f921fc5000c41	56531c13eb6f921fc5100c41	e958c0428e950d17	dynamic
z1.jinxinxin.5653194aeb6f921fc5000c3c	5653194aeb6f921fc5100c3c	edea5432d964c972	dynamic
```


# 备注

为了减轻输入子命令的痛苦，上面的`list-stream`也可以缩短为`ls`，比如如下的命令也是一样的效果：

```
qlive hub ls
```