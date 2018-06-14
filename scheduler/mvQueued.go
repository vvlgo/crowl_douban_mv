package scheduler

type MvQueuedScheduler struct {
	requestChan chan string
	workerChan  chan chan string
}

func (s *MvQueuedScheduler) Submit(r *string) {
	s.requestChan <- *r
}

func (s *MvQueuedScheduler) ConfigureMasterWorkerChan(chan string) {
	panic("implement me")
}

func (s *MvQueuedScheduler) WorkerReady(w chan string) {
	s.workerChan <- w
}


func (s *MvQueuedScheduler) Run() {
	s.requestChan = make(chan string)
	s.workerChan = make(chan chan string)
	go func() {
		var requestQ []string
		var workerQ []chan string
		for {
			var activeRequest string
			var activeWorker chan string
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}

	}()
}

