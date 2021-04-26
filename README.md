# go_port_scanner

A Go tool to scan ports on IPs using Goroutines to be more efficient.

# Build
`go build ./`

# Test
`go test -v ./`

# Examples
## Scan one IP and port.
`go_port_scanner -h 8.8.8.8 -p 53`

## Scan more than one host on more than one port.
`go_port_scanner -h 8.8.8.8,8.8.4.4. -p 53-55`