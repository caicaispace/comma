// 测试sego并行分词速度

package segment_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"time"

	"comma/pkg/library/db"

	"github.com/caicaispace/gohelper/setting"

	participleService "comma/pkg/service/segment"
)

var (
	segmenter  = participleService.GetSegmenterServiceInstance()
	numThreads = runtime.NumCPU()
	task       = make(chan []byte, numThreads*40)
	done       = make(chan bool, numThreads)
	numRuns    = 50
)

func worker() {
	for line := range task {
		segmenter.Segment(line)
	}
	done <- true
}

func Test_Goroutines(t *testing.T) {
	// 将线程数设置为CPU数
	runtime.GOMAXPROCS(numThreads)

	// 载入词典
	db.New(&setting.DBSetting{
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		DbName:   "comma",
	})
	// sego.LoadDb("root:phpcj@tcp(localhost:3306)/es_help?charset=utf8")
	segmenter.LoadDict()

	// 打开将要分词的文件
	file, err := os.Open("../testdata/bailuyuan.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 逐行读入
	scanner := bufio.NewScanner(file)
	size := 0
	var lines [][]byte
	for scanner.Scan() {
		var text string
		fmt.Sscanf(scanner.Text(), "%s", &text)
		content := []byte(text)
		size += len(content)
		lines = append(lines, content)
	}

	// 启动工作线程
	for i := 0; i < numThreads; i++ {
		go worker()
	}
	log.Print("开始分词")

	// 记录时间
	t0 := time.Now()

	// 并行分词
	for i := 0; i < numRuns; i++ {
		for _, l := range lines {
			task <- l
		}
	}
	close(task)

	// 确保分词完成
	for i := 0; i < numThreads; i++ {
		<-done
	}

	// 记录时间并计算分词速度
	t1 := time.Now()
	log.Printf("分词花费时间 %v", t1.Sub(t0))
	log.Printf("分词速度 %f MB/s", float64(size*numRuns)/t1.Sub(t0).Seconds()/(1024*1024))
}
