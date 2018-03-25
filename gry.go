package main

import (
    "encoding/binary"
    "encoding/json"
    "net"

    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/host"
    psnet "github.com/shirou/gopsutil/net"
)

type Recon struct {
    Cpu          []cpu.InfoStat          `json: "cpu"`
    Connections  []psnet.ConnectionStat  `json: "connections"`
    Partitions   []disk.PartitionStat    `json: "partitions"`
    Host         *host.InfoStat          `json: "host"`
    Temperature  []host.TemperatureStat  `json: "temperature"`
    Users        []host.UserStat         `json: "users"`
}

func (h Recon) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}

func main() {
    srvAddr := "localhost:8899"
    tcpAddr, _ := net.ResolveTCPAddr("tcp", srvAddr)

    cpu_info, _ := cpu.Info()
    connections, _ := psnet.Connections("all")
    disks, _ := disk.Partitions(false)
    host_info, _ := host.Info()
    temperature, _ := host.SensorsTemperatures()
    users, _ := host.Users()
    recon := &Recon{Cpu: cpu_info, Connections: connections, Partitions: disks, Host: host_info, Temperature: temperature, Users: users}
    reconString := recon.String()
    payloadSize := make([]byte, 4)
    binary.BigEndian.PutUint32(payloadSize, uint32(len(reconString)))

    conn, _ := net.DialTCP("tcp", nil, tcpAddr)
    _, _ = conn.Write(payloadSize)
    _, _ = conn.Write([]byte(reconString))
}
