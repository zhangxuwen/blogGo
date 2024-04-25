package snowflake

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

func GetID() int64 {
	return node.Generate().Int64()
}

func main() {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	id := GetID()
	fmt.Println(id)
}
