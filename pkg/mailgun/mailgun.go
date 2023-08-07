package mailgun

import (
	"context"
	"fmt"

	"github.com/mailgun/mailgun-go/v4"
)

type Mailgun struct {
	ApiKey string
}

func New(apiKey string) *Mailgun {
	return &Mailgun{ApiKey: apiKey}
}

func (m *Mailgun) CopyTemplate(ctx context.Context, fromDomain, toDomain, templateName string) error {
	fromM := mailgun.NewMailgun(fromDomain, m.ApiKey)
	toM := mailgun.NewMailgun(toDomain, m.ApiKey)

	template, err := fromM.GetTemplate(ctx, templateName)
	if err != nil {
		return fmt.Errorf("reading template: %w", err)
	}

	err = toM.CreateTemplate(ctx, &template)

	if err != nil {
		return fmt.Errorf("creating template: %w", err)
	}

	return nil
}
