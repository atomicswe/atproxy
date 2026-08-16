# Contributing to atproxy

Thanks for wanting to help. atproxy is a small forward proxy that only lets through traffic for domains listed in `config.json`. Contributions of all sizes are welcome—bug reports, docs, tests, and code.

## Before you start

- Open an [issue](https://github.com/atomicswe/atproxy/issues) for larger changes so we can agree on the approach first.
- For typos, docs, or small bug fixes, a pull request is enough.
- Please do not open PRs that only reformat unrelated files.

## Development setup

You need the Go version in [`go.mod`](go.mod).

Run from source (writes `./config.json` if it does not exist):

```bash
make run
```

Build a local binary (`./atproxy`):

```bash
make build
```

Run the test suite:

```bash
make test
```

`make run` sets `ATPROXY_CONFIG_PATH=./config.json`. That file is gitignored — keep secrets and local allowlists out of commits.

### Config while developing

Default listen address is `:11111` (empty `address` in config). Set `allowed_domains` or every request is rejected.

```json
{
  "server": {
    "port": 11111,
    "address": ""
  },
  "validator": {
    "allowed_domains": ["example.com"]
  }
}
```

Override the config path with `ATPROXY_CONFIG_PATH` if you do not want to use `./config.json`.

## Project layout

```
cmd/atproxy/          # main: load config, start the proxy
internal/config/      # config.json load/save
internal/server/      # HTTP listen address and server start
internal/proxy/       # allow/deny, HTTP forward, HTTPS CONNECT tunnel
internal/request/     # domain allowlist checks
```

Keep new code in the package that already owns that behavior. The public entrypoint stays `cmd/atproxy`.

## Code style

- Format with `gofmt` (or `go fmt ./...`) before you push.
- Follow existing naming: tests use `sut` for the system under test, and names like `TestAllowedAllowsRequestWithAllowedDomain`.
- Prefer the standard library. This module currently has no third-party dependencies — keep it that way unless there is a strong reason.
- Match the surrounding file: small functions, `log` for operational messages, no extra abstraction for a one-off.
- Do not commit `config.json`, `atproxy`, `cover.out`, or `tests.out`.

## Tests

Add or update tests for any behavior change.

- Unit tests live next to the code: `internal/request/validator_test.go`, `internal/server/server_test.go`.
- Name tests after the behavior, not the implementation.
- `make test` runs `go test -v ./...` with atomic coverage (`cover.out`). Coverage artifacts are gitignored.

If you change proxy forwarding or tunneling, include a test that covers the allow/deny path at minimum. Full network tests are welcome when they stay hermetic (no real external hosts).

## Pull requests

1. Branch off `main`.
2. Keep the PR focused on one change.
3. Run `make test` and `go fmt ./...` locally.
4. Describe **why** the change exists, not only what you edited. Link the issue if there is one.
5. Expect review from [@atomicswe](https://github.com/atomicswe) (see `.github/CODEOWNERS`).

## Issues

When filing a bug, include:

- atproxy commit
- OS and `go version`
- relevant `config.json` (`allowed_domains`, `port`, `address`)
- what you expected vs what happened
- logs from the process

Feature requests should explain the use case, not only the API you want.

## Security

atproxy is an allowlist proxy. Treat bypasses, accidental open proxies, and unexpected forwarding as security issues.

Please **do not** open a public issue for a suspected bypass. Contact the maintainer privately via X ([@atomicswe](https://x.com/atomicswe)).

## Questions

If something in this guide is wrong or missing, open an issue or a docs PR. That is a useful contribution too.
