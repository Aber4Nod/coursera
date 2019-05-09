package main

import (
	"sync"
	"sort"
	"fmt"
	"strconv"
)

const (
	thrMax = 6
)

func wrapper(fnc job, in, out chan interface{}) {
	fnc(in, out)
	close(out)
}

func ExecutePipeline(freeFlowJobs ...job) {
	var in chan interface{}

	for idx, fnc := range freeFlowJobs {
		out := make(chan interface{}, 1)
		if idx < len(freeFlowJobs) - 1 {
			go wrapper(fnc, in, out)
			in = out
		} else {
			wrapper(fnc, in, out)
		}
	}
}

func SingleHash(in, out chan interface{}) {
	var wgt sync.WaitGroup

	for data := range in {
		var wg sync.WaitGroup
		var tData, tMd5 string

		wgt.Add(1); wg.Add(1)
		data := fmt.Sprintf("%v", data)
		go func(data string) {
			tData = DataSignerCrc32(data)
			wg.Done()
		}(data)
		md5Data := DataSignerMd5(data)
		go func(data string) {
			tMd5 = DataSignerCrc32(data)
			wg.Wait()
			out <- tData + "~" + tMd5; wgt.Done()
		}(md5Data)
	}
	wgt.Wait()
}

func MultiHash(in, out chan interface{}) {
	var wgt sync.WaitGroup

	for data := range in {
		var wg sync.WaitGroup
		var mThr [thrMax]string

		wgt.Add(1)
		data := fmt.Sprintf("%v", data)
		for thr := 0; thr < thrMax; thr++ {
			wg.Add(1)
			go func(thr int, data string) {
				mThr[thr] = DataSignerCrc32(strconv.Itoa(thr) + data)
				wg.Done()
			}(thr, data)
		}
		go func() {
			var rt string
			wg.Wait()
			for thr := 0; thr < thrMax; thr++ {
				rt = rt + mThr[thr]
			}
			out <- rt; wgt.Done()
		}()
	}
	wgt.Wait()
}

func CombineResults(in, out chan interface{}) {
	fn := make([]string, 0)
	var rt string

	for data := range in {
		fn = append(fn, fmt.Sprintf("%v", data))
	}
	sort.Slice(fn, func(i, j int) bool { return fn[i] < fn[j] })
	for idx, data := range fn {
		if idx == 0 {
			rt += data
		} else {
			rt += "_" + data
		}
	}
	out <- rt
}