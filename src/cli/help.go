package cli
import "fmt"

const VERISON = "1.0.0"

var helpInfo = map[string]map[string]string{
	"hub":map[string]string{
		"reg":"qlive hub reg -ak <AccessKey> -sk <SecretKey> -hub <HubName>",
		"stat":"qlive hub stat",
		"create-stream":"qlive hub create-stream -t <Title> -pbk <PublishKey> -pbs <PublishSecurity>",
		"get-stream":"qlive hub get-stream <StreamId>",
		"list-stream":"qlive hub list-stream -s <Status> -l <Limit> -p <Prefix>",
	},
}

func Version() {
	fmt.Println("qlive", VERISON)
}

func Help() {

}

func CmdHelp(cmd string, subCmd string) {
	fmt.Println("Usage:", helpInfo[cmd][subCmd])
}