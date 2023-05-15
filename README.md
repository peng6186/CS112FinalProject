# This is the final project for CS112 Networds
---

# ShadowTunnel
ShadowTunnel is a socks5 proxy designed to bypass any potential network blockage and censorship to make internet browsing safer.
## 🚚 Project Structure

```
/
├── st-client/
│   ├── build/
│   ├── frontend/
│   ├── app.go
│   ├── wails.json
│   └── main.go
└── st-server/
    └── main.go
```

- `st-client` is a desktop app built with Go, React and Wails. The client listens and intercepts local traffic on a specific port and forwards them to the proxy server after encryption.
    - `/frontend/` - Frontend project files
    - `/build/` - Project build directory
    - `/app.go` - Methods exposed to front end
    - `/main.go` - The main application
    - `/wails.json` - The project configuration
- `st-server` is the proxy server that receives requests from clients, decrypts messages, and initializes messages to the actual server.
    - `/main.go` - The main application

## 💻 Dev Setup

### Client

Install Go, Node.js and Wails first.

| Command        | Action                                         |
| :------------- | :--------------------------------------------- |
| `wails dev`    | Run application with hot reload                |

### Server

Using Go implement it