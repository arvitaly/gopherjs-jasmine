package jasmine

import "github.com/gopherjs/gopherjs/js"

// Describe discribes a test suit containing multiple specs.
// The specs are run inside the callback function fn.
func Describe(name string, fn func()) {
	js.Global.Call("describe", name, fn)
}

// XDescribe is a canncled or commented out suit which will not be run.
func XDescribe(name string, fn func()) {
	js.Global.Call("xdescribe", name, fn)
}

//For callback in goroutines
func callAsyncInGo(fn func(func())) func(done func()) {
	return func(done func()) {
		go func() {
			fn(done)
		}()
	}
}
func callSyncInGo(fn func()) func(done func()) {
	return func(done func()) {
		go func() {
			defer func() {
				r := recover()
				if r != nil {
					Fail(r.(*js.Error).String())
					done()
				}
			}()
			fn()
			done()
		}()
	}
}

// It is a spec to be tested with Expectations.
func It(behavior string, fn func()) {
	js.Global.Call("it", behavior, createFuncWithVisibleDone(callSyncInGo(fn)))
}

// XIt is a cancled or commented out spec which will not be run.
func XIt(behavior string, fn func()) {
	js.Global.Call("xit", behavior, createFuncWithVisibleDone(callSyncInGo(fn)))
}

func ItAsync(behavior string, fn func(func())) {
	js.Global.Call("it", behavior, createFuncWithVisibleDone(callAsyncInGo(fn)))
}

func XitAsync(behavior string, fn func(func())) {

	js.Global.Call("xit", behavior, createFuncWithVisibleDone(callAsyncInGo(fn)))
}

func BeforeEach(fn func()) {
	js.Global.Call("beforeEach", createFuncWithVisibleDone(callSyncInGo(fn)))
}

func BeforeEachAsync(fn func(func())) {
	js.Global.Call("beforeEach", createFuncWithVisibleDone(callAsyncInGo(fn)))
}

func BeforeAllAsync(fn func(func())) {
	js.Global.Call("beforeAll", createFuncWithVisibleDone(callAsyncInGo(fn)))
}

func AfterEach(fn func()) {
	js.Global.Call("afterEach", createFuncWithVisibleDone(callSyncInGo(fn)))
}

func AfterEachAsync(fn func(func())) {
	js.Global.Call("afterEach", createFuncWithVisibleDone(callAsyncInGo(fn)))
}
func AfterAllAsync(fn func(func())) {
	js.Global.Call("afterAll", createFuncWithVisibleDone(callAsyncInGo(fn)))
}

func Expect(value interface{}) *Expectation {
	return &Expectation{o: js.Global.Call("expect", value)}
}

func SetDefaultTimeoutInterval(interval int) {
	js.Global.Get("jasmine").Set("DEFAULT_TIMEOUT_INTERVAL", interval)
}

func Fail(message string) {
	js.Global.Call("fail", message)
}

func Run(calls func()) bool {
	calls()
	return true
}

func createFuncWithVisibleDone(fn func(func())) *js.Object {
	return js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		return js.Global.Get("eval").Invoke("var a = function(cb){ return function(done){ return cb(done); }; }; a")
	}).Invoke().Invoke(fn)
}
