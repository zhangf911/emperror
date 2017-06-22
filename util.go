package emperror

// HandleRecover recovers from a panic and handles the error.
//
// 		go emperror.HandleRecover(errorHandler)
func HandleRecover(handler Handler) {
	err := Recover(recover())
	if err != nil {
		handler.Handle(err)
	}
}

// HandleIfErr handles an error whenever it occures.
func HandleIfErr(handler Handler, err error) {
	if err != nil {
		handler.Handle(err)
	}
}
