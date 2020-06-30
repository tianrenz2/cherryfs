package task

import (
	"encoding/json"
	"fmt"
	"cherryfs/pkg/context"
)

func TaskHandler(taskTypeId int, taskData []byte) (error) {
	taskType := TaskType(taskTypeId)

	//var info interface{}

	switch taskType {
		case CopyObjects:
			var info []CopyTaskInfo

			json.Unmarshal(taskData, &info)
			fmt.Printf("task data: %v\n", info)
			newTask := Task{
				RunnerHostId: context.GlobalChunkCtx.HostId,
				State: PROCESSING,
				Type: CopyObjects,
				Info: info,
			}
			go newTask.ExecuteTask()
			break
	}



	return nil
}
