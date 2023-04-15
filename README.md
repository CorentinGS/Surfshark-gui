# OpenVPN-GUI

A simple GUI implementation for OpenVPN to easily add my configuration options.

Been tested with surfshark config files

## How to configure

### OpenVPN

Set your environment variable

```
VPN_USERNAME=[your username]
VPN_PASSWORD=[your password]
```

### Wireguard

Download the config file from your VPN provider and place it in the `/etc/wireguard` folder.
Make sure to rename the file to the following format: `fr.par.conf` (country.city.conf)

## License

This project is licensed under the ISC License - see the [LICENSE](LICENSE) file for details

## Disclaimer

This project is not affiliated with any third party software or products. It's a personal tool that I decided to share
with others.