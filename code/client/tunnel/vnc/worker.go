package vnc

import (
	"fmt"
	"natpass/code/client/tunnel/vnc/worker"

	"github.com/gorilla/websocket"
	"github.com/lwch/runtime"
)

// RunWorker run vnc worker
func RunWorker(port uint16, cursor bool) {
	worker := worker.NewWorker(cursor)
	if worker == nil {
		panic("build context failed")
	}
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://127.0.0.1:%d", port), nil)
	runtime.Assert(err)
	worker.Do(conn)
}
