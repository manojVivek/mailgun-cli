package main

import (
	"github.com/alecthomas/kong"

	"github.com/manojVivek/mailgun-cli/pkg/mailgun"
)

var CLI struct {
	mailgun.Globals

	Template mailgun.TemplateCmd `cmd:"" help:"Manage templates."`
}

func main() {
	ctx := kong.Parse(&CLI, kong.Name("mailgun"), kong.Description("Mailgun CLI"), kong.UsageOnError())
	ctx.Run(&CLI.Globals)
}
