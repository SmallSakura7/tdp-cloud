package worker

import (
	"log"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/host"

	"tdp-cloud/helper/socket"
	"tdp-cloud/internal/workhub"
)

type SocketData = workhub.SocketData

type RecvPod struct {
	*socket.JsonPod
}

type RespPod struct {
	*socket.JsonPod
}

type SendPod struct {
	*socket.JsonPod
}

func Daemon(url string) {

	args := []string{}
	info, _ := host.Info()

	if osType := info.OS; len(osType) > 0 {
		args = append(args, "OSType="+osType)
	}

	if hostName := info.Hostname; len(hostName) > 0 {
		args = append(args, "HostName="+hostName)
	}

	if hostId := info.HostID; len(hostId) > 0 {
		args = append(args, "HostId="+hostId)
	}

	if len(args) > 0 {
		url += "?" + strings.Join(args, "&")
	}

	log.Println("Connecting", url)
	pod, err := socket.NewJsonPodClient(url)

	if err != nil {
		return
	}

	defer pod.Close()

	go Sender(pod)
	Receiver(pod)

}

func Sender(pod *socket.JsonPod) error {

	send := &SendPod{pod}

	for {

		if _, err := send.Ping(); err != nil {
			log.Println("Send:error", err)
			return err
		}

		time.Sleep(time.Second * 15)

	}

}

func Receiver(pod *socket.JsonPod) error {

	recv := &RecvPod{pod}
	resp := &RespPod{pod}

	for {
		var rs *SocketData

		if err := pod.Read(&rs); err != nil {
			log.Println("Read:error", err)
			return err
		}

		switch rs.Method {
		case "Exec":
			recv.Exec(rs)
		case "Ping:resp":
			resp.Ping(rs)
		default:
			log.Println("Task:unknown", rs)
		}
	}

}
