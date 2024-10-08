package consts

import (
	"os"
	"regexp"
)

const (
	BotApiLink  = "https://core.telegram.org/bots/api"
	PackageName = "github.com/GoBotApiOfficial/gobotapi"
	AGMessage   = "// Code AutoGenerated; DO NOT EDIT."
)

var OutputFolder string
var FolderPermission os.FileMode
var (
	CommonFields = []string{
		"from",
		"chat",
		"message",
		"date",
	}
	GenericInputRgx  = regexp.MustCompile("^Input.*Media$")
	MediaInputRgx    = regexp.MustCompile("^Input.*Media\\w+")
	InputTypeNameRgx = regexp.MustCompile("^(Input.*Media|)(.*?)$")
	BasicTypesRgx    = regexp.MustCompile(`^(int|string|bool|float64|\[])`)
)
