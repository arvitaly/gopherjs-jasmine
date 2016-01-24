package jasmine

import "github.com/gopherjs/gopherjs/js"

func CreateSpy(name string, originalFn interface{}) *Spy {
	return &Spy{Object: js.Global.Get("jasmine").Call("createSpy", name, originalFn)}
}

type Spy struct {
	*js.Object
	And            SpyAnd         `js:"and"`
	Calls          Calls          `js:"calls"`
	MostRecentCall MostRecentCall `js:"mostRecentCall"`
	ArgsForCall    []*js.Object   `js:"argsForCall"`
	WasCalled      bool           `js:"wasCalled"`
}
type MostRecentCall struct {
	*js.Object
	Args []*js.Object `js:"args"`
}

type SpyAnd struct {
	*js.Object
}

/** By chaining the spy with and.callThrough, the spy will still track all calls to it but in addition it will delegate to the actual implementation. */
func (a SpyAnd) CallThrough() Spy {
	return Spy{Object: a.Call("callThrough")}
}

/** By chaining the spy with and.returnValue, all calls to the function will return a specific value. */
func (a SpyAnd) ReturnValue(val interface{}) {
	a.Call("returnValue", val)
}

/** By chaining the spy with and.callFake, all calls to the spy will delegate to the supplied function. */
func (a SpyAnd) CallFake(fn interface{}) Spy {
	return Spy{Object: a.Call("callFake", fn)}
}

/** By chaining the spy with and.throwError, all calls to the spy will throw the specified value. */
func (a SpyAnd) ThrowError(msg string) {
	a.Call("throwError", msg)
}

/** When a calling strategy is used for a spy, the original stubbing behavior can be returned at any time with and.stub. */
func (a SpyAnd) Stub() Spy {
	return Spy{Object: a.Call("stub")}
}

type Calls struct {
	*js.Object
}

/** By chaining the spy with calls.any(), will return false if the spy has not been called at all, and then true once at least one call happens. **/
func (c Calls) Any() bool {
	return c.Call("any").Bool()
}

/** By chaining the spy with calls.count(), will return the number of times the spy was called **/
func (c Calls) Count() int {
	return c.Call("count").Int()
}

/** By chaining the spy with calls.argsFor(), will return the arguments passed to call number index **/
func (c Calls) ArgsFor(index int) []*js.Object {
	var source = c.Call("argsFor", index)
	var dest []*js.Object
	for i := 0; i < source.Length(); i++ {
		dest = append(dest, source.Index(i))
	}
	return dest
}

/** By chaining the spy with calls.allArgs(), will return the arguments to all calls **/
func (c Calls) AllArgs() []*js.Object {
	var source = c.Call("allArgs")
	var dest []*js.Object
	for i := 0; i < source.Length(); i++ {
		dest = append(dest, source.Index(i))
	}
	return dest
}

/** By chaining the spy with calls.all(), will return the context (the this) and arguments passed all calls **/
func (c Calls) All() []CallInfo {
	var source = c.Call("all")
	var dest []CallInfo
	for i := 0; i < source.Length(); i++ {
		dest = append(dest, CallInfo{o: source.Index(i)})
	}
	return dest
}

/** By chaining the spy with calls.mostRecent(), will return the context (the this) and arguments for the most recent call **/
func (c Calls) MostRecent() CallInfo {
	return CallInfo{o: c.Call("mostRecent")}
}

/** By chaining the spy with calls.first(), will return the context (the this) and arguments for the first call **/
func (c Calls) First() CallInfo {
	return CallInfo{o: c.Call("first")}
}

/** By chaining the spy with calls.reset(), will clears all tracking for a spy **/
func (c Calls) Reset() {
	c.Call("reset")
}

type CallInfo struct {
	o *js.Object
	/** The context (the this) for the call */
	Object *js.Object `js:"object"`
	/** All arguments passed to the call */
	args []interface{} `js:"args"`
}
