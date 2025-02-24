package pool

import (
	"natpass/code/network"
	"time"

	"google.golang.org/protobuf/proto"
)

// SendShellData send shell data
func (conn *Conn) SendShellData(to string, toIdx uint32, id string, data []byte) uint64 {
	dup := func(data []byte) []byte {
		ret := make([]byte, len(data))
		copy(ret, data)
		return ret
	}
	var msg network.Msg
	msg.To = to
	msg.ToIdx = toIdx
	msg.XType = network.Msg_shell_data
	msg.LinkId = id
	msg.Payload = &network.Msg_Sdata{
		Sdata: &network.ShellData{
			Data: dup(data),
		},
	}
	select {
	case conn.write <- &msg:
		data, _ := proto.Marshal(&msg)
		return uint64(len(data))
	case <-time.After(conn.parent.cfg.WriteTimeout):
		return 0
	}
}

// SendShellResize send shell resize
func (conn *Conn) SendShellResize(to string, toIdx uint32, id string, rows, cols uint32) {
	var msg network.Msg
	msg.To = to
	msg.ToIdx = toIdx
	msg.XType = network.Msg_shell_resize
	msg.LinkId = id
	msg.Payload = &network.Msg_Sresize{
		Sresize: &network.ShellResize{
			Rows: rows,
			Cols: cols,
		},
	}
	select {
	case conn.write <- &msg:
	case <-time.After(conn.parent.cfg.WriteTimeout):
	}
}
