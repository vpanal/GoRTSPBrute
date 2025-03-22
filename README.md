# GoRTSPBrute - RTSP Path Bruteforce

RTSPBrute is a tool written in Go to perform brute force attacks on RTSP paths in IP cameras and streaming servers. It uses the **DESCRIBE** method to verify if a stream actually exists on the tested path, avoiding false positives.

## Features
✅ Fast and multithreaded scanning.  
✅ Uses the **DESCRIBE** method instead of **OPTIONS**, improving accuracy.  
✅ Filters out false paths by verifying that the response contains stream information.  
✅ Compatible with standard RTSP servers.

## Requirements
- Go 1.18 or higher

## Installation
Clone the repository and compile the binary:
```bash
git clone https://github.com/vpanal/gortspbrute.git
cd gortspbrute
go build main.go -o gortspbrute
```

## Usage
Run the binary specifying the IP:PORT of the RTSP server and the path wordlist:
```bash
./gortspbrute 192.168.1.54:8554 wordlist.txt
```

Example wordlist:
```
stream
live.sdp
video
cam1
axis-media/media.amp
```

## Example Output
```bash
[+] Valid stream found: rtsp://192.168.1.54:8554/live.sdp
[-] Invalid path or no stream: rtsp://192.168.1.54:8554/test
[+] Valid stream found: rtsp://192.168.1.54:8554/axis-media/media.amp
Bruteforce completed.
```

## Notes
- If a path returns **401 Unauthorized**, it means it is valid but requires authentication.
- If VLC cannot open the URL, the stream is likely not accessible.

## Author
Created by vpanal for educational purposes and security testing in controlled environments. I am not responsible for misuse.

