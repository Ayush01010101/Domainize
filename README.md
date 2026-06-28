# domainize

A command line tool for serving your local development servers under real custom
domains over HTTPS. Map a port like `localhost:8080` to a domain like
`example.com`, and `domainize` runs a reverse proxy that forwards traffic to your
app with a trusted, locally signed TLS certificate.

## About

When you run a local server, you usually reach it at `http://localhost:<port>`.
That is fine for quick checks, but it breaks down when you need a real domain and
real HTTPS, for example when testing OAuth callbacks, cookies scoped to a domain,
subdomains, or anything that behaves differently outside of `localhost`.

`domainize` solves this by tying three pieces together:

- It adds an entry to your system hosts file so the chosen domain resolves to your
  machine (`127.0.0.1`).
- It uses [mkcert](https://github.com/FiloSottile/mkcert) to generate a TLS
  certificate that your browser trusts, so the domain loads over HTTPS with no
  warnings.
- It runs a reverse proxy on ports 80 and 443 that forwards requests for the
  domain to the local port your app is listening on.

The result is that `https://example.com` on your machine transparently serves the
app running on `http://localhost:8080`.

## Features

- Map any local port to a custom domain.
- Automatic HTTPS using locally trusted certificates (no browser warnings).
- Automatic hosts file management (entries are added on link and cleaned up on
  overwrite).
- Cross platform: Linux, macOS, and Windows.
- mkcert is downloaded and configured for you during setup.

## Requirements

- Administrative privileges. `domainize` binds to the privileged ports 80 and 443
  and edits the system hosts file, so most commands must be run with `sudo` on
  Linux and macOS, or from an elevated (Administrator) terminal on Windows.
- Go 1.26 or newer, only if you are building from source.
- An active internet connection the first time you run `setup`, so mkcert can be
  downloaded.

## Installation

### Build from source

```bash
git clone https://github.com/Ayush01010101/Custom-Domain-CLI.git
cd Custom-Domain-CLI
go build -o domainize
```

Then move the binary somewhere on your `PATH` so you can call it from anywhere:

```bash
sudo mv domainize /usr/local/bin/
```

On Windows, build with `go build -o domainize.exe` and place `domainize.exe` in a
directory that is on your `PATH`.

### Download a release

Prebuilt archives are published on the
[releases page](https://github.com/Ayush01010101/Custom-Domain-CLI/releases).
Download the archive for your operating system and architecture, extract it,
rename the binary to `domainize` (or `domainize.exe` on Windows), and move it onto
your `PATH` as shown above.

## Usage

The typical workflow is: run `setup` once, `link` a port to a domain, then `start`
the proxy.

### 1. Set up

Run this once to download mkcert, install the local certificate authority, and
create the configuration file.

```bash
sudo domainize setup
```

If a configuration already exists and you want to start fresh, reset it:

```bash
sudo domainize setup --reset
```

### 2. Link a port to a domain

Map a local port to the domain you want to use. The first argument is the port and
the second is the domain.

```bash
sudo domainize link 8080 example.com
```

This adds a `127.0.0.1 example.com` entry to your hosts file, generates a trusted
certificate for the domain, and records the mapping in the configuration file. If
a domain is already configured, you will be asked whether to overwrite it.

### 3. Start the proxy

Start serving every linked domain. Make sure your application is already running on
the port you linked.

```bash
sudo domainize start
```

The proxy now listens on ports 80 and 443. Open `https://example.com` in your
browser and you will reach the app running on `http://localhost:8080`.

### Check status

Verify that `domainize` is set up correctly.

```bash
domainize status
```

### Help

List all commands and flags at any time:

```bash
domainize --help
domainize <command> --help
```

## How it works

```
Browser  ->  https://example.com (resolved to 127.0.0.1 via hosts file)
         ->  domainize reverse proxy on ports 80 and 443 (TLS via mkcert)
         ->  http://localhost:8080 (your app)
```

## Configuration

`domainize` stores its configuration, the downloaded mkcert binary, and the
generated certificates in your user configuration directory:

- Linux: `~/.config/domainize/`
- macOS: `~/Library/Application Support/domainize/`
- Windows: `%AppData%\Roaming\domainize\`

The configuration file (`config.json`) looks like this after linking a domain:

```json
{
  "name": "domainize",
  "domain": {
    "example.com": {
      "https": true,
      "port": 8080
    }
  }
}
```

Hosts file entries are written to:

- Linux and macOS: `/etc/hosts`
- Windows: `C:\Windows\System32\drivers\etc\hosts`

## Troubleshooting

- Permission denied or "address already in use" on start: ensure you are running
  with administrative privileges and that nothing else is using ports 80 or 443.
- Certificate not trusted: re-run `sudo domainize setup` so the local certificate
  authority is installed in your system trust store.
- Domain does not resolve: confirm the `127.0.0.1 <domain>` line exists in your
  hosts file, which `link` should have added.

## License

See the [LICENSE](LICENSE) file for details.
