# reddit-dl
a reddit video/gif downloader

## External Dependencie(s) used

<!--[aria2](https://aria2.github.io/) version 1.36.0-->

[ffmpeg](https://ffmpeg.org/) version 5.1.2

## STEPS

```
go install github.com/noornee/reddit-dl@latest

```

make sure you have `ffmpeg` installed before running the following command:

## With flags
`reddit-dl -url '<reddit_video_url>' `

## Without flags
`reddit-dl '<reddit_video_url>' `
 
## PREVIEW
### Note:
	theres a minor change in the downloader because it no longer uses this external dependency [aria2](https://aria2.github.io/) to download the files, it now uses a go library (grab)["github.com/cavaliergopher/grab/v3"] aside that, everything works as is.

[reddit-dl.webm](https://user-images.githubusercontent.com/71889751/206926034-1022447d-b104-4998-b06c-edf6b7c04633.webm)

