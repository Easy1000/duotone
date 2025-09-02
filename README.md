# Duotone CLI

A simple command-line tool to apply a **duotone effect** on images.  
Currently supports **Linux**.

---

## Installation

### 1. Download Binary (Linux only)

1. Go to the [Releases](https://github.com/Easy1000/duotone/releases) page.
2. Download the latest binary for Linux.
3. Make it executable:

   ```bash
   chmod +x duotone
   ```

4. Move it to a directory in your PATH, for example

   ```bash
   sudo mv duotone /usr/local/bin/
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
