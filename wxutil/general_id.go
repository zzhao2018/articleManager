package wxutil

import (
	"fmt"
	"github.com/sony/sonyflake"
	"log"
)

var(
	sonyFlake *sonyflake.Sonyflake
	sonyMachineId uint16
)


//初始化
func InitSonyFlake(machineId uint16)(error){
	sonyConf:=sonyflake.Settings{
		MachineID: func() (u uint16, e error) {
			return machineId,e
		},
		CheckMachineID: nil,
	}
	sonyMachineId=machineId
	sonyFlake=sonyflake.NewSonyflake(sonyConf)
	return nil
}


//获得全局id
func GetGeneralId()(uint64,error){
	if sonyFlake==nil {
		log.Printf("flake is nil\n")
		err:=fmt.Errorf("flake is nil")
		return 0,err
	}
	return sonyFlake.NextID()
}