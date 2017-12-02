package engine


/* DEPRICATED */

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type Texture struct {
	textureImage raylib.Image
	texture raylib.Texture2D
	FilterMode raylib.TextureFilterMode
}

func RegisterTextureType(state *lua.LState) {
	fmt.Println("[Polyplex:raylib]", "Register type: Texture...")
	mt := state.NewTypeMetatable("texture")
	
	state.SetGlobal("texture", mt)
	
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), textureMembers))
}

func NewTexture(state *lua.LState) int {
	tex := &Texture{}
	ud := state.NewUserData()
	ud.Value = tex
	state.SetMetatable(ud, state.GetTypeMetatable("texture"))
	state.Push(ud)
	return 1
}

func checkTexture(L *lua.LState) *Texture {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Texture); ok {
		return v
	}
	L.ArgError(1, "Expected Texture!")
	return nil
}

var textureMembers = map[string]lua.LGFunction {
	"save_img": func (L *lua.LState) int {
		this := checkTexture(L)
		if L.GetTop() == 2 {
			name := string(L.CheckString(2))
			raylib.SaveImageAs(name, this.textureImage)
			L.Push(lua.LBool(true))
			return 1
		}
		L.Push(lua.LBool(false))
		return 1
	},
}
