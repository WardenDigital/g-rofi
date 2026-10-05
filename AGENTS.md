# AGENTS.md

Go CLI that wraps the `rofi` launcher to provide prompt/selection UIs for shell workflows. Built with Cobra. Distributed via a Nix flake with a Home Manager module.

## Commands

- Build: `make build` (outputs `bin/g-rofi`, uses `CGO_ENABLED=0 go build ... *.go`)
- Run from source: `go run . <subcommand>`
- Nix build: `nix build` (result symlink lands in repo root; gitignored)
- There are **no tests** and no lint/CI config. `make test`, `fmt`, `vet`, `deps`, `help` are declared in `.PHONY` but have **no targets defined** — they do nothing.
- `make install` is broken: it only `chmod`s `bin/g-rofi` and never copies it to `INSTALL_DIR` (`/usr/local/bin`). Don't rely on it.

## CLI surface

Cobra commands are defined in `cmd/` and registered via `init()` on the package-level `rootCmd` (calling `cmd.Execute()` from `main.go`). Subcommand names are camelCase:

- `g-rofi browserSearch [-b <browser>] [-e <engine>]` — defaults: `brave-browser`, `google`
- `g-rofi shellCommand` — prompts for a command string, runs it synchronously, and retries with a masked rofi password prompt via `sudo -S` when the command fails needing a password

Flow: `cmd/` → handler in `internal/handlers/*` → `internal/rofi` (spawns rofi) → action. Errors in handlers are handled inconsistently: `search.Search` **panics** on error, `terminal.OneShotCommand` returns an error that the cobra `Run` closure silently discards.

## Config resolution (non-obvious)

All rofi invocations pass `-config <path>/config.rasi`. The config directory is resolved in `internal/rofi/general.go`:

1. `G_ROFI_CONFIG` env var (set by the Home Manager module)
2. `THEME_DIR` env var (set by the Nix `wrapProgram`)
3. `.` (current directory) as fallback — running from the repo root works because `config/config.rasi` sits at `./config.rasi`. Inside `config/`, `config.rasi` does `@import "myConfig"`, which is how `myConfig.rasi`/`myTheme.rasi` get pulled in — don't rename them.

The inline prompt additionally passes `-theme <dir>/inlinePrompt.rasi`. `G_ROFI_CONFIG` takes precedence over `THEME_DIR` — both are set when installed via the Home Manager module with the flake's wrapper, so `G_ROFI_CONFIG` always wins in that case.

## Gotchas

- **Externals required**: nothing works without the `rofi` binary on `PATH` and an X11/Wayland session. The flake wrapper adds rofi to PATH via `--prefix PATH`.
- **`NewBrowser` argument reversal bug** (`internal/handlers/search/browsers.go:55`): `strings.Contains("brave-browser", strings.ToLower(n))` checks whether the *constant* contains the *flag value*, reversed. Any value that isn't a substring of `"brave-browser"` (e.g. `librewolf`) fails the Brave check and falls through to the Librewolf branch only if it's a substring of `"librewolf"` (which works since the value `librewolf` is compared against itself); anything else silently becomes a Brave browser with the user's binary name. Handle with care.
- **`ExecuteCommand` runs synchronously with password retry** (`internal/command/command.go`): since `shellCommand` uses `sh -c`, sudo-style commands fail without a tty ("a terminal is required to read the password"). `ExecuteCommand` runs the command with `cmd.Run()` (stderr tee'd to the user *and* captured), and if stderr matches `passwordSignals` (password/permission denied/no tty/askpass), it prompts via `rofi.Password` (masked `-password` dmenu) and retries with `sudo -p "" -S sh -c <cmd>`, up to `maxPasswordAttempts` (3), re-prompting only while the failure still looks password-related. Note the retry runs the command **as root** — the masked prompt is the only consent gate. Browser `Open*` methods are still fire-and-forget (`.Start()`, errors ignored).
- **No URL encoding**: search URLs are built by plain string concatenation (`"https://www.google.com/search?q=" + query`); queries with spaces/`&` will produce broken URLs.
- **`Select` pipes options via `printf`**: `internal/rofi/select.go` builds a `\n`-joined string and pipes it into rofi through an `os.Pipe`; also prints debug lines ("Piped options string created") to stdout — noisy when used programmatically.
- **Keep flake vendorHash in sync**: `flake.nix` pins `vendorHash` for `buildGoModule`; adding a Go dependency requires updating it (or `nix flake update`/`nix build` will fail hash verification).
- **Nix packaging**: `postInstall` copies `config/*` to `$out/share/config` and sets `THEME_DIR` and `XDG_DATA_DIRS`. The Home Manager module (`nix/hm-module.nix`) activates `g-rofi`, exposes `programs.g-rofi.{enable,package,configPath}`, and sets the `G_ROFI_CONFIG` session variable (custom `configPath` or the package's `share/config`). Home-manager module input is wired via `flake.output.homeManagerModules.default`.
- **Go version**: `go.mod` declares `go 1.25.6`; module path is `github.com/WardenDigital/g-rofi` (imports use this prefix, not a local path). Only external dependency: `github.com/spf13/cobra`.
- **Style**: no tests anywhere; gopls flags a `WriteString` string-concat warning in `select.go:47` — leave it or fix deliberately (existing code style is minimal/unpolished).

## Layout

```
main.go                    entry point → cmd.Execute()
cmd/                       cobra commands (root, browserSearch, shellCommand)
internal/rofi/             rofi process helpers: general.go (config + cmd builders), prompt.go, select.go
internal/handlers/search/  browser/engine interfaces + factories (browsers.go, engines.go, search.go)
internal/handlers/terminal/ one-shot shell command
internal/command/          sh -c execution
config/                    rofi .rasi files (config.rasi imports myConfig; myTheme.rasi is Catppuccin Mocha)
nix/hm-module.nix          Home Manager module
flake.nix                  package + homeManagerModules.default output
```