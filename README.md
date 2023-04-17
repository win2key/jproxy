# JProxy

JProxy is a simple proxy server that enables you to run a SOCKS5 proxy on a Jump Host. It can run as a Windows Service in the background, even when no one is logged into the system.

## Table of Contents

- [JProxy](#jproxy)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Troubleshooting](#troubleshooting)
  - [License](#license)
  
## Features

- Runs as a Windows Service, enabling background operation without user login
- Supports SOCKS5 proxy

## Requirements

- Go version 1.15 or higher
- Windows operating system

## Installation

1. Clone the repository to your local machine:
    ```bash
    git clone https://github.com/win2key/jproxy.git
    ```
2. Navigate to the JProxy directory:
    ```bash
    cd JProxy
    ```
3. Compile the JProxy project for the Windows platform:
    ```go
    GOOS=windows GOARCH=amd64 go build -o jproxy.exe
    ```
4. Copy the compiled `jproxy.exe` file to the desired location on your Windows Jump Host, for example:
    ```makefile
    C:\Program Files\Common Files\Services\jproxy.exe
    ```
5. Follow the instructions in the video tutorial "[JProxy Installation Guide for Windows Jump Hosts](https://www.youtube.com/watch?v=cjs_e5pVAwU)" for setting up the JProxy service on your Windows system:
  
    [![Video Thumbnail](https://img.youtube.com/vi/cjs_e5pVAwU/maxresdefault.jpg)](https://www.youtube.com/watch?v=cjs_e5pVAwU)

    For reference, the commands used in the video for CMD are:
    ```makefile
    sc create "JProxy" binPath= "C:\Program Files\Common Files\Services\jproxy.exe" 
    DisplayName= "JProxy Service" start= auto
    sc start JProxy
    ```

## Usage

Once the JProxy service is installed and started, it will run automatically in the background, listening on the specified address (default: `0.0.0.0:20202`). You can now use your Windows Jump Host as a SOCKS5 proxy.

## Troubleshooting

If you encounter any issues with JProxy, please check the logs for any error messages or unexpected behavior. The logs can provide valuable information to help diagnose the problem.

If you are still unable to resolve the issue, please open an issue on the [GitHub repository](https://github.com/win2key/jproxy/issues), providing as much detail as possible.

## License

JProxy is released under the [MIT License](https://github.com/win2key/jproxy/blob/main/LICENSE).