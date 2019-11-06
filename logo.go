package kara

import (
	"strings"
)

func GetLogo(httpPort string, grpcPort string) string {
	var tmpl = `
_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
 _   _
| |/ /              
| ' / __ _ _ __ __ _         ${repo}
|  < / _  |  __/ _  |        ${version}
| . \ (_| | | | (_| |        ${http_port}
|_|\_\__,_|_|  \__,_|        ${grpc_port}
_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
`
	var logo = strings.Replace(tmpl, "${repo}", "Repo: https://github.com/fwhezfwhez/kara", -1)
	logo = strings.Replace(logo, "${version}", "Version: "+Version, -1)
	logo = strings.Replace(logo, "${http_port}", "Http: "+httpPort, -1)
	logo = strings.Replace(logo, "${grpc_port}", "Grpc: "+grpcPort, -1)
	return logo
}
