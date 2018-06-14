package scheduler

type SimpleScheduler struct {
	workChan chan string
}

func (s *SimpleScheduler) Submit(url *string){
	go func() {
		s.workChan <- *url
	}()

}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan string){
	s.workChan =in
}
