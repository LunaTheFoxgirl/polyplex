package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type Vector2 = raylib.Vector2

type Camera struct {
	iCam raylib.Camera2D
	Position Vector2
	Origin Vector2
	Rotation float32
	Zoom float32
}

func RegisterCameraType(state *lua.LState) {
	fmt.Println("[Polyplex:raylib]", "Register type: Camera2D...")
	mt := state.NewTypeMetatable("camera2d")
	
	state.SetGlobal("camera2d", mt)
	state.SetGlobal("Camera2D", state.NewFunction(NewCamera))
	state.SetField(mt, "new", state.NewFunction(NewCamera))
	
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), cameraMembers))
}


func NewCamera(state *lua.LState) int {
	cam := &Camera{}
	pos := state.CheckTable(1)
	ori := state.CheckTable(2)
	rot := state.CheckNumber(3)
	zom := state.CheckNumber(4)
	cam.Position = Vector2{float32(pos.RawGetInt(1).(lua.LNumber)), float32(pos.RawGetInt(2).(lua.LNumber))}
	cam.Origin = Vector2{float32(ori.RawGetInt(1).(lua.LNumber)), float32(ori.RawGetInt(2).(lua.LNumber))}
	cam.Rotation = float32(rot)
	cam.Zoom = float32(zom)
	ud := state.NewUserData()
	ud.Value = cam
	state.SetMetatable(ud, state.GetTypeMetatable("camera2d"))
	state.Push(ud)
	return 1
}

func (c *Camera) Update() {
	c.iCam.Offset = c.Position
	c.iCam.Target = c.Origin
	c.iCam.Rotation = c.Rotation
	c.iCam.Zoom = c.Zoom
}

func checkCamera(L *lua.LState) *Camera {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Camera); ok {
		return v
	}
	L.ArgError(1, "Expected Camera!")
	return nil
}

func camUpdate(state *lua.LState) int {
	cam := checkCamera(state)
	if cam == nil {
		return 0
	}
	cam.Update()
	return 1
}


var cameraMembers = map[string]lua.LGFunction {
	"update": camUpdate,
	"position": func(state *lua.LState) int {
		this := checkCamera(state)
		if this == nil {
			return 0
		}
		if state.GetTop() >= 2 {
			this.Position = Vector2{float32(state.CheckNumber(2)), float32(state.CheckNumber(3))}
			return 0
		}
		state.Push(lua.LNumber(this.Position.X))
		state.Push(lua.LNumber(this.Position.Y))
		return 2
	},
	"origin": func(state *lua.LState) int {
		this := checkCamera(state)
		if state.GetTop() >= 2 {
			x := state.CheckNumber(2)
			y := state.CheckNumber(3)
			this.Origin = Vector2{float32(x), float32(y)}
			return 0
		}
		state.Push(lua.LNumber(this.Origin.X))
		state.Push(lua.LNumber(this.Origin.Y))
		return 2
	},
	"rotation": func(state *lua.LState) int {
		this := checkCamera(state)
		
		if state.GetTop() == 2 {
			x := state.CheckNumber(2)
			this.Rotation = float32(x)
			return 0
		}
		state.Push(lua.LNumber(this.Rotation))
		return 1
	},
	"zoom": func(state *lua.LState) int {
		this := checkCamera(state)
		
		if state.GetTop() == 2 {
			x := state.CheckNumber(2)
			this.Zoom = float32(x)
			return 0
		}
		state.Push(lua.LNumber(this.Zoom))
		return 1
	},
}
