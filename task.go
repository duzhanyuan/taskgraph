package meritop

// Task is a logic repersentation of a computing unit.
// Each task contain at least one Node.
// Each task has exact one master Node and might have multiple salve Nodes.
type Task interface {
	// This is useful to bring the task up to speed from scratch or if it recovers.
	Init(taskID uint64, framework Framework, config Config)

	// Task need to finish up for exit, last chance to save work?
	Exit()

	// These are called by framework implementation so that task implementation can
	// reacts to parent or children restart.
	ParentRestart(parentID uint64)
	ChildRestart(childID uint64)

	ParentDie(parentID uint64)
	ChildDie(childID uint64)

	// Ideally, we should also have the following:
	ParentReady(parentID uint64, data []byte)
	ChildReady(childID uint64, data []byte)

	// This give the task an opportunity to cleanup and regroup.
	SetEpoch(epochID uint64)
}

// Backupable is an interface that task need to implement if they want to have
// hot standby copy. This is another can of beans.
type Backupable interface {
	// Some hooks that need for master slave etc.
	BecameMaster()
	BecameBackup()
}
