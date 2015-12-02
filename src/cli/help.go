package cli

import "fmt"

const VERISON = "1.0.0"

var helpInfo = map[string]map[string]string{
	"hub": map[string]string{
		"reg":           "qlive hub reg -ak <AccessKey> -sk <SecretKey> -hub <HubName>",
		"stat":          "qlive hub info",
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
		"hls-play":  "qlive stream hls-play <StreamId> -s <Start> -e <End>",
		"status":    "qlive stream status <StreamId>",
		"saveas":    "qlive stream saveas -n <Name> -f <Format> -s <Start> -e <End> -c <NotifyUrl>",
		"snapshot":  "qlive stream snapshot -n <Name> -f <Format> -t <Time> -c <NotifyUrl>",
	},
}

func Version() {
	fmt.Println("Qlive", VERISON)
}

func Help() {
	fmt.Println("QLive", VERISON)
	fmt.Println()
	fmt.Println("Commands for hub:")
	for cmd, help := range helpInfo["hub"] {
		fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, help))
	}

	fmt.Println()
	fmt.Println("Commands for stream:")
	for cmd, help := range helpInfo["stream"] {
		fmt.Println(fmt.Sprintf("%15s\t\t%s", cmd, help))
	}
	fmt.Println()
}

func CmdHelp(cmd string, subCmd string) {
	fmt.Println("Usage:", helpInfo[cmd][subCmd])
}
