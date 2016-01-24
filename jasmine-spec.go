package jasmine

import "time"

// +testing
func caller(f func()) {
	f()
}
func RunJasmineTests() {
	Describe("testig basic jasmine functions", func() {
		Describe("nesed siute", func() {
			XIt("schould not be called", func() {
				Expect(1).ToBe(2)
			})
		})

		XDescribe("this suite should not be called", func() {
			It("should not be called", func() {
				Expect(1).ToBe(2)
			})
		})
	})

	Describe("testing basic expectations", func() {
		It("test spec ToBe", func() {
			Expect(true).ToBe(true)
			Expect(true).Not.ToBe(false)
		})

		It("test spec ToEqual", func() {
			Expect(1).ToEqual(1)
			Expect(1).Not.ToEqual(2)
		})

		It("test spec ToMatch", func() {
			Expect("foo").ToMatch("foo")
			Expect("foo").Not.ToMatch("bar")
		})

		It("test spec ToBeDefined", func() {
			Expect(1).ToBeDefined()
		})

		It("test spec ToBeNull", func() {
			Expect(nil).ToBeNull()
		})

		It("test spec ToBeTruthy", func() {
			Expect(true).ToBeTruthy()
			Expect(false).Not.ToBeTruthy()
		})

		It("test spec ToBeFalsy", func() {
			Expect(false).ToBeFalsy()
			Expect(true).Not.ToBeFalsy()
		})

		It("test spec ToContain", func() {
			array := []int{1, 2, 3}
			exp := Expect(array)
			exp.ToContain(2)
			exp.Not.ToContain(5)
		})

		It("test spec ToBeLessThan", func() {
			Expect(1).ToBeLessThan(4)
		})

		It("test spec ToBeGreaterThan", func() {
			Expect(4).ToBeGreaterThan(1)
		})
	})

	Describe("testing setup and teardown methods", func() {
		var a = 0

		BeforeEach(func() {
			a = 1
		})

		AfterEach(func() {
			a = 2
		})

		It("test setup and teardown", func() {
			Expect(a).ToBe(1)
		})

	})

	Describe("testing spy like functinallity with ItAsync", func() {
		var ch chan bool

		BeforeEach(func() {
			ch = make(chan bool)
		})
		ItAsync("test toHaveBeenCalled", func(done func()) {
			caller(func() {
				go func() {
					ch <- true
				}()
			})
			go func() {
				b := <-ch
				Expect(b).ToBeTruthy()
				done()
			}()
		})
	})
	Describe("testing sync functions with blocked channels", func() {
		SetDefaultTimeoutInterval(10)
		BeforeEach(func() {
			var c = make(chan bool)
			go func() {
				time.AfterFunc(time.Millisecond*3, func() {
					c <- true
				})
			}()
			<-c
		})
		It("testing It, BeforeEach, AfterEach", func() {
			var c = make(chan bool)
			go func() {
				time.AfterFunc(time.Millisecond*3, func() {
					c <- true
				})
			}()
			<-c
		})
		AfterEach(func() {
			var c = make(chan bool)
			go func() {
				time.AfterFunc(time.Millisecond*3, func() {
					c <- true
				})
			}()
			<-c
		})
	})
	Describe("testing async functions", func() {
		SetDefaultTimeoutInterval(10)
		BeforeEachAsync(func(done func()) {
			var c = make(chan bool)
			go func() {
				time.AfterFunc(time.Millisecond*3, func() {
					c <- true
				})
			}()
			<-c
			done()
		})
		ItAsync("ItAsync with BeforeEachAsync and AfterEachAsync", func(done func()) {
			var c = make(chan bool)
			go func() {
				time.AfterFunc(time.Millisecond*3, func() {
					c <- true
				})
			}()
			<-c
			done()
		})
		AfterEachAsync(func(done func()) {
			var c = make(chan bool)
			go func() {
				time.AfterFunc(time.Millisecond*3, func() {
					c <- true
				})
			}()
			<-c
			done()
		})
	})
	Describe("Spy", func() {
		var spy1 *Spy
		BeforeEach(func() {
			spy1 = CreateSpy("Test", func(arg1 int) int {
				return arg1 + 2
			})
		})
		It("And return value", func() {
			spy1.And.ReturnValue("val1")
			Expect(spy1.Invoke()).ToBe("val1")
		})
		It("And call fake", func() {
			spy1.And.CallFake(func(arg1 int, arg2 int) int {
				return arg1 + arg2
			})
			Expect(spy1.Invoke(1, 2)).ToBe(3)
		})
		It("And callThrough", func() {
			spy1.And.CallThrough()
			Expect(spy1.Invoke(2)).ToBe(4)
		})
		It("And stub", func() {
			spy1.And.Stub()
			Expect(spy1.Invoke(1, 2, 3)).ToBeUndefined()
		})
		It("Calls any()", func() {
			Expect(spy1.Calls.Any()).ToBeFalsy()
			spy1.Invoke()
			Expect(spy1.Calls.Any()).ToBeTruthy()
		})
		It("Calls Count()", func() {
			spy1.Invoke()
			spy1.Invoke()
			Expect(spy1.Calls.Count()).ToBe(2)
		})
		It("Calls ArgsFor()", func() {
			spy1.Invoke(1, 2)
			spy1.Invoke(3, 4)
			Expect(spy1.Calls.ArgsFor(0)[0]).ToBe(1)
			Expect(spy1.Calls.ArgsFor(1)).ToEqual([]interface{}{3, 4})
		})
		It("Calls AllArgs()", func() {
			spy1.Invoke(5, "6")
			spy1.Invoke(7)
			Expect(spy1.Calls.AllArgs()).ToEqual([]interface{}{[]interface{}{5, "6"}, []interface{}{7}})
		})
		It("Calls All()", func() {
			spy1.Invoke(1)
			Expect(spy1.Calls.All()[0].args).ToEqual([]interface{}{1})
		})
		It("Calls MostRecent", func() {
			spy1.Invoke(1)
			spy1.Invoke(2)
			spy1.Invoke(3)
			Expect(spy1.Calls.MostRecent().args).ToEqual([]interface{}{3})
		})
		It("Calls First", func() {
			spy1.Invoke(4)
			spy1.Invoke(2)
			spy1.Invoke(3)
			Expect(spy1.Calls.First().args).ToEqual([]interface{}{4})
		})
		It("Calls First", func() {
			spy1.Invoke(4)
			spy1.Invoke(4)
			Expect(spy1.Calls.Count()).ToBe(2)
			spy1.Calls.Reset()
			spy1.Invoke(2)
			Expect(spy1.Calls.Count()).ToBe(1)
		})
	})
}
