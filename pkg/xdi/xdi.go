package xdi

// dependency injection

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Say() {
	fmt.Println("Hello")
}

func Call(plug any) {
	// fmt.Println(plug)
	FuncAnalyse(plug)
	// typeOf := reflect.TypeOf(plug)
	// fmt.Println("typeOf", typeOf)
	// fmt.Println("kind", typeOf.Kind())
	// fmt.Println("valueOf", reflect.ValueOf(plug))
	// fmt.Println(reflect.TypeOf(plug).Name())
	// name := reflect.TypeOf(plug).Name()
	// valOf := reflect.ValueOf(plug)
	// fmt.Println(valOf)
	// valOf.MethodByName("Say").Call([]reflect.Value{reflect.ValueOf("aman")})
	// fmt.Println(valOf.MethodByName("Version").Call([]reflect.Value{}))
}

type App struct {
	constructs []interface{}
	plugins    map[string]reflect.Value
}

func New() *App {
	return &App{
		plugins:    make(map[string]reflect.Value),
		constructs: make([]interface{}, 0),
	}
}

func (a *App) AddConstruct(c any) {
	// Call(plug)
	// a.constructs[reflect.TypeOf(c).String()] = c
	// fmt.Println(reflect.TypeOf(c).Kind())
	// If user pass other kind e..g struct, string other than Func type
	kind := reflect.TypeOf(c).Kind().String()
	fmt.Println(reflect.TypeOf(c))
	if kind != "func" {
		fmt.Println("wrong constructor passed it should be func type")
		return
	}
	a.constructs = append(a.constructs, c)
}

func (a App) ListConstructs() []interface{} {
	return a.constructs
}

func (a *App) AddPlugin(name string, plug reflect.Value) {
	// Call(plug)
	// a.plugins[reflect.TypeOf(plug).String()] = plug
	a.plugins[name] = plug
}

func (a App) ListPlugins() map[string]reflect.Value {
	return a.plugins
}

func (a *App) RunConstruct(c any) {
	typeOf := reflect.TypeOf(c)
	fmt.Println("typeOf", typeOf)

	numIn := typeOf.NumIn() //Count inbound parameters
	// fmt.Println("numIn", numIn)

	funcArgs := []reflect.Value{}
	for i := 0; i < numIn; i++ {

		inV := typeOf.In(i)
		paramTypeName := inV.String() // demoplug.Config

		// fmt.Println("paramTypeName", paramTypeName)

		for plugName, plug := range a.plugins {
			// fmt.Println("plugName", plugName)
			if plugName == paramTypeName {
				funcArgs = append(funcArgs, plug)
			}
		}
		// fmt.Println(inV.String())
		// in_Kind := inV.Kind() //func
		// fmt.Printf("\nParameter IN: "+strconv.Itoa(i)+"\nKind: %v\nName: %v\n-----------", in_Kind, inV.Name())
	}

	if len(funcArgs) != numIn {
		fmt.Println("Not all the function args provided")
		return
	}

	// name := reflect.TypeOf(c).Name()
	valOf := reflect.ValueOf(c)
	// outs := valOf.Call([]reflect.Value{})
	outs := valOf.Call(funcArgs)
	// fmt.Println(outs)
	for _, o := range outs {
		// fmt.Println(o.Type().String())
		a.AddPlugin(o.Type().String(), o)
		// typeOf := reflect.TypeOf(o)
		// fmt.Println(typeOf.Kind())
		// fmt.Println(typeOf.Name())
		// fmt.Println(typeOf.Elem())

		// val := reflect.Indirect(reflect.ValueOf(o))
		// fmt.Println(val.Field(0).Type().Name())

		// value := reflect.Value(o)
		// fmt.Println(value.Elem())

	}

}

func (a *App) Init() {
	for _, v := range a.constructs {
		a.RunConstruct(v)
	}
}

type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}

// type Pool[T any] struct {
// 	p *sync.Pool
// }

// func NewPool[T any](new func() T) Pool[T] {
// 	return Pool[T]{
// 		p: &sync.Pool{
// 			New: func() interface{} {
// 				return new()
// 			},
// 		},
// 	}
// }

func GetPlugin[T any](app *App, def T) (T, error) {
	// fmt.Println("aa", reflect.TypeOf(def))
	typeof := reflect.TypeOf(def)
	for p, v := range app.plugins {
		// fmt.Println("EEEE", p)
		if p == typeof.String() {
			return v.Interface().(T), nil

		}
	}
	return def, errors.New("not found")
}

func (a *App) GetPlugin1(new any) any {
	return new
}

func FuncAnalyse(m interface{}) {

	//Reflection type of the underlying data of the interface
	x := reflect.TypeOf(m)

	numIn := x.NumIn()   //Count inbound parameters
	numOut := x.NumOut() //Count outbounding parameters

	fmt.Println("Method:", x.String())
	fmt.Println("Variadic:", x.IsVariadic()) // Used (<type> ...) ?
	fmt.Println("Package:", x.PkgPath())

	for i := 0; i < numIn; i++ {

		inV := x.In(i)
		in_Kind := inV.Kind() //func
		fmt.Printf("\nParameter IN: "+strconv.Itoa(i)+"\nKind: %v\nName: %v\n-----------", in_Kind, inV.Name())
	}
	for o := 0; o < numOut; o++ {

		returnV := x.Out(0)
		return_Kind := returnV.Kind()
		fmt.Printf("\nParameter OUT: "+strconv.Itoa(o)+"\nKind: %v\nName: %v\n", return_Kind, returnV.Name())
	}

}

// app.Say()
// app.Call(demoplug.NewConfig)
// app1 := xdi.New()
// 	// app1.AddConstruct(NewConfigMgrConfig)
// 	app1.AddConstruct(configmgr.New)
// 	// app1.AddConstruct(NewInterfaceConfig)
// 	// app1.AddConstruct(cli.New)
// 	// app1.AddConstruct(logger.New)
// 	// app1.AddConstruct(new(dbconn.Config))
// 	// app1.AddConstruct(dbconn.New)
// 	// app1.AddConstruct(NewGormConfig)
// 	// app1.AddConstruct(gormdb.New)
// 	// app1.AddConstruct(NewGormDb)
// 	// app1.AddConstruct(demoplug.NewConfig)
// 	// app1.AddConstruct(demoplug.New)
// 	// fmt.Println(app1.ListConstructs())
// 	app1.Init()
// fmt.Println(app1.ListPlugins())

// p1, err := app.GetPlugin(app1, &configmgr.ConfigMgr{})
// fmt.Println(p1, err)
