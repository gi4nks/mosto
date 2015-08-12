package main

import (
	"github.com/codegangsta/cli"
	"github.com/gi4nks/quant"
	"os"
)

var tracer = quant.NewTrace("mosto")

func main() {
	tracer.Light()
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
		{
			Name:    "functions",
			Aliases: []string{"f"},
			Usage:   "functions",
			Action:  functionsSample,
		},
		{
			Name:    "tideot",
			Aliases: []string{"t"},
			Usage:   "tideot",
			Action:  tideotSample,
		},
		{
			Name:    "randoms",
			Aliases: []string{"r"},
			Usage:   "randoms",
			Action:  randomsSample,
		},
		{
			Name:    "loggers",
			Aliases: []string{"l"},
			Usage:   "loggers",
			Action:  loggersSample,
		},
		{
			Name:    "compressors",
			Aliases: []string{"c"},
			Usage:   "compressors",
			Action:  compressorSample,
		},
		{
			Name:    "docker",
			Aliases: []string{"d"},
			Usage:   "docker",
			Action:  dockerSample,
		},
	}

	app.Run(os.Args)
}

func pointersSample(ctx *cli.Context) {
	tracer.Notice("pointers ...")

	var sample = Sample{}

	tracer.News("Sample{} --> " + sample.String())

	sample.Stateless()
	tracer.News("sample.Stateless() --> " + sample.String())

	sample.Stateful()
	tracer.News("sample.Stateful() --> " + sample.String())

}

func interfacesSample(ctx *cli.Context) {
	tracer.Notice("interfaces ...")

	var animal Animal

	tracer.News("Declaring interface variable (var animal Animal)")

	var cat = Cat{}
	animal = cat

	tracer.News("Cat (animal.Greet()) --> " + animal.Greet())

	var dog = Dog{}
	animal = dog

	tracer.News("Dog (animal.Greet()) --> " + animal.Greet())

	var bird = Bird{}
	animal = bird

	tracer.News("Bird (animal.Greet()) --> " + animal.Greet())
}

func errorsSample(ctx *cli.Context) {
	tracer.Notice("errors ...")

	i1 := -1
	i2 := 1

	if result, err := GenerateResult(i1); err == nil {
		tracer.News("-1 --> This is a good result: " + result.String())
	} else {
		tracer.Warning("-1 --> This is a bad result: " + result.String() + " - " + err.Error())
	}

	if result, err := GenerateResult(i2); err == nil {
		tracer.News("1 --> This is a good result: " + result.String())
	} else {
		tracer.Warning("1 --> This is a bad result: " + result.String() + " - " + err.Error())
	}
}

func functionsSample(ctx *cli.Context) {
	tracer.Notice("functions ...")

	Functions()
}

func tideotSample(ctx *cli.Context) {
	tracer.Notice("tideot ...")

	//EmbeddedExample()
	PersistentSample1()
}

func randomsSample(ctx *cli.Context) {
	tracer.Notice("randoms ...")

	RandomsSample1()
	RandomsSample2()
	RandomsSample3()
	RandomsSample4()

}

func loggersSample(ctx *cli.Context) {
	tracer.Notice("loggers ...")

	LoggerSample1()
	LoggerSample2()
	LoggerSample3()

}

func loggersSample(ctx *cli.Context) {
	tracer.Notice("loggers ...")

	LoggerSample()
}

func compressorSample(ctx *cli.Context) {
	tracer.Notice("compressors ...")

	var b = Compress("hello world")

	tracer.News(b)

	var b1 = Uncompress(b)

	tracer.News(b1)
}

func dockerSample(ctx *cli.Context) {
	tracer.Notice("docker ...")

	DockerClientTest()
}
