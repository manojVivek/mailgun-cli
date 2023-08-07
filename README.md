# mailgun-cli

A cli tool with utilities for common mailgun tasks.

## Install

```
# Download binary for your OS and architecture
curl -LO https://github.com/manojVivek/mailgun-cli/releases/latest/download/mailgun-cli_$(uname)_$(uname -m)
# Verify the checksum
curl -sL https://github.com/manojVivek/mailgun-cli/releases/latest/download/mailgun-cli_checksums.txt | shasum --ignore-missing -a 256 --check
# Make the binary executable
chmod a+x mailgun-cli_$(uname)_$(uname -m)
# Move to path
sudo mv mailgun-cli_$(uname)_$(uname -m) /usr/local/bin/mailgun-cli
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
