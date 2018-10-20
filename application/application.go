package app_helper

//Application is your service
type Application struct {

	//Context
	Context ApplicationContext

	//When application start
	OnStart func()

	//When appliction stop
	OnStop func()
}
