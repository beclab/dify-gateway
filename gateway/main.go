package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"wzinc/database"
	"wzinc/dify"
	"wzinc/http_server"

	"syscall"

	"github.com/rs/zerolog"
	cli "gopkg.in/urfave/cli.v1"
	"wzinc/inotify"
)

var (
	OriginCommandHelpTemplate = `{{.Name}}{{if .Subcommands}} command{{end}}{{if .Flags}} [command options]{{end}} {{.ArgsUsage}}
{{if .Description}}{{.Description}}
{{end}}{{if .Subcommands}}
SUBCOMMANDS:
  {{range .Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}{{end}}{{if .Flags}}
OPTIONS:
{{range $.Flags}}   {{.}}
{{end}}
{{end}}`
)
var app *cli.App
var HTTPServer *http_server.Server

const DefaultPort = "6317"

func init() {
	app = cli.NewApp()
	app.Version = "v0.2.13"
	app.Commands = []cli.Command{
		commandStart,
	}

	cli.CommandHelpTemplate = OriginCommandHelpTemplate
}

var commandStart = cli.Command{
	Name:   "start",
	Usage:  "start loading contract gas fee",
	Flags:  []cli.Flag{},
	Action: Start,
}

func Start(ctx *cli.Context) {
	database.InitDB()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	inotify.WatchDir = os.Getenv("WATCH_DIR")
	fmt.Fprintln(os.Stderr, inotify.WatchDir)
	if inotify.WatchDir == "" {
		inotify.WatchDir = "/data"
	}
	dify.InitDifyV2()
	//inotify.WatchPath(inotify.WatchDir)
	//inotify.ListApp()

	// 初始化 PathToDatasetMap
	inotify.InitializePathToDatasetMap(inotify.WatchDir, dify.DatasetId)

	inotify.WatchPath(inotify.PathToDatasetMap)

	// 创建 HTTP 服务器
	HTTPServer = http_server.NewServer()

	// 启动服务器
	err := HTTPServer.Start(":6317")
	if err != nil {
		log.Fatal(err)
	}
	waitToExit()
}

func main() {
	// 启动定时执行逻辑的goroutine
	go dify.StartScheduledExecution()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	if !signal.Ignored(syscall.SIGHUP) {
		signal.Notify(sc, syscall.SIGHUP)
	}
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range sc {
			fmt.Printf("received exit signal:%v", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}
