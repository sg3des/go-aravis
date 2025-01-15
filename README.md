# go-aravis

Go wrapper around libaravis 0.10

go bindings to libaravis 0.8 - https://github.com/hybridgroup/go-aravis

go bindings to libaravis 0.6 - https://github.com/thinkski/go-aravis

## Quickstart

How to get the number of connected devices:

```go
import aravis
import log

func main() {
    aravis.UpdateDeviceList()

    n, err := aravis.GetNumDevices()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Devices:", n)
}
```

## Troubleshooting

GigE Vision cameras often use large packet sizes, well in excess of the typical MTU (maximum transmit unit) of most network interface cards, to save on packet overhead. Be sure to first set the MTU of the network interface(s) with GigE Vision cameras to 9000 bytes. For instance, if the network interface is `enp2s0`:

    ip link set enp2s0 mtu 9000
