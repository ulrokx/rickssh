# RickSSH

This is a work in progress SSH server written in Go. I am following along with the following RFCs:

- [RFC 4253](https://datatracker.ietf.org/doc/rfc4253/) - The Secure Shell (SSH) Transport Layer Protocol
    - This RFC helped to understand a lot of the components of the SSH protocol. 
- [RFC 4250](https://datatracker.ietf.org/doc/rfc4250/) - The Secure Shell (SSH) Protocol Assigned Numbers
    - I used (will use) this RFC to get the numbers for the different stages of the protocol.
- [RFC 4251](https://datatracker.ietf.org/doc/rfc4251/) - The Secure Shell (SSH) Protocol Architecture
    - This RFC helped to understand the architecture of the SSH protocol.

## Usage

```bash
go run cmd/main/rickssh.go
```

You can specify the following flags:

- `-d` enables debug logging
- `-t` enables trace logging and debug logging