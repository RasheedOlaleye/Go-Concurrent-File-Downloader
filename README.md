# Go Concurrent File Downloader

A concurrent file downloader written in Go that downloads multiple files simultaneously using goroutines, channels, and sync.WaitGroup.

This project was built to learn Go concurrency by implementing a real-world application instead of simple examples.

## Features

- Concurrent file downloads
- Goroutines for parallel execution
- sync.WaitGroup for synchronization
- Channels for communicating download results
- Automatic creation of the downloads directory
- Reads download URLs from a `urls.txt` file
- Error handling for failed downloads

## Project Structure

```
.
├── downloads/
├── urls.txt
├── main.go
├── go.mod
└── README.md
```

## How It Works

1. Read URLs from `urls.txt`.
2. Launch one goroutine per download.
3. Synchronize all downloads using `sync.WaitGroup`.
4. Send status updates through a channel.
5. Print download results as they complete.

## Technologies

- Go
- Goroutines
- Channels
- sync.WaitGroup
- net/http
- File I/O

## What I Learned

- How goroutines execute concurrently.
- Synchronizing goroutines with `sync.WaitGroup`.
- Communicating between goroutines using channels.
- Avoiding deadlocks.
- Reading files using `bufio.Scanner`.
- Downloading files over HTTP.

## Future Improvements

- Retry failed downloads
- Progress bars
- Resume interrupted downloads
- Download speed tracking
- CLI arguments
- Docker support
- Unit tests

## Run Locally

```bash
git clone <repository-url>
cd go

go run .
```

## Engineering Notes

This project is part of my Go Engineering roadmap, where I build production-inspired software to master backend engineering, concurrency, networking, and distributed systems.

## License

MIT
