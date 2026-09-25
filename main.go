package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type PortStatus string

const (
	ConnOpen    PortStatus = "OPEN"
	ConnRef     PortStatus = "REFUSED"
	ConnTimeout PortStatus = "TIMEOUT"
	ConnUnknown PortStatus = "UNKNOWN"
	ConnUnreach PortStatus = "UNREACHEBLE"
)

func ClassifyError(err error) PortStatus {
	if err == nil {
		return ConnOpen
	}

	var NetErr net.Error
	if errors.As(err, &NetErr) && NetErr.Timeout() {
		return ConnTimeout
	}

	var errno syscall.Errno
	if errors.As(err, &errno) {
		switch errno {
		case syscall.ECONNREFUSED:
			return ConnRef
		case syscall.ETIMEDOUT:
			return ConnTimeout
		case syscall.EHOSTUNREACH:
			return ConnUnreach
		}
	}

	return ConnUnknown
}

func GetPorts(ports string) (int, int, error) {
	ports = strings.TrimSpace(ports)

	var start, end int

	_, err := fmt.Sscanf(ports, "%d-%d", &start, &end)
	if err == nil {
		if start > end {
			return 0, 0, fmt.Errorf("First port (%d) cant be bigger than second (%d)", start, end)
		}
		return start, end, nil
	}

	singlePort, err := strconv.Atoi(ports)
	if err == nil {
		return singlePort, singlePort, nil
	}

	return 0, 0, fmt.Errorf("Wrong input format!")
}

func main() {
	var host string
	var ports string
	fmt.Print("Enter host: ")
	fmt.Scan(&host)
	fmt.Print("Enter ports range: ")
	fmt.Scan(&ports)

	f_port, s_port, err := GetPorts(ports)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := f_port; i < s_port; i++ {
		HostPort := host + ":" + strconv.Itoa(i)

		conn, err := net.DialTimeout("tcp", HostPort, 1*time.Second)

		status := ClassifyError(err)
		if status == ConnOpen {
			fmt.Println(host, "->", i, ": OPEN")
			conn.Close()
		} else if status == ConnRef {
			fmt.Println(host, "->", i, ": REFUSED")
		} else if status == ConnTimeout {
			fmt.Println(host, "->", i, ": TIMEOUT")
		} else if status == ConnUnreach {
			fmt.Println(host, "->", i, ": UNREACHABLE")
		} else {
			fmt.Println(host, "->", i, ": UNKNOWN")
		}
	}
}
