package timeTaskServiceModules

import (
	"errors"
	"fmt"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strconv"
	"time"
)

type timeTask struct {
	TaskID      int64
	TaskName    string
	Status      bool
	Cron        string
	TaskFunc    func()
	LastRunTime time.Time
	RunTimes    int
}

// 初始化任务列表
var tasks = []timeTask{
	{
		TaskID:      utils.GenID(),
		TaskName:    "task1",
		Cron:        "@every 1s",
		Status:      false,
		LastRunTime: time.Now(),
		RunTimes:    0,
		TaskFunc: func() {
			fmt.Println("task1")
		},
	},
	{
		TaskID:      utils.GenID(),
		TaskName:    "task2",
		Cron:        "*/10 * * * * *",
		Status:      false,
		LastRunTime: time.Now(),
		RunTimes:    0,
		TaskFunc: func() {
			fmt.Println("task2")
		},
	},
}

func init() {
	for _, task := range tasks {
		err := utils.TimeTask.AddTask(strconv.FormatInt(task.TaskID, 10), task.Cron, task.TaskFunc)
		if err != nil {
			utils.Log.Error(err)
		}
	}
}

var TimeTaskService = &timeTaskService{
	Task: tasks,
}

type timeTaskService struct {
	Task []timeTask
}

func (t *timeTaskService) GetTimeTask() []dto.TimeTaskResponse {
	var timeTaskResponse []dto.TimeTaskResponse
	for _, task := range t.Task {

		// 获取当前时间
		now := time.Now()
		// 定义一个目标时间
		target := time.Date(task.LastRunTime.Year(), task.LastRunTime.Month(), task.LastRunTime.Day(), task.LastRunTime.Hour(), task.LastRunTime.Minute(), task.LastRunTime.Second(), 0, task.LastRunTime.Location())
		// 计算时间差值

		var diff time.Duration
		if task.Status {
			diff = now.Sub(target)
		}

		timeTaskResponse = append(timeTaskResponse, dto.TimeTaskResponse{
			TimeTaskID:   strconv.FormatInt(task.TaskID, 10),
			TimeTaskName: task.TaskName,
			Cron:         task.Cron,
			Status:       task.Status,
			LastRunTime:  task.LastRunTime.Format("2006/01/02 03:04:05"),
			RunTimes:     diff.Seconds(),
		})
	}
	return timeTaskResponse
}

func (t *timeTaskService) UpdateTask(id string, params *dto.UpdateTimeTask) error {
	// 校验 cron 表达式是否正确
	if err := utils.TimeTask.ParseCron(params.Cron); err != nil {
		return err
	}

	var updated bool
	for i, task := range t.Task {
		if strconv.FormatInt(task.TaskID, 10) == id {
			t.Task[i].Cron = params.Cron
			t.Task[i].TaskName = params.TimeTaskName

			if err := utils.TimeTask.RemoveTask(strconv.FormatInt(t.Task[i].TaskID, 10)); err != nil {
				return err
			}

			if err := utils.TimeTask.AddTask(strconv.FormatInt(t.Task[i].TaskID, 10), t.Task[i].Cron, t.Task[i].TaskFunc); err != nil {
				return err
			}

			updated = true
			if params.Status {
				return t.StartTimeTask(strconv.FormatInt(task.TaskID, 10))
			} else {
				return t.StopTimeTask(strconv.FormatInt(task.TaskID, 10))
			}
		}
	}
	if !updated {
		return errors.New("任务不存在")
	}
	return nil
}

func (t *timeTaskService) StartTimeTask(TimeTaskID string) error {
	var err error
	for i, task := range t.Task {
		if strconv.FormatInt(task.TaskID, 10) == TimeTaskID {
			t.Task[i].Status = true
			t.Task[i].LastRunTime = time.Now()
			err = utils.TimeTask.StartTask(TimeTaskID)
		}
	}
	return err
}

func (t *timeTaskService) StopTimeTask(TimeTaskID string) error {
	if err := utils.TimeTask.StopTask(TimeTaskID); err != nil {
		return err
	} else {
		for i, task := range t.Task {
			if strconv.FormatInt(task.TaskID, 10) == TimeTaskID {
				t.Task[i].Status = false
				break
			}
		}
		return nil
	}
}

func (t *timeTaskService) RemoveTimeTask(TimeTaskID string) error {
	if err := utils.TimeTask.RemoveTask(TimeTaskID); err != nil {
		return err
	}
	return nil
}
