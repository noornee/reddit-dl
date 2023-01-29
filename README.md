[![License: MIT](https://img.shields.io/badge/License-MIT-success.svg?style=flat-square)](https://github.com/noornee/reddit-dl/blob/main/LICENSE)

# reddit-dl
reddit-dl is a reddit media downloader written in Go.

## External Dependency used

[ffmpeg](https://ffmpeg.org/) version 5.1.2

## Installation (with `go install`)

```
go install github.com/noornee/reddit-dl/cmd/reddit-dl@latest
```

⚠  make sure you have `ffmpeg` installed before running the following command:

### With Flags
`reddit-dl -url '<reddit_url>'`

### Without Flags
`reddit-dl '<reddit_url>'`


## Installation (with binary package)
To install reddit-dl binary, go to the [releases tab](https://github.com/noornee/reddit-dl/releases/tag/build), and download the appropriate file for your OS and extract it.

🛑 do you have have `ffmpeg` installed ? 🤨


<details>
<summary> Linux </summary>

download and extract [Linux Build](https://github.com/noornee/reddit-dl/releases/download/build/reddit-dl_linux.tar.gz)  
- 64 bit `reddit-dl_linux_amd64`
- 32 bit `reddit-dl_linux_arm`
</details>

<details>
<summary> MacOs </summary>

download and extract [MacOs Build](https://github.com/noornee/reddit-dl/releases/download/build/reddit-dl_darwin.tar.gz)  
- 64 bit `reddit-dl_darwin_amd64`
- 32 bit `reddit-dl_darwin_arm64`
</details>

<details>
<summary> Windows </summary>

download and extract [Windows Build](https://github.com/noornee/reddit-dl/releases/download/build/reddit-dl_windows.tar.gz)  
then launch 
- 64 bit `reddit-dl_windows_amd64.exe`
- 32 bit `reddit-dl_windows_386.exe`
</details>



## PREVIEW
[reddit-dl.webm](https://user-images.githubusercontent.com/71889751/213936098-84d7db2f-f55b-400a-9597-09155d94e48d.webm)

<br>
