# 七牛直播命令行工具

## 简介
本项目是七牛直播业务的命令行工具，可以用在直播业务的开发和调试过程中。

## 下载
可以下载已编译版本，直接使用。

|版本|平台|链接|
|---|-----|-----|
|QLive v1.0.0|Windows，Linux，Mac|[下载]()|

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

## 命令详解

**所有的支持的命令列表**

```
QLive 1.0.0

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
         saveas		qlive stream saveas -n <Name> -f <Format> -s <Start> -e <End> -c <NotifyUrl>
       snapshot		qlive stream snapshot -n <Name> -f <Format> -t <Time> -c <NotifyUrl>
```

