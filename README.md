# atproxy aka Allowed-only Traffic Proxy

atproxy is a forward proxy which only allows specified traffic. The user defines the allowed traffic in the `config.json` file, by setting the `allowed_domains` (array of strings). All requests from domains that are not in the `allowed_domains` list will be rejected.

atproxy supports both *HTTP* and *HTTPS* traffic.

## Useful links

- [Eli Bendersky's proxy series](https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-1-http-proxies/)
