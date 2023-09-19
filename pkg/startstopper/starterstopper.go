package startstopper

type Starter interface {
	Start()
}

type Stopper interface {
	Stop()
}

// StartStopper - Combination of both Starter and Stopper interface.
type StartStopper interface {
	Starter
	Stopper
}
