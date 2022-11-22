package initialize

import (
	"fmt"

	"server/config"
	"server/global"
	"server/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
}

func Tasks() {
	// log.Println(global.GVA_CONFIG)
	record := 0
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Tasks {
			go func(detail config.Task) {
				global.GVA_Timer.AddTaskByFunc("Tasking", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.Tasking(detail.TaskName, detail.Output, detail.Interval)
					record++
					if record == 10 {
						global.GVA_Timer.StopTask("Tasking")
					}
					if err != nil {
						fmt.Println("tasking error:", err)
					}
				})
			}(global.GVA_CONFIG.Timer.Tasks[i])
		}
	}
}
