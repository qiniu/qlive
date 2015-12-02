package cli

import (
	"flag"
	"fmt"
	"github.com/jemygraw/pili-sdk-go/pili"
	"os"
)

var liveHub = LiveHub{}
var hubSubCmdAlias = map[string]string{
	"cs": "create-stream",
	"gs": "get-stream",
	"ls": "list-stream",
}

var supportedHubSubCmds = map[string]func(string, string){
	"info":          GetHubInfo,
	"reg":           RegisterHub,
	"create-stream": CreateStream,
	"get-stream":    GetStream,
	"list-stream":   ListStream,
}

func Hub(subCmd string) {
	//add alias support
	if vCmd, vOk := hubSubCmdAlias[subCmd]; vOk {
		subCmd = vCmd
	}

	//parse and exec sub cmd
	if subCmdFunc, ok := supportedHubSubCmds[subCmd]; ok {
		subCmdFunc("hub", subCmd)
	} else {
		fmt.Println("Unknown cmd", fmt.Sprintf("`%s`", subCmd), "for hub")
	}
}

///////////////////////////////////////////////////////////////

func GetHubInfo(cmd string, subCmd string) {
	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	fmt.Println(liveHub.String())
}

func RegisterHub(cmd string, subCmd string) {
	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var ak string
	var sk string
	var hub string

	flagSet.StringVar(&ak, "ak", "", "access key")
	flagSet.StringVar(&sk, "sk", "", "secret key")
	flagSet.StringVar(&hub, "hub", "", "hub name")

	flagSet.Parse(os.Args[3:])

	if ak == "" || sk == "" || hub == "" {
		CmdHelp(cmd, subCmd)
		return
	}

	regErr := liveHub.Set(ak, sk, hub)
	if regErr != nil {
		fmt.Println(regErr)
	}
}

//create a new stream
func CreateStream(cmd string, subCmd string) {
	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var title string
	var publishKey string
	var publishSecurity string

	flagSet.StringVar(&title, "t", "", "title")
	flagSet.StringVar(&publishKey, "pbk", "", "publish key")
	flagSet.StringVar(&publishSecurity, "pbs", "", "publish security")

	flagSet.Parse(os.Args[3:])

	if publishSecurity != "" {
		if !(publishSecurity == "static" || publishSecurity == "dynamic") {
			fmt.Println("Publish Security can only be 'static' or 'dynamic'")
			return
		}
	}

	options := pili.OptionalArguments{
		Title:           title,
		PublishKey:      publishKey,
		PublishSecurity: publishSecurity,
	}

	stream, cErr := hub.CreateStream(options)
	if cErr != nil {
		fmt.Println("Create stream error,", cErr)
		return
	}

	jsonStr, jsonErr := stream.ToJSONString()
	if jsonErr != nil {
		fmt.Println("Create stream error,", jsonErr)
		return
	}

	fmt.Println(jsonStr)
}

//get stream info
func GetStream(cmd string, subCmd string) {
	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	streamId := os.Args[3]
	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)
	stream, sErr := hub.GetStream(streamId)
	if sErr != nil {
		fmt.Println("Get stream error,", sErr)
		return
	}

	jsonStr, jsonErr := stream.ToJSONString()
	if jsonErr != nil {
		fmt.Println("Get stream error,", jsonErr)
		return
	}

	fmt.Println(jsonStr)
}

//list streams
func ListStream(cmd string, subCmd string) {
	gErr := liveHub.Get()
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	cred := pili.NewCredentials(liveHub.AccessKey, liveHub.SecretKey)
	hub := pili.NewHub(cred, liveHub.Hub)

	flagSet := flag.NewFlagSet(subCmd, flag.ExitOnError)
	flagSet.Usage = func() {
		CmdHelp(cmd, subCmd)
	}

	var status string
	var marker string
	var limit int
	var prefix string

	flagSet.StringVar(&status, "s", "", "stream status")
	flagSet.StringVar(&marker, "m", "", "list offset marker")
	flagSet.IntVar(&limit, "l", 0, "limit count")
	flagSet.StringVar(&prefix, "t", "", "title prefix")

	flagSet.Parse(os.Args[3:])

	options := pili.OptionalArguments{
		Status: status,
		Marker: marker,
		Limit:  uint(limit),
		Title:  prefix,
	}

	maxRetries := 5

	fCnt := 0

	for {
		listResult, lErr := hub.ListStreams(options)
		if lErr != nil {
			fmt.Println("List stream error,", lErr)
			if maxRetries > 0 {
				fmt.Println("Retrying...")
				maxRetries -= 1
				continue
			} else {
				break
			}
		}

		for _, item := range listResult.Items {
			fmt.Println(fmt.Sprintf("%s\t%s\t%s\t%s", item.Id, item.Title, item.PublishKey, item.PublishSecurity))
		}

		fCnt += len(listResult.Items)

		if (limit > 0 && fCnt >= limit) || listResult.End {
			break
		}

		options.Marker = listResult.Marker
		//reset
		maxRetries = 5
	}
}
