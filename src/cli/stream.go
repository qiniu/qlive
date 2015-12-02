package cli

import (
	"fmt"
	"github.com/jemygraw/pili-sdk-go/pili"
	"flag"
	"os"
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
		fmt.Println("Unknown cmd ", subCmd, "for stream")
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

	fmt.Println("Updated Stream")
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

}

func GetRtmpLiveAddress(cmd string, subCmd string) {

}

func GetHlsLiveAddress(cmd string, subCmd string) {

}

func GetFlvLiveAddress(cmd string, subCmd string) {

}

func GetHlsPlayAddress(cmd string, subCmd string) {

}

func GetStreamStatus(cmd string, subCmd string) {

}

func SaveStreamAsVideo(cmd string, subCmd string) {

}

func TakeStreamSnapshot(cmd string, subCmd string) {

}
