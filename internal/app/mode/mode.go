package mode

// Mode defines the behavior required for different game running execution strategies,
// such as interactive or non-interactive workflows.
type Mode interface {
	Run() error
}
