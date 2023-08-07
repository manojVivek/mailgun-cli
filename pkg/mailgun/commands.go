package mailgun

import (
	"context"
)

type Globals struct {
	ApiKey string `short:"k" help:"Mailgun API key."`
}

type TemplateCopy struct {
	FromDomain   string `name:"from-domain" short:"f" help:"Domain to copy from."`
	ToDomain     string `name:"to-domain" short:"t" help:"Domain to copy to."`
	TemplateName string `name:"template-name" short:"n" help:"Template name to copy."`
}

type TemplateCmd struct {
	Copy TemplateCopy `cmd:"" help:"Copy a template."`
}

func (cmd *TemplateCopy) Run(globals *Globals) error {

	m := New(globals.ApiKey)

	err := m.CopyTemplate(context.Background(), cmd.FromDomain, cmd.ToDomain, cmd.TemplateName)
	return err
}
