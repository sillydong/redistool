package main
import (
	"os"
	"fmt"
	"github.com/sillydong/cli"
	"gopkg.in/redis.v3"
)

const APPVERSION = "20151010"
var app *cli.App

func main() {
	app = cli.NewApp()
	app.Name = "RedisTool"
	app.Author = "Chen.Zhidong"
	app.Copyright = "http://sillydong.com"
	app.Usage = "Help do something that redis-cli can not do"
	app.Version = APPVERSION
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "keys",
			Usage:  "show all keys matching pattern",
			Action: keys,
			Flags:[]cli.Flag{
				cli.StringFlag{"h", "127.0.0.1", "Server hostname (default: 127.0.0.1).", ""},
				cli.IntFlag{"p", 6379, "Server port (default: 6379).", ""},
				cli.StringFlag{"a", "", "Password to use when connecting to the server.", ""},
				cli.IntFlag{"n", 0, "Database number.", ""},
				cli.StringFlag{"r", "<pattern>", `find keys matching pattern
	*	matches all
	h?llo	matches hello, hallo and hxllo
	h*llo	matches hllo and heeeello
	h[ae]llo	matches hello and hallo, but not hillo
	h[^e]llo	matches hallo, hbllo, ... but not hello
	h[a-b]llo	matches hallo and hbllo`, ""},
				cli.BoolFlag{"d","delete keys matching pattern",""},
			},
		},
	}
	app.Run(os.Args)
}

func keys(ctx *cli.Context){
	host:=ctx.String("h")
	password:=ctx.String("a")
	regex := ctx.String("r")
	port := ctx.Int("p")
	database := int64(ctx.Int("n"))
	delete := ctx.Bool("d")
	
	if len(ctx.Args())>0 || regex == "<pattern>"{
		fmt.Println("Incorrect Usage.")
		fmt.Println("")
		cli.ShowCommandHelp(ctx,"mdel")
	}else{
		client:=redis.NewClient(&redis.Options{
			Addr:fmt.Sprintf("%s:%d",host,port),
			Password:password,
			DB:database,
		})
		keys,err:=client.Keys(regex).Result()
		if err != nil {
			fmt.Printf("%v\n",err)
		}else{
			if len(keys)==0{
				fmt.Printf("no keys match pattern \"%s\"\n",regex)
			}else{
				if delete{
					count, err := client.Del(keys...).Result()
					if err != nil {
						fmt.Printf("%v\n", err)
					}else {
						fmt.Printf("%d keys match pattern \"%s\" has been deleted\n", count, regex)
					}
				}else{
					fmt.Println("matched keys:")
					for _,key := range keys{
						fmt.Println("\t"+key)
					}
				}
			}
		}
	}
}
