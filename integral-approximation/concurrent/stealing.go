package concurrent
import (
	"sync"
)
func NewWorkStealingExecutor(capacity int,  wg *sync.WaitGroup) ExecutorService {
	return NewWorkBalancingExecutor(capacity, 0, wg)
 }