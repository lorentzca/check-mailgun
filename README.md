# check-mailgun

## Description

Check state that registered domain on mailgun.

## Download

- [Releases · Lorentzca/check-mailgun · GitHub](https://github.com/Lorentzca/check-mailgun/releases)

## Setting

```
[plugin.checks.mailgun]
command = "path/to/check-mailgun -p <your mailgun apikey> -d <your domain>"
```

When you want to set interval etc, see [Adding monitors for script checks - Mackerel Docs](https://mackerel.io/docs/entry/custom-checks).

```
[plugin.checks.mailgun]
command = "path/to/check-mailgun -p <your mailgun apikey> -d <your domain>"
notification_interval = 60
max_check_attempts = 2
```

## Options

```
Application Options:
  -p, --apikey= Mailgun Api Key
  -d, --domain= Mailgun Domain

Help Options:
  -h, --help    Show this help message
```