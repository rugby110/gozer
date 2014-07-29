package mesos

type stateErrorId int

const (
	errorNone stateErrorId = iota
	errorNotInitialized
	errorNotConnected
	errorNotReady
	//...
)

func stateError(d *Driver) stateFn {
	// Handle any type of error state
	d.log.Info.Println("STATE: Error, driver =", d)
	return stateStop
}
