// Attack must be implemented by a service client.
type Attack interface {

	// Setup should establish the connection to the service
	// It may want to access the config of the runner.
	Setup(c Config) error

	// Do performs one request and is executed in a separate goroutine.
	Do(ctx context.Context) DoResult

	// Teardown can be used to close the connection to the service
	Teardown() error
	
	// Clone should return a fresh new Attack
	// Make sure the new Attack has values for shared struct fields initialized at Setup.
	Clone() Attack
}