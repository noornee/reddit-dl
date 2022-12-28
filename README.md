# reddit-dl
a reddit video/gif downloader

## External Dependency used

<!--[aria2](https://aria2.github.io/) version 1.36.0-->

[ffmpeg](https://ffmpeg.org/) version 5.1.2

## STEPS

```
go install github.com/noornee/reddit-dl/cmd/reddit-dl@latest
```

make sure you have `ffmpeg` installed before running the following command:

### With Flags
`reddit-dl -url '<reddit_url>'`

### Without Flags
`reddit-dl '<reddit_url>'`
 

### Note:
>there's a minor change in the downloader; it no longer uses the external dependency [aria2](https://aria2.github.io/) to download files, it now uses a go library [grab]("github.com/cavaliergopher/grab/v3"). Aside from this change, everything works as is.


## PREVIEW
[reddit-dl.webm](https://user-images.githubusercontent.com/71889751/206926034-1022447d-b104-4998-b06c-edf6b7c04633.webm)
