package app_helper

type fContext struct {
}

//Init your custom context obj, such as db connections, log configs etc by implements ApplicationContext interface
func (self fContext) Init() {

}

func (self fContext) MyOwnMethod() {

}

func demo() {
	context := fContext{}
	context.Init()
	app := NewApplication()
	app.SetContext(context)
	app.SetOnStart(onStart)
	app.SetOnStop(onStop)
	app.Run(run)
}

func onStart(c ApplicationContext) {

}

func onStop(c ApplicationContext) {
	c.(fContext).MyOwnMethod()
}

func run(c ApplicationContext) {

}
