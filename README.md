# atproxy aka Allowed-only Traffic Proxy

atproxy is a forward proxy which only allows specified traffic. The user defines the allowed traffic in the `config.json` file, by setting the `allowed_domains` (array of strings). All requests from domains that are not in the `allowed_domains` list will be rejected.

atproxy supports both *HTTP* and *HTTPS* traffic.

## Contributing

Want to help? See the [contributing guide](CONTRIBUTING.md).

## How to use

### 1. Requirements

The only real requirement to use `atproxy` is having go installed. See `go.mod` for the exact version needed.

### 2. Using atproxy

#### Through the source code

To use `atproxy` you don't really need any fancy configurations, builds or... You can simply do `make run` and you are ready to use it. Just make sure to set the `allowed_domains` in the config.

#### Through an executable

To run `atproxy` through an executable you have to build it first, but don't worry, it's very simple as well. Simple do `make build` and a new executable will appear in your current directory. You can do with this executable whatever you want, the simplest one would be to use `nohup` to run it in the background.

Something like this:
```bash
nohup ./atproxy &
```

#### Through Docker

You have a Dockerfile in the root of the project that you can use to deploy `atproxy` with Docker or Kubernetes.

## AI Usage

No code has been written by AI in this repo. The only AI usage was to ask questions regarding the behavior of proxies, not code specific.

## Useful links

- [Eli Bendersky's proxy series](https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-1-http-proxies/)
