package trylock

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{ch: make(chan struct{}, 1)}
}

func (M *Mutex) Lock() {
	M.ch <- struct{}{}
}

func (M *Mutex) Unlock() {
	select {
	case <-M.ch:
		return
	default:
		panic("unlock of unlocked mutex")
	}
}

func (M *Mutex) Trylock() bool {
	select {
	case M.ch <- struct{}{}:
		return true
	default:
		return false
	}
	return false
}

func (M *Mutex) Islocked() bool {
	return len(M.ch) == 1
}
