package main

import (
	"github.com/codegangsta/cli"
	"github.com/op/go-logging"
	"os"
)

var logger = logging.MustGetLogger("runTool")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func main() {
	// Configuring logger
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	// Set the backends to be used.
	logging.SetBackend(backendFormatter)
	// -------------------

	app := cli.NewApp()
	app.Name = "mosto"
	app.Usage = "Mosto is a sample application to sho Golang functionalities"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "pointers",
			Aliases: []string{"p"},
			Usage:   "pointers",
			Action:  pointersSample,
		},
		{
			Name:    "interfaces",
			Aliases: []string{"i"},
			Usage:   "interfaces",
			Action:  interfacesSample,
		},
		{
			Name:    "errors",
			Aliases: []string{"e"},
			Usage:   "errors",
			Action:  errorsSample,
		},
	}

	app.Run(os.Args)
}

func pointersSample(ctx *cli.Context) {
	logger.Notice("pointers ...")

	var sample = Sample{}

	logger.Info("Sample{} --> " + sample.String())

	sample.Stateless()
	logger.Info("sample.Stateless() --> " + sample.String())

	sample.Stateful()
	logger.Info("sample.Stateful() --> " + sample.String())

}

func interfacesSample(ctx *cli.Context) {
	logger.Notice("interfaces ...")

	var animal Animal

	logger.Info("Declaring interface variable (var animal Animal)")

	var cat = Cat{}
	animal = cat

	logger.Info("Cat (animal.Greet()) --> " + animal.Greet())

	var dog = Dog{}
	animal = dog

	logger.Info("Dog (animal.Greet()) --> " + animal.Greet())

	var bird = Bird{}
	animal = bird

	logger.Info("Bird (animal.Greet()) --> " + animal.Greet())
}

func errorsSample(ctx *cli.Context) {
	logger.Notice("errors ...")

	i1 := -1
	i2 := 1

	if result, err := GenerateResult(i1); err == nil {
		logger.Info("-1 --> This is a good result: " + result.String())
	} else {
		logger.Warning("-1 --> This is a bad result: " + result.String() + " - " + err.Error())
	}

	if result, err := GenerateResult(i2); err == nil {
		logger.Info("1 --> This is a good result: " + result.String())
	} else {
		logger.Warning("1 --> This is a bad result: " + result.String() + " - " + err.Error())
	}
}
