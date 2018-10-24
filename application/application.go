package app_helper

import (
	log "log"
	"runtime/debug"
)

//Application is your service
type Application struct {

	//Context
	context *ApplicationContext

	//When application start
	onStart func(*ApplicationContext)

	//When appliction stop
	onStop func(*ApplicationContext)
}

//Setup application context
func (self *Application) SetContext(c ApplicationContext) {
	self.context = &c
}

//Setup application start event
func (self *Application) SetOnStart(c func(*ApplicationContext)) {
	self.onStart = c
}

//Setup application stop event
func (self *Application) SetOnStop(c func(*ApplicationContext)) {
	self.onStop = c
}

//Run application
func (self *Application) Run(c func(*ApplicationContext)) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Application Uncaught Error Occour:", r.(error).Error())
			debug.PrintStack()
		}
	}()
	self.onStart(self.context)
	c(self.context)
	self.onStop(self.context)
}
