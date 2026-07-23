# saveup
Dead-simple backup utility for self-hosted servers.

## Features
- ZIP backup
- Cloudflare R2 upload
- Configuration validation
- Linux daemon support (planned)

## Installation
Download files from [here.](https://github.com/AltheaIX/saveup/releases/latest) Then, configure the config.yaml with your R2 credentials.

Please use this reference to setting your R2 credentials.
[Click Here](https://developers.cloudflare.com/r2/get-started/s3/#2-generate-api-credentials)

## Usage
```
Usage:
  saveup diag
  saveup backup
  saveup store [file] - use absolute path from backup's output
  # Use `saveup store latest` to store the latest version after running `saveup backup`
```

## Roadmap
- [x] ZIP backup
- [x] Cloudflare R2 upload
- [ ] Automatic daemon
- [ ] Restore
- [ ] Multiple compression formats
- [ ] Support multiple workloads
