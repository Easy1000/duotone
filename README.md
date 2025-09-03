# Duotone CLI

A simple command-line tool to apply a **duotone effect** on images.  
Currently supports **Linux**.

> ⚡ This project is inspired by the duotone profile pictures people are currently using as a show of support for the protests in Indonesia.

---

## Installation

### 1. Install via `.deb` Package (Linux)

1. Go to the [Releases](https://github.com/Easy1000/duotone/releases) page.
2. Download the latest `.deb` package (e.g. `duotone_0.1.2_amd64.deb`).
3. Install it using `dpkg`:

   ```bash
   sudo dpkg -i duotone_0.1.2_amd64.deb
   ```

4. Verify installation:

   ```bash
   duotone --help
   ```

### 2. Install via Go

If you have Go installed, you can install directly from source

```bash
go install github.com/Easy1000/duotone@latest
```

This will place the duotone binary in your `GOPATH/bin` or `$HOME/go/bin`

## Usage

```bash
duotone convert <input-file>
```

example:

```bash
duotone convert input.png
```

This will generate duotone.jpg with duotone effect applied
