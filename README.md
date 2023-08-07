# mailgun-cli

A cli tool with utilities for common mailgun tasks.

## Install

```
go install github.com/manojVivek/mailgun-cli/cmd/mailgun-cli@latest
```

## Usage

```
$ mailgun-cli --help
```

To copy a template from one domain to another:

```
$ mailgun-cli template copy --help
```

So to copy a template named `my-template` from domain `my-domain.com` to domain `my-other-domain.com`:

```
$ mailgun-cli -k <api-key> template copy --domain my-domain.com --template my-template --to-domain my-other-domain.com
```

## License

MIT
