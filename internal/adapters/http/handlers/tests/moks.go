package tests

type testLogger struct{}

func (tl *testLogger) Info(msg string)  {}
func (tl *testLogger) Warn(msg string)  {}
func (tl *testLogger) Error(msg string) {}
