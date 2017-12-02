package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type Shader struct {
	shd raylib.Shader
}

func RegisterShaderType(state *lua.LState) {
	fmt.Println("[Polyplex:raylib]", "Register type: Shader...")
	mt := state.NewTypeMetatable("shader")
	
	state.SetGlobal("shader", mt)
	
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), shaderMembers))
}


func NewShader(state *lua.LState, shd raylib.Shader) *Shader {
	shdr := &Shader{shd: shd}
	ud := state.NewUserData()
	ud.Value = shdr
	state.SetMetatable(ud, state.GetTypeMetatable("shader"))
	state.Push(ud)
	return shdr
}

func checkShader(L *lua.LState, where int) *Shader {
	ud := L.CheckUserData(where)
	if v, ok := ud.Value.(*Shader); ok {
		return v
	}
	L.ArgError(1, "Expected Shader!")
	return nil
}

var shaderMembers = map[string]lua.LGFunction {
	"begin_pass": func(L *lua.LState) int {
		this := checkShader(L, 1)
		if this == nil {
			return 0
		}
		raylib.BeginShaderMode(this.shd)
		return 0
	},
	"end_pass": func(L *lua.LState) int {
		raylib.EndShaderMode()
		return 0
	},
	"set_uniform": func(L *lua.LState) int {
		this := checkShader(L, 1)
		if this == nil {
			return 0
		}
		loc := raylib.GetShaderLocation(this.shd, string(L.CheckString(2)))
		pos := L.CheckTable(3)
		val := []float32{}
		for i := 0; i < pos.Len(); i++ {
			val = append(val, float32(pos.RawGetInt(i+1).(lua.LNumber)))
		}
		raylib.SetShaderValue(this.shd, loc, val, int32(len(val)))
		return 0
	},
}