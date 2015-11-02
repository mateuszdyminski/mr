package mapreduce

import (
	"container/list"
	"fmt"
	"sync"
)

type WorkerInfo struct {
	address string
	// You can add definitions here.
}

// Clean up all workers by sending a Shutdown RPC to each one of them Collect
// the number of jobs each work has performed.
func (mr *MapReduce) KillWorkers() *list.List {
	l := list.New()
	for _, w := range mr.Workers {
		DPrintf("DoWork: shutdown %s\n", w.address)
		args := &ShutdownArgs{}
		var reply ShutdownReply
		ok := call(w.address, "Worker.Shutdown", args, &reply)
		if ok == false {
			fmt.Printf("DoWork: RPC %s shutdown error\n", w.address)
		} else {
			l.PushBack(reply.Njobs)
		}
	}
	return l
}

func (mr *MapReduce) RunMaster() *list.List {
	fmt.Printf("Run master...\n")

	var wg sync.WaitGroup
	wg.Add(mr.nMap)

	fmt.Printf("Map phase - number of map calls: %d \n", mr.nMap)
	for i := 0; i < mr.nMap; i++ {
		go mr.callMap(i, &wg)
	}
	wg.Wait()

	wg.Add(mr.nReduce)
	fmt.Printf("Reduce phase - number of reduce calls: %d \n", mr.nReduce)
	for i := 0; i < mr.nReduce; i++ {
		go mr.callReduce(i, &wg)
	}
	wg.Wait()

	return mr.KillWorkers()
}

func (mr *MapReduce) callMap(jobId int, wg *sync.WaitGroup) {
	args := DoJobArgs{mr.file, Map, jobId, mr.nReduce}
	reply := DoJobReply{}
	worker := <-mr.WorkersPool
	fmt.Printf("Starting map phase %d...\n", jobId)
	ok := call(worker, "Worker.DoJob", args, &reply)
	if ok {
		wg.Done()
		mr.WorkersPool <- worker
		fmt.Printf("Map phase(%d) success\n", jobId)
	} else {
		fmt.Printf("Map phase(%d) failed. Restarting...\n", jobId)
		go mr.callMap(jobId, wg)
	}
}

func (mr *MapReduce) callReduce(jobId int, wg *sync.WaitGroup) {
	fmt.Printf("Reduce phase - job: %d \n", jobId)
	args := DoJobArgs{mr.file, Reduce, jobId, mr.nMap}
	reply := DoJobReply{}
	worker := <-mr.WorkersPool
	ok := call(worker, "Worker.DoJob", args, &reply)
	if ok {
		wg.Done()
		mr.WorkersPool <- worker
		fmt.Printf("Reduce phase(%d) success\n", jobId)
	} else {
		fmt.Printf("Reduce phase(%d) failed. Restarting...\n", jobId)
		go mr.callReduce(jobId, wg)
	}
}
