# Shell Packet

Shell Packet is a cross-platform SSH VPN client that routes all of your machine's traffic through an SSH tunnel. It's a simple and secure way to encrypt your traffic and bypass network restrictions.

## Features

*   **Cross-platform:** Works on Linux, Windows, and macOS.
*   **Secure:** Encrypts all of your traffic using SSH.
*   **Performant:** Uses modern ciphers and compression to optimize performance.
*   **Easy to use:** Simple command-line interface.

## How it works

Shell Packet works by creating a virtual network interface (a TUN device) and then routing all of your machine's traffic through it. The traffic is then forwarded over an SSH connection to a remote server, which then forwards it to the internet.

This has two main benefits:

1.  **Security:** All of your traffic is encrypted, so it's safe from eavesdropping.
2.  **Bypassing restrictions:** You can use Shell Packet to bypass firewalls and other network restrictions.

## Getting Started

To get started with Shell Packet, you'll need to have a remote server that you can connect to via SSH. You'll also need to have Go installed on your local machine.

### Building the application

You can build the application using the included `Makefile`:

```bash
# Build for your current OS
make build

# Build for Linux
make build-linux

# Build for Windows
make build-windows

# Build for macOS
make build-mac
```

This will create a binary file in the current directory.

### Running the application

To run the application, you'll need to provide your SSH credentials:

```bash
sudo ./shell-packet <user> <host:port> <password>
```

**Note:** You'll need to run the application with `sudo` because it needs to create a TUN device and modify the routing table.

## Platform-specific instructions

### Linux

On Linux, Shell Packet uses the `ip` command to configure the TUN device. You'll need to have the `iproute2` package installed.

### Windows

On Windows, Shell Packet uses the `netsh` command to configure the TUN device. You'll also need to install the `wintun` driver. You can download it from the [wintun website](https://www.wintun.net/).

### macOS

On macOS, Shell Packet uses the `ifconfig` and `route` commands to configure the TUN device. These commands should be available by default.
