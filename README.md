# inonius_v3cli

inonius_v3cli is a command line tool for [inonius speedtest](https://inonius.net) v3.


```bash
Usage:
  inonius_v3cli [flags]

Flags:
  -c, --config string          --config <CONFIG_PATH> YML, TOML and JSON are available. (default ./config.yml)
  -d, --debug                  Debug mode
      --deviceid string        custom device id (default: hostname based generate)
  -e, --endpoint string        Use: client --endpoint <ENDPOINT> (default "https://api.inonius.net")
  -?, --help                   Show help
      --icmp                   Use ICMP ping (default: http ping)
  -k, --ignore-tls-error       Ignore tls error
  -4, --ipv4                   Force IPv4
      --ipv4-endpoint string   Use: client --ipv4-endpoint <ENDPOINT> (default "https://ipv4-api.inonius.net")
  -6, --ipv6                   Force IPv6
      --ipv6-endpoint string   Use: client --ipv4-endpoint <ENDPOINT> (default "https://ipv6-api.inonius.net")
  -O, --orgtag string          OrgTag if you have
  -q, --quiet                  Quiet mode
      --json                   Output as JSON
  -i, --interface string       Interface Name
  -s, --source string          Source address
  -v, --version                version for inonius_v3cli
```


<details>
<summary>Json output example:</summary>
  
```json
{
    "timestamp": 1734670201,
    "ipv4_available": true,
    "ipv6_available": true,
    "ipv4_info": {
        "ip": "203.0.113.133",
        "port": 59892,
        "is_ipv4": true,
        "org": "AS64512 Alice Corp."
    },
    "ipv6_info": {
        "ip": "3fff:1:0:1001::200e:20",
        "port": 26811,
        "is_ipv4": false,
        "org": "AS64513 Bob Inc."
    },
    "result": [
        {
            "speedtest_type": "IPv4",
            "timestamp": 1734671365,
            "server": "ipv4-librespeed2",
            "upload": 114.76,
            "download": 67.23,
            "ping": 7,
            "jitter": 0.53
        },
        {
            "speedtest_type": "IPv6",
            "timestamp": 1734671365,
            "server": "ipv6-librespeed1",
            "upload": 34.86,
            "download": 49.2,
            "ping": 94.82,
            "jitter": 26.05
        }
    ]
}
```
</details>
