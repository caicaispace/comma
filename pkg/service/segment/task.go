package segment

import (
	"sync"

	"comma/pkg/library/core/l"
	"comma/pkg/library/setting"
	"comma/pkg/library/util/crontab"
)

type TaskService struct {
	Segmenter *SegmenterService
}

var (
	cron     = crontab.New()
	lastTime = 0
	minTime  = 0
)

var (
	taskService     *TaskService
	taskServiceOnce sync.Once
)

func GetTaskServiceInstance() *TaskService {
	taskServiceOnce.Do(func() {
		taskService = &TaskService{}
		if setting.Server.Env != "dev" {
			taskService.loadDictOnTime()
		}
		seg := GetSegmenterServiceInstance()
		taskService.Segmenter = seg
		seg.LoadDict()
	})
	return taskService
}

func (ts *TaskService) loadDictOnTime() {
	lastTime = GetLastCreateTime()

	var err error
	var job, job2 *crontab.JobModel

	// 增量更新（每分钟）
	job, err = crontab.NewJobModel("00 * * * * *", func() {
		minTime = minTime + 1
		time := GetLastCreateTime()
		if GetLastCreateTime() > lastTime {
			l.Infof("change so load dic")
			ts.Segmenter.LoadDict()
			lastTime = time
			minTime = 0
		}
	})
	if err != nil {
		panic(err.Error())
	}
	err = cron.DynamicRegister("one minute", job)
	if err != nil {
		panic(err.Error())
	}

	// 全量更新（每天）
	job2, err = crontab.NewJobModel("00 00 03 * * *", func() {
		minTime = 0
		l.Infof("24 time load dic")
		ts.Segmenter.LoadDict()
	})
	if err != nil {
		panic(err.Error())
	}
	err = cron.DynamicRegister("one day", job2)
	if err != nil {
		panic(err.Error())
	}
}
