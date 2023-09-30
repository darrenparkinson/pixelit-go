# PixelIt CLI

[![Status](https://img.shields.io/badge/status-wip-yellow)](https://github.com/darrenparkinson/pixelit-go) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/darrenparkinson/pixelit-go) ![GitHub](https://img.shields.io/github/license/darrenparkinson/pixelit-go?color=brightgreen) [![GoDoc](https://pkg.go.dev/badge/darrenparkinson/pixelit-go)](https://pkg.go.dev/github.com/darrenparkinson/pixelit-go) [![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/pixelit-go)](https://goreportcard.com/report/github.com/darrenparkinson/pixelit-go)


> *PixelIt is an ESP8266/ESP32 and WS2812B LED Matrix based PixelArt display*

You can find details for the project on their [documentation page](https://pixelit-project.github.io/) or their [github repository](https://github.com/pixelit-project)

The project supports an [HTTP REST-based interface](https://pixelit-project.github.io/api.html); an [MQTT interface](https://pixelit-project.github.io/api.html); and has a [node-red](https://pixelit-project.github.io/nodered.html) node for interactivity.

This repository contains two additional components for use with that [PixelIt Matrix Display](https://pixelit-project.github.io/):

* [Command Line Interface (CLI)](#command-line-interface-cli) for interacting with the PixelIt API;
* [Go Library](#go-library) used by this CLI which can also be used in your own projects;

> Note: It is still a work in progress.

## Command Line Interface (CLI)

![cli](pixelit.gif)

### Features

The CLI currently provides the ability to send the following supported screens to your PixelIt display:

* Clock - currently only default options (more options to come)
* Text - currently text only (more options to come)
* Bitmap - currently only default bitmap (more options to come)
* Social - see below

The social command currently supports the following sub-commands in order to display follower counts for the supplied username:

* Instagram - send follow count for given instagram username
* TikTok - send follow count for given tiktok username

### Installation

You can get started with the CLI by downloading an executable for your platform from the [releases page](https://github.com/darrenparkinson/pixelit-go/releases). 

Alternatively if you have Go installed, you can use the following command to get started:

```sh
$ go install github.com/darrenparkinson/pixelit-go/cmd/pixelit-cli
```

### Usage

You can get help with the CLI by using the `--help` option:

```sh
$ pixelit-cli --help

PixelIt CLI is used to interact with the PixelIt API.

Usage:
  pixelit-cli [command]

Available Commands:
  bitmap      Send a bitmap to PixelIt by ID or as data
  clock       Send a clock to PixelIt
  help        Help about any command
  serve       Read configuration and perform timed actions
  social      Send social network follower counts to PixelIt
  text        Send text to PixelIt

Flags:
      --config string   config file (default $HOME/.pixelit.json)
  -h, --help            help for pixelit-cli
      --host string     host name or ip of your PixelIt (default "192.168.86.25")
  -v, --verbose         provide verbose output
      --version         version for pixelit-cli

Use "pixelit-cli [command] --help" for more information about a command.
```

You are then able to run commands such as the following:

* Send a standard clock:
```sh
$ ./pixelit-cli clock
```
* Send bitmap by ID:
```sh
$ ./pixelit-cli bitmap --id 878
```
* Send text:
```sh
$ ./pixelit-cli text "PixelIt is Awesome!"
```
* Send instagram follower count for apple:
```sh
$ ./pixelit-cli social instagram --username apple 
```
* Send tiktok follower count for apple
```sh
$ ./pixelit-cli social tiktok --username apple 
```

## Go Library

The Go library currently provides the ability to send any screen type to your PixelIt, along with some helper functions that support the CLI.

### Usage

You can install the library in the usual way as follows:

```sh
$ go get github.com/darrenparkinson/pixelit-go
```

In your code, import the library:

```go
import "github.com/darrenparkinson/pixelit-go"
```

You can then construct a new PixelIt client, and use the various methods on the client to access different parts of the API.  For example:

```go
pxc, _ := pixelit.NewClient("192.168.86.25", nil)
pxc.SendInstagramFollowers(username)
```

You can also send raw "Screens" to your PixelIt which should cover all the options you need, e.g.:

```go
pxc, _ := pixelit.NewClient("192.168.86.25", nil)
s := pixelit.Screen{
		Clock: &pixelit.Clock{
			Show:         pixelit.Bool(true),
			SwitchAktiv:  pixelit.Bool(true),
			WithSeconds:  pixelit.Bool(false),
			SwitchSec:    pixelit.Int(7),
			DrawWeekDays: pixelit.Bool(true),
			Color: &pixelit.Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	pxc.SendScreen(&s)
```

See the go docs to see all available fields.

You will notice the use of some helper functions, such as `pixelit.Bool` and `pixelit.Int`.   These allow distinguishing between unset fields and those set to a zero value. This can cause challenges when retrieving fields since you may encounter a panic if you access a nil pointer. Clearly this isn't a very nice user experience, so where appropriate, "getter" accessor functions are generated automatically for structs with pointer fields to enable you to safely retrieve values:

```go
clock := s.GetClock()
```


## TODO

* [ ] Add configuration options for clock command
* [ ] Add configuration options for bitmap command
* [ ] Add configuration options for text command
* [ ] Add options to send text with bitmap

 