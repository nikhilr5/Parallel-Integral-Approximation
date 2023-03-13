package concurrent
type Runnable interface {
	Run() // Starts the execution of a Runnable
}

// Callable represents a task that will return a value.
type Callable interface {
	Call() interface{} // Starts the execution of a Callable
}

// Future represents the value that is returned after executing a Runnable or Callable task.
type Future interface {
	Get() interface{}
}

// ExecutorService represents a service that can run om Runnable and/or Callable tasks concurrently.
type ExecutorService interface {

	// Submits a task for execution and returns a Future representing that task.
	Submit(task interface{}) Future

	Shutdown()
}