package command

import "context"

//Bus defines the expected behavior from a command bus.
type Bus interface {
	//Dispatch is the used to dispactch new commands
	Dispatch(context.Context, Command) error
	//Register is the method used to register a new command handler.
	Register(Type, Handler)
}

//Type represents an application command type.
type Type string

//Comand represent an application command.
type Command interface {
	Type() Type
}

//Handler defines the expected behavior from
type Handler interface {
	Handle(context.Context, Command) error
}
