package common

type Thing interface {
	GetName() string                                                           // name will be part of identification
	GetStates() []string                                                       // the states that thing can be
	GetProperties() []string                                                   // the properties name
	GetServices() map[string]func(map[string]interface{}) (interface{}, error) // service function
}
