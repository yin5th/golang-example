package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Number int
	Id     int
}

type Result struct {
	job *Job
	sum int
}

func calc(job *Job, result chan *Result) {
	var sum int
	number := job.Number
	for number != 0 {
		// 获取个位数
		tmp := number % 10
		sum += tmp
		number /= 10
	}
	r := &Result{
		job: job,
		sum: sum,
	}
	result <- r
}

func Worker(jobChan chan *Job, resultChan chan *Result) {
	for job := range jobChan {
		go calc(job, resultChan)
	}
}

func startWokePool(num int, jobChan chan *Job, resultChan chan *Result) {
	//num线程数量
	for i := 0; i < num; i++ {
		go Worker(jobChan, resultChan)
	}
}

func printResult(resultChan chan *Result) {
	for result := range resultChan {
		fmt.Printf("Job ID:%v Number:%d sum:%d\n", result.job.Id, result.job.Number, result.sum)
	}
}

func main() {
	jobChan := make(chan *Job, 100)
	resultChan := make(chan *Result, 100)

	// 开启工作池 线程数128 任务队列jobChan 结果保存在resultChan
	startWokePool(128, jobChan, resultChan)
	// 打印结果
	go printResult(resultChan)

	var id int
	for {
		id++
		number := rand.Int()
		job := &Job{
			Id:     id,
			Number: number,
		}

		jobChan <- job
	}
}
