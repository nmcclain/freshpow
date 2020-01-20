# FreshPow: ski resort lift and trail status alerts

`FreshPow` tracks public ski resort lift and trail status, alerting you about changes.

## Usage

* You'll need to set the SLACK_WEBHOOK and SLACK_CHANNEL environment variables to enable notifications.
* Try it from the command-line to test things out.
* Run it as a service for "production" usage.
* Please be considerate when lowering the `--delay` setting below the default of 5m.

```
freshpow: trail and lift status alerts
Usage:
  freshpow [options] <resort>
  freshpow --help
  freshpow --version

Options:
  -l, --list                           List supported resort names.
  -d, --delay=<delay>                  Delay between requests [default: 5m].
  --debug                              Display debugging messages.
  -h, --help                           Show this screen.
  --version                            Show version.<Paste>
```

Example:
```
$ ./freshpow --debug Vail
```

Supported ski resorts:
```
$ ./freshpow -l
Supported resort names: Eldora, Steamboat, Keystone, CrestedButte, Copper, Snowshoe, Blue, Breckenridge, BeaverCreek, Dev, Stratton, Tremblant, WinterPark, Vail
```

## Contributing

**Something bugging you?** Please open an [Issue](https://github.com/nmcclain/freshpow/issues) or [Pull Request](https://github.com/nmcclain/freshpow/pulls) - we're here to help!

**New Feature Ideas?** Please open a [Pull Request](https://github.com/nmcclain/freshpow/pulls).
 
**All Humans Are Equal In This Project And Will Be Treated With Respect.**
