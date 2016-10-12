# 七牛直播命令行工具

## 简介
本项目是七牛直播业务的命令行工具，可以用在直播业务的开发和调试过程中。该工具基于[PILI v1](https://github.com/pili-engineering/pili-sdk-go) 的服务端SDK。

## 下载
可以下载已编译版本，直接使用。

|版本|支持平台|链接|
|---|-----|-----|
|QLive v1.0.3|Windows，Linux，Mac|[下载](http://devtools.qiniu.com/qlive-v1.0.3.zip)|

## 使用方法

**开通直播服务**

1. 首先需要向七牛申请开通直播服务，拥有直播的账号和`Hub`信息。
2. 然后根据七牛的技术支持的提示，配置直播服务相关的域名。

**准备工作**

1. 本例使用`qlive`做讲解，请重命名下载后的工具名称为`qlive`或者`qlive.exe`，或者使用原始名称；
2. 如果是`Linux`或者`Mac`系统，请先使用`chmod +x qlive`来为工具添加可执行权限；
3. 本工具是命令行工具，在`Windows`系统下，请不要直接双击打开，而是从命令行窗口运行；
4. 在`Windows 7`及其以上版本，可以在工具所在目录使用`Shift+右键`快速打开命令行窗口。
5. 可以将命令所在的目录添加到系统的环境变量`PATH`中，这样可以从任何路径运行命令。

**命令概述**

1. 可以使用`qlive -h`或`qlive.exe -h`来查看所有命令的帮助信息；
2. 可以使用`qlive hub 子命令`或者`qlive.exe hub 子命令`来查看子命令的帮助信息；
3. 该工具命令分为`hub`组命令和`stream`组命令，其他都是组下面的子命令。


**注意事项**

为了能够支持多hub的情况，该工具设计的时候将hub的信息写入到工具当前执行时所在的目录。这样我们就可以规划一些目录，专门做为工具的运行目录。
比如如下的目录`hub1`和`hub2`专门做为工具的执行目录，其中每个目录下面的`hub.json`存储了hub的配置信息。

```
➜  hubs  pwd
/Users/jemy/Temp/hubs

➜  hubs  tree -a
.
├── hub1
│   └── .qlive
│       └── hub.json
└── hub2
    └── .qlive
        └── hub.json

4 directories, 2 files
```

## 命令列表

**所有的支持的命令列表**

```
QLive 1.0.1

Commands for hub:
            reg		qlive hub reg -ak <AccessKey> -sk <SecretKey> -hub <HubName>
           info		qlive hub info
  create-stream		qlive hub create-stream -t <Title> -pbk <PublishKey> -pbs <PublishSecurity>
     get-stream		qlive hub get-stream <StreamId>
    list-stream		qlive hub list-stream -s <Status> -l <Limit> -p <Prefix>

Commands for stream:
         update		qlive stream update <StreamId> -pbk <PublishKey> -pbs <PublishSecurity>
         delete		qlive stream delete <StreamId>
        disable		qlive stream disable <StreamId>
         enable		qlive stream enable <StreamId>
       rtmp-pub		qlive stream rtmp-pub <StreamId>
      rtmp-live		qlive stream rtmp-live <StreamId>
       hls-live		qlive stream hls-live <StreamId>
       flv-live		qlive stream flv-live <StreamId>
         status		qlive stream status <StreamId>
       hls-play		qlive stream hls-play <StreamId> -s <Start> -e <End>
         saveas		qlive stream saveas <StreamId> -n <Name> -f <Format> -s <Start> -e <End> -p <Pipeline> -c <NotifyUrl>
       snapshot		qlive stream snapshot <StreamId> -n <Name> -f <Format> -t <Time> -c <NotifyUrl>
```


**Hub组命令**

|命令|描述|详细|
|----|-----|-----|
|reg|设置工具的`AccessKey`，`SecretKey`和`Hub`信息|[详细](docs/qlive_cmd_hub_reg.md)|
|info|查看工具当前设置的`AccessKey`，`SecretKey`和`Hub`信息|[详细](docs/qlive_cmd_hub_info.md)|
|create-stream|创建一个新的直播流，可以选择性设置`PublishKey`和`PublishSecurity`信息|[详细](docs/qlive_cmd_hub_create_stream.md)|
|get-stream|根据流Id获取一个流的基本信息|[详细](docs/qlive_cmd_hub_get_stream.md)|
|list-stream|根据列举选项，获取当前`Hub`中的流列表|[详细](docs/qlive_cmd_hub_list_stream.md)|

**Stream组命令**

|命令|描述|详细|
|-----|-----|-----|
|update|更新流的`PublishKey`或者`PublishSecurity`信息|[详细](docs/qlive_cmd_stream_update.md)|
|delete|根据流Id删除一个流|[详细](docs/qlive_cmd_stream_delete.md)|
|disable|根据流Id禁用一个流|[详细](docs/qlive_cmd_stream_disable.md)|
|enable|根据流Id启用一个流|[详细](docs/qlive_cmd_stream_enable.md)|
|rtmp-pub|根据流Id获取该流的RTMP推流地址|[详细](docs/qlive_cmd_stream_rtmp_pub.md)|
|rtmp-live|根据流Id获取该流的RTMP直播播放地址|[详细](docs/qlive_cmd_stream_rtmp_live.md)|
|hls-live|根据流Id获取该流的HLS直播播放地址|[详细](docs/qlive_cmd_stream_hls_live.md)|
|flv-live|根据流Id获取该流的FLV直播播放地址|[详细](docs/qlive_cmd_stream_flv_live.md)|
|status|根据流Id获取该流的当前状态|[详细](docs/qlive_cmd_stream_status.md)|
|hls-play|根据流Id和起始，结束时间获取直播流的HLS回放地址|[详细](docs/qlive_cmd_stream_hls_play.md)|
|saveas|根据流Id和起始，结束时间将直播流转存为视频格式，并返回视频地址|[详细](docs/qlive_cmd_stream_saveas.md)|
|snapshot|根据流Id和截取图片时间从直播流中截取一张图片，并返回截图地址|[详细](docs/qlive_cmd_stream_snapshot.md)|