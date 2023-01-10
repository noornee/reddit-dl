[![Windows Build](https://github.com/noornee/reddit-dl/actions/workflows/build_windows.yml/badge.svg)](https://github.com/noornee/reddit-dl/actions/workflows/build_windows.yml) [![Linux Build](https://github.com/noornee/reddit-dl/actions/workflows/build_linux.yml/badge.svg)](https://github.com/noornee/reddit-dl/actions/workflows/build_linux.yml) [![MacOS Build](https://github.com/noornee/reddit-dl/actions/workflows/build_macos.yml/badge.svg)](https://github.com/noornee/reddit-dl/actions/workflows/build_macos.yml)


# reddit-dl
reddit-dl is a reddit media downloader written in Go.

## External Dependency used

[ffmpeg](https://ffmpeg.org/) version 5.1.2

## Installation (with `go install`)

```
go install github.com/noornee/reddit-dl/cmd/reddit-dl@latest
```

make sure you have `ffmpeg` installed before running the following command:

### With Flags
`reddit-dl -url '<reddit_url>'`

### Without Flags
`reddit-dl '<reddit_url>'`

<br>

## Installation (with binary package)
To install reddit-dl binary, go to the [releases tab](https://github.com/noornee/reddit-dl/releases/tag/build), and download the appropriate zip for your OS + ARCH, and extract it.


make sure you have `ffmpeg` installed

### Windows
download and extract [Windows Build](https://github.com/noornee/reddit-dl/releases/download/build/reddit-dl_windows.zip)  
then launch 
- 64 bit `reddit-dl_amd64-windows.exe`
- 32 bit `reddit-dl_i386-windows.exe`

### MacOs
download and extract [MacOs Build](https://github.com/noornee/reddit-dl/releases/download/build/reddit-dl_osx.tar.gz)  
- 64 bit `reddit-dl_amd64-osx`
- 32 bit `reddit-dl_arm64-osx`

### Linux
download and extract [Linux](https://github.com/noornee/reddit-dl/releases/download/build/reddit-dl_linux.tar.gz)  
- 64 bit `reddit-dl_amd64-linux`
- 32 bit `reddit-dl_i386-linux`
<br>

## PREVIEW
[reddit-dl.webm](https://user-images.githubusercontent.com/71889751/206926034-1022447d-b104-4998-b06c-edf6b7c04633.webm)

<br>

## DISCLAIMER:
<b>reddit-dl</b> is still in its early stages and might be prone to bugs.
the source code is also quite a mess but there's going to be improvements as time goes
