package cli

import (
	"flag"
	"fmt"
	"github.com/jemygraw/pili-sdk-go/pili"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var supportedStreamSubCmds = map[string]func(string, string){
	"update":    UpdateStream,
	"delete":    DeleteStream,
	"disable":   DisableStream,
	"enable":    EnableStream,
	"rtmp-pub":  GetRtmpPubAddress,
	"rtmp-live": GetRtmpLiveAddress,
	"hls-live":  GetHlsLiveAddress,
	"flv-live":  GetFlvLiveAddress,
	"hls-play":  GetHlsPlayAddress,
	"status":    GetStreamStatus,
	"saveas":    SaveStreamAsVideo,
	"snapshot":  TakeStreamSnapshot,
}

func Stream(subCmd string) {
	//parse and exec sub cmd
	if subCmdFunc, ok := supportedStreamSubCmds[subCmd]; ok {
		subCmdFunc("stream", subCmd)
	} else {
		fmt.Println("Unknown cmd", fmt.Sprintf("`%s`", subCmd), "for stream")
	}
}

/////////////////////////////////////////////////////////

func UpdateStream(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var publishKey string
	var publishSecurity string

	flagSet.StringVar(&publishKey, "pbk", "", "publish key")
	flagSet.StringVar(&publishSecurity, "pbs", "", "publish security")

	flagSet.Parse(os.Args[4:])

	if publishKey == "" && publishSecurity == "" {
		CmdHelp(cmd, subCmd)
		return
	}

	if publishSecurity != "" {
		if !(publishSecurity == "static" || publishSecurity == "dynamic") {
			fmt.Println("Publish Security can only be 'static' or 'dynamic'")
			return
		}
	}

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	stream.PublishKey = publishKey
	stream.PublishSecurity = publishSecurity

	nStream, uErr := stream.Update()
	if uErr != nil {
		fmt.Println("Update stream error,", uErr)
		return
	}

	fmt.Println("Updated Stream:")
	fmt.Println()
	fmt.Println(" Stream Id:\t\t", nStream.Id)
	fmt.Println(" Stream Title:\t\t", nStream.Title)
	fmt.Println(" Publish Key:\t\t", nStream.PublishKey)
	fmt.Println(" Publish Security:\t", nStream.PublishSecurity)
}

func DeleteStream(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	_, dErr := stream.Delete()
	if dErr != nil {
		fmt.Println("Delete stream error,", dErr)
		return
	}
	fmt.Println("Done!")
}

func DisableStream(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	_, dErr := stream.Disable()
	if dErr != nil {
		fmt.Println("Disable stream error,", dErr)
		return
	}
	fmt.Println("Done!")
}

func EnableStream(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	_, dErr := stream.Enable()
	if dErr != nil {
		fmt.Println("Enable stream error,", dErr)
		return
	}
	fmt.Println("Done!")
}

func GetRtmpPubAddress(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	rmptPubUrl := stream.RtmpPublishUrl()
	fmt.Println(rmptPubUrl)
}

func GetRtmpLiveAddress(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	rtmpLiveUrls, gErr := stream.RtmpLiveUrls()
	if gErr != nil {
		fmt.Println("Get rtmp live urls error,", gErr)
		return
	}

	for tag, liveUrl := range rtmpLiveUrls {
		fmt.Println(tag, liveUrl)
	}
}

func GetHlsLiveAddress(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	hlsLiveUrls, gErr := stream.HlsLiveUrls()
	if gErr != nil {
		fmt.Println("Get hls live urls error,", gErr)
		return
	}

	for tag, liveUrl := range hlsLiveUrls {
		fmt.Println(tag, liveUrl)
	}
}

func GetFlvLiveAddress(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	flvLiveUrls, gErr := stream.HttpFlvLiveUrls()
	if gErr != nil {
		fmt.Println("Get flv live urls error,", gErr)
		return
	}

	for tag, liveUrl := range flvLiveUrls {
		fmt.Println(tag, liveUrl)
	}
}

//@param start 20151202102000
//@param end
func GetHlsPlayAddress(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var startTime string
	var endTime string

	flagSet.StringVar(&startTime, "s", "", "start time")
	flagSet.StringVar(&endTime, "e", "", "end time")

	flagSet.Parse(os.Args[4:])

	if startTime == "" || endTime == "" {
		CmdHelp(cmd, subCmd)
		return
	}

	var start int64
	var end int64
	var pErr error

	if start, pErr = parseTime(startTime); pErr != nil {
		CmdHelp(cmd, subCmd)
		return
	}

	if end, pErr = parseTime(endTime); pErr != nil {
		CmdHelp(cmd, subCmd)
		return
	}

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	hlsPlayUrls, gErr := stream.HlsPlaybackUrls(start, end)
	if gErr != nil {
		fmt.Println("Get hls playback urls error,", gErr)
		return
	}

	for tag, playUrl := range hlsPlayUrls {
		fmt.Println(tag, playUrl)
	}

}

func parseTime(str string) (ts int64, err error) {
	if len(str) == 14 {
		//parse as 20151010000000
		yearStr := str[:4]
		monthStr := str[4:6]
		dayStr := str[6:8]
		hourStr := str[8:10]
		minuteStr := str[10:12]
		secondStr := str[12:14]

		var year int64
		var month int64
		var day int64
		var hour int64
		var minute int64
		var second int64
		var pErr error

		if year, pErr = strconv.ParseInt(yearStr, 10, 64); pErr != nil {
			err = pErr
			return
		}
		if month, pErr = strconv.ParseInt(monthStr, 10, 64); pErr != nil {
			err = pErr
			return
		}
		if day, pErr = strconv.ParseInt(dayStr, 10, 64); pErr != nil {
			err = pErr
			return
		}
		if hour, pErr = strconv.ParseInt(hourStr, 10, 64); pErr != nil {
			err = pErr
			return
		}
		if minute, pErr = strconv.ParseInt(minuteStr, 10, 64); pErr != nil {
			err = pErr
			return
		}
		if second, pErr = strconv.ParseInt(secondStr, 10, 64); pErr != nil {
			err = pErr
			return
		}

		loc, _ := time.LoadLocation("Asia/Shanghai")
		date := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
		ts = date.Unix()
	} else {
		ts, err = strconv.ParseInt(str, 10, 64)
	}

	return
}

func GetStreamStatus(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	platform := runtime.GOOS
	for {

		status, gErr := stream.Status()
		if gErr != nil {
			fmt.Println("Get stream status error,", gErr)
			return
		}

		var outputStatus string

		if platform == "windows" {
			outputStatus += fmt.Sprintf("%s", strings.ToUpper(status.Status))
		} else {
			if status.Status == "connected" {
				outputStatus += fmt.Sprintf("\033[32m%s\033[37m", strings.ToUpper(status.Status))
			} else {
				outputStatus += fmt.Sprintf("\033[31m%s\033[37m", strings.ToUpper(status.Status))
			}
		}

		if status.Status == "connected" {

			if platform == "windows" {
				outputStatus += fmt.Sprintf(" %s", status.StartFrom)
				outputStatus += fmt.Sprintf(" BPS: %.2f", status.BytesPerSecond)
				outputStatus += fmt.Sprintf(" FPSA: %.2f FPSV: %.2f FPSD: %.2f", status.FramesPerSecond.Audio,
					status.FramesPerSecond.Video, status.FramesPerSecond.Data)
			} else {
				outputStatus += fmt.Sprintf(" \033[33m%s\033[37m", status.StartFrom)
				outputStatus += fmt.Sprintf(" BPS: \033[32m%.2f\033[37m", status.BytesPerSecond)
				outputStatus += fmt.Sprintf(" FPSA: \033[32m%.2f\033[37m FPSV: \033[32m%.2f\033[37m FPSD: \033[32m%.2f\033[37m",
					status.FramesPerSecond.Audio, status.FramesPerSecond.Video, status.FramesPerSecond.Data)
			}
		}

		fmt.Print("\033[2K\r")
		fmt.Printf("%s", outputStatus)

		<-time.After(time.Millisecond * 500)
	}
}

func SaveStreamAsVideo(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var name string
	var format string
	var startTime string
	var endTime string
	var notifyUrl string

	flagSet.StringVar(&name, "n", "", "name")
	flagSet.StringVar(&format, "f", "", "format")
	flagSet.StringVar(&startTime, "s", "", "start time")
	flagSet.StringVar(&endTime, "e", "", "end time")
	flagSet.StringVar(&notifyUrl, "c", "", "notify url")

	flagSet.Parse(os.Args[4:])

	if name == "" || format == "" || startTime == "" || endTime == "" {
		CmdHelp(cmd, subCmd)
		return
	}

	var start int64
	var end int64
	var pErr error

	if start, pErr = parseTime(startTime); pErr != nil {
		CmdHelp(cmd, subCmd)
		return
	}

	if end, pErr = parseTime(endTime); pErr != nil {
		CmdHelp(cmd, subCmd)
		return
	}

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	options := pili.OptionalArguments{
		NotifyUrl: notifyUrl,
	}

	ret, sErr := stream.SaveAs(name, format, start, end, options)
	if sErr != nil {
		fmt.Println("Save stream as video error,", sErr)
		return
	}

	fmt.Println("M3u8Url:\t", ret.Url)
	fmt.Println("TargetUrl:\t", ret.TargetUrl)
	fmt.Println()
	fmt.Println(fmt.Sprintf("See http://api.qiniu.com/status/get/prefop?id=%s", ret.PersistentId))
}

func TakeStreamSnapshot(cmd string, subCmd string) {
	if len(os.Args) < 4 {
		CmdHelp(cmd, subCmd)
		return
	}

	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	streamId := os.Args[3]
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var name string
	var format string
	var snapshotTime string
	var notifyUrl string

	flagSet.StringVar(&name, "n", "", "name")
	flagSet.StringVar(&format, "f", "", "format")
	flagSet.StringVar(&snapshotTime, "t", "", "snapshot time")
	flagSet.StringVar(&notifyUrl, "c", "", "notify url")

	flagSet.Parse(os.Args[4:])

	if name == "" || format == "" || snapshotTime == "" {
		CmdHelp(cmd, subCmd)
		return
	}

	var shotTime int64
	var pErr error

	if shotTime, pErr = parseTime(snapshotTime); pErr != nil {
		CmdHelp(cmd, subCmd)
		return
	}

	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		fmt.Println("Get stream error,", gErr)
		return
	}

	options := pili.OptionalArguments{
		Time:      shotTime,
		NotifyUrl: notifyUrl,
	}

	ret, sErr := stream.Snapshot(name, format, options)
	if sErr != nil {
		fmt.Println("Snapshot stream error,", sErr)
		return
	}

	fmt.Println("TargetUrl:", ret.TargetUrl)
	fmt.Println()
	fmt.Println(fmt.Sprintf("See http://api.qiniu.com/status/get/prefop?id=%s", ret.PersistentId))
}
