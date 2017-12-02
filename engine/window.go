package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type Window struct {
	Title string
	Width int32
	Height int32
}

func (this *Window) Init(L *lua.LState) int {
	err := L.CallByParam(lua.P{	Fn: L.GetGlobal("game_init"), NRet: 0, Protect: true})
	if err != nil {
		fmt.Println("[Polyplex:raylib:WARNING]", "game_init is not defined!")
		return 0
	}
	return 1
}

func (this *Window) Update(L *lua.LState) int {
	err := L.CallByParam(lua.P{	Fn: L.GetGlobal("game_update"), NRet: 0, Protect: true})
	if err != nil {
		fmt.Println("[Polyplex:raylib:ERROR]\nTry implementing game_update?\n\n--------Lua error information------\n", err)
		return 0
	}
	return 1
}

func (this *Window) Draw(L *lua.LState) int {
	err := L.CallByParam(lua.P{	Fn: L.GetGlobal("game_draw"), NRet: 0, Protect: true})
	if err != nil {
		fmt.Println("[Polyplex:raylib:ERROR]\nTry implementing game_draw?\n\n--------Lua error information------\n", err)
		return 0
	}
	return 1
}



func RegisterWindowType(state *lua.LState) {
	fmt.Println("[Polyplex:raylib]", "Register type: Window...")
	mt := state.NewTypeMetatable("window")
	
	state.SetGlobal("window", mt)
	state.SetGlobal("Window", state.NewFunction(NewWindow))
	state.SetField(mt, "new", state.NewFunction(NewWindow))
	
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), windowMembers))
}

func NewWindow(state *lua.LState) int {
	win := &Window{Title: state.CheckString(1), Width: int32(state.CheckInt(2)), Height: int32(state.CheckInt(3))}
	ud := state.NewUserData()
	ud.Value = win
	state.SetMetatable(ud, state.GetTypeMetatable("window"))
	state.Push(ud)
	return 1
}

func checkWindow(L *lua.LState) *Window {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Window); ok {
		return v
	}
	L.ArgError(1, "Expected Window!")
	return nil
}

func startGame(L *lua.LState) int {
	this := checkWindow(L)
	if this == nil {
		return 0
	}
	raylib.InitWindow(this.Width, this.Height, this.Title)
	raylib.ShowLogo()
	this.Init(L)
	for !raylib.WindowShouldClose() {
		err := this.Update(L)
		if err == 0 {
			raylib.CloseWindow()
			return 0
		}
		
		raylib.BeginDrawing()
		this.Draw(L)
		if err == 0 {
			raylib.CloseWindow()
			return 0
		}
		raylib.EndDrawing()
	}
	err := L.CallByParam(lua.P{	Fn: L.GetGlobal("game_cleanup"), NRet: 0, Protect: true})
	if err != nil {
		fmt.Println("[Polyplex:raylib:WARNING]", "game_cleanup is not defined!\n\nGame cleanup is important!")
		return 0
	}
	return 1

	return 1
}

func setResizable(state *lua.LState) int {
	if state.GetTop() == 2 {
		raylib.SetWindowResizable(state.CheckBool(2))
		return 1
	}
	state.ArgError(1, "Expected window state boolean!")
	return 0
}

func getWidth(state *lua.LState) int {
	this := checkWindow(state)
	if this == nil {
		return 0
	}
	state.Push(lua.LNumber(this.Width))
	return 1
}

func getPosition(state *lua.LState) int {
	this := checkWindow(state)
	if this == nil {
		return 0
	}
	i := raylib.GetGameWindowInfo()
	state.Push(lua.LNumber(i.Size.Position.X))
	state.Push(lua.LNumber(i.Size.Position.Y))
	return 2
}


func getHeight(state *lua.LState) int {
	this := checkWindow(state)
	if this == nil {
		return 0
	}
	state.Push(lua.LNumber(this.Height))
	return 1
}

func getSetTitle(state *lua.LState) int {
	this := checkWindow(state)
	if this == nil {
		return 0
	}
	
	if state.GetTop() == 2 {
		this.Title = state.CheckString(2)
		return 0
	}
	state.Push(lua.LString(this.Title))
	return 1
}

func targetFPS(state *lua.LState) int {
	if state.GetTop() == 2 {
		raylib.SetTargetFPS(int32(state.CheckInt(2)))
		return 0
	}
	return 1
}

var windowMembers = map[string]lua.LGFunction {
	"set_resizable": setResizable,
	"start_game": startGame,
	"width": getWidth,
	"height": getHeight,
	"position": getPosition,
	"title": getSetTitle,
	"target_fps": targetFPS,
}
