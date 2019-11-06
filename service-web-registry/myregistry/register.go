//将传入的interf中的方法包装成map[string] HttpWR ，供路由绑定使用


package myregistry

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"net/http"
	"reflect"
)
type ControllerMapsType map[string]reflect.Value
type HttpWR  func(w http.ResponseWriter, r *http.Request)
type RfAddr struct {}
type Base struct{
	CrMap ControllerMapsType
	FuncRegistry map[string] HttpWR
	//rfValue rfAddr
}

//type Handle interface{
//	init()
//	TestUserLogin() HttpWR
//	Login() HttpWR
//}

var Register *Base

func init(){
	Register = &Base{}
	Register.FuncRegistry = make(map[string] HttpWR)
	Register.CrMap = make(ControllerMapsType, 0)
}

//注册函数，通过反射将handles中的方法，打包成字典存入 Register.FuncRegistry中 key为对应的方法名，value为对应的方法
func (b* Base)Registery(handles interface{}){

	//创建反射变量，注意这里需要传入ruTest变量的地址；
	//不传入地址就只能反射Routers静态定义的方法
	vf := reflect.ValueOf(handles)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	fmt.Println("NumMethod:", mNum)
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		b.CrMap[mName] = vf.Method(i)
		f:= b.CrMap[mName].Call(nil)
		_, ifOK := b.FuncRegistry[mName]
		if ifOK {
			panic("重复注册方法 -----> " + mName)
		}
		b.FuncRegistry[mName] = f[0].Interface().(HttpWR)
	}
	if len(b.FuncRegistry) == 0{
		return
	}
	fmt.Println("FuncRegistry: ---->", b.FuncRegistry)
}

//外部使用此接口，将url与handle绑定，也可以在外部直接绑定，不使用此方法
func BindUrlHandle(service web.Service, patter , method string){
	f, ok := Register.FuncRegistry[method]
	if !ok {
		panic("绑定的方法未注册！")
	}
	service.HandleFunc(patter,f)
}