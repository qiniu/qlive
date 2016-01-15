package cli

import "fmt"

const VERISON = "1.0.2"

var hubCmdOrder = []string{"reg", "info", "create-stream", "get-stream", "list-stream"}
var streamCmdOrder = []string{"update", "delete", "disable", "enable", "rtmp-pub", "rtmp-live",
	"hls-live", "flv-live", "status", "hls-play", "saveas", "snapshot"}

var helpInfo = map[string]map[string]string{
	"hub": map[string]string{
		"reg":           "qlive hub reg -ak <AccessKey> -sk <SecretKey> -hub <HubName>",
		"info":          "qlive hub info",
		"create-stream": "qlive hub create-stream -t <Title> -pbk <PublishKey> -pbs <PublishSecurity>",
		"get-stream":    "qlive hub get-stream <StreamId>",
		"list-stream":   "qlive hub list-stream -s <Status> -l <Limit> -p <Prefix>",
	},
	"stream": map[string]string{
		"update":    "qlive stream update <StreamId> -pbk <PublishKey> -pbs <PublishSecurity>",
		"delete":    "qlive stream delete <StreamId>",
		"disable":   "qlive stream disable <StreamId>",
		"enable":    "qlive stream enable <StreamId>",
		"rtmp-pub":  "qlive stream rtmp-pub <StreamId>",
		"rtmp-live": "qlive stream rtmp-live <StreamId>",
		"hls-live":  "qlive stream hls-live <StreamId>",
		"flv-live":  "qlive stream flv-live <StreamId>",
		"status":    "qlive stream status <StreamId>",
		"hls-play":  "qlive stream hls-play <StreamId> -s <Start> -e <End>",
		"saveas":    "qlive stream saveas <StreamId> -n <Name> -f <Format> -s <Start> -e <End> -c <NotifyUrl>",
		"snapshot":  "qlive stream snapshot <StreamId> -n <Name> -f <Format> -t <Time> -c <NotifyUrl>",
	},
}

func Version() {
	fmt.Println("Qlive", VERISON)
}

func Help() {
	fmt.Println("QLive", VERISON)
	fmt.Println()
	fmt.Println("Commands for hub:")
	for _, cmd := range hubCmdOrder {
		fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, helpInfo["hub"][cmd]))
	}

	fmt.Println()
	fmt.Println("Commands for stream:")
	for _, cmd := range streamCmdOrder {
		fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, helpInfo["stream"][cmd]))
	}
	fmt.Println()
}

func CmdHelp(cmd string, subCmd string) {
	fmt.Println("Usage:", helpInfo[cmd][subCmd])
}
