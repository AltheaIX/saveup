# saveup
Dead-simple backup utility for self-hosted servers.

## Features
- ZIP backup
- Supports S3 compatible object storage
- Configuration validation
- Cross-platform daemon support (Windows / Linux)

## Installation
Download files from [here.](https://github.com/AltheaIX/saveup/releases/latest) Then, configure the config.yaml with your S3 credentials.

Please use this reference to setting your S3 credentials.
[Click Here](https://developers.cloudflare.com/r2/get-started/s3/#2-generate-api-credentials)

## Usage
```
Usage:
  saveup diag
  saveup backup
  saveup store [file] - use absolute path from backup's output
  # Use `saveup store latest` to store the latest version after running `saveup backup`
  saveup daemon - start daemon for auto-backup
```

## Roadmap
- [x] ZIP backup
- [x] Cloudflare R2 upload
- [X] Automatic daemon
- [X] Retention Policy
- [ ] Restore or Download File directly
- [ ] Multiple compression formats
- [ ] Support multiple workloads
