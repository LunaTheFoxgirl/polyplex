package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type Sprite struct {
	iTex *raylib.Image
	Tex raylib.Texture2D
	shdr *Shader
	fliph bool
	flipv bool
}

func RegisterSpriteType(state *lua.LState) {
	fmt.Println("[Polyplex:raylib]", "Register type: Sprite...")
	mt := state.NewTypeMetatable("sprite")
	
	state.SetGlobal("sprite", mt)
	
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), spriteMembers))
}


func NewSprite(state *lua.LState, baseTex *raylib.Image) *Sprite {
	tex := &Sprite{iTex: baseTex, Tex: raylib.LoadTextureFromImage(baseTex)}
	ud := state.NewUserData()
	ud.Value = tex
	state.SetMetatable(ud, state.GetTypeMetatable("sprite"))
	state.Push(ud)
	return tex
}

func checkSprite(L *lua.LState) *Sprite {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Sprite); ok {
		return v
	}
	L.ArgError(1, "Expected Sprite!")
	return nil
}

var spriteMembers = map[string]lua.LGFunction {
	"attach_shader": func(L *lua.LState) int {
		spr := checkSprite(L)
		if spr == nil {
			return 0
		}
		if L.GetTop() == 1 {
			spr.shdr = nil
			return 0
		}
		shd := checkShader(L, 2)
		if shd == nil {
			return 0
		}
		spr.shdr = shd
		return 0
	},
	"draw": func(L *lua.LState) int {
		this := checkSprite(L)
		if this == nil {
			return 0
		}
		if this.shdr != nil {
			raylib.BeginShaderMode(this.shdr.shd)
		}
		pos := L.CheckTable(2)
		src := L.CheckTable(3)
		rot := L.CheckNumber(4)
		col := L.CheckTable(5)
		posR := raylib.Rectangle{int32(pos.RawGetInt(1).(lua.LNumber)), int32(pos.RawGetInt(2).(lua.LNumber)), int32(pos.RawGetInt(3).(lua.LNumber)), int32(pos.RawGetInt(4).(lua.LNumber))}
		srcR := raylib.Rectangle{int32(src.RawGetInt(1).(lua.LNumber)), int32(src.RawGetInt(2).(lua.LNumber)), int32(src.RawGetInt(3).(lua.LNumber)), int32(src.RawGetInt(4).(lua.LNumber))}
		if this.fliph {
			srcR = raylib.Rectangle{srcR.X+srcR.Width, srcR.Y, -srcR.Width, srcR.Height}
		}
		if this.flipv {
			srcR = raylib.Rectangle{srcR.X, srcR.Y+srcR.Height, srcR.Width, -srcR.Height}
		}
		colR := raylib.Color{uint8(col.RawGetInt(1).(lua.LNumber)), uint8(col.RawGetInt(2).(lua.LNumber)), uint8(col.RawGetInt(3).(lua.LNumber)), uint8(col.RawGetInt(4).(lua.LNumber))}
		
		raylib.DrawTexturePro(this.Tex, srcR, posR, Vector2{0, 0}, float32(rot), colR)
		//TODO: Draw functionality
		
		if this.shdr != nil {
			raylib.EndShaderMode()
		}
		return 0
	},
	"width": func(L *lua.LState) int {
		this := checkSprite(L)
		if this == nil {
			return 0
		}
		L.Push(lua.LNumber(this.iTex.Width))
		return 1
	},
	"height": func(L *lua.LState) int {
		this := checkSprite(L)
		if this == nil {
			return 0
		}
		L.Push(lua.LNumber(this.iTex.Height))
		return 1
	},
	"flip_h": func(L *lua.LState) int {
		this := checkSprite(L)
		if this == nil {
			return 0
		}
		if L.GetTop() == 2 {
			this.fliph = bool(L.CheckBool(2))
			return 0
		}
		L.Push(lua.LBool(this.fliph))
		return 1
	},
	"flip_v": func(L *lua.LState) int {
		this := checkSprite(L)
		if this == nil {
			return 0
		}
		if L.GetTop() == 2 {
			this.flipv = bool(L.CheckBool(2))
			return 0
		}
		L.Push(lua.LBool(this.flipv))
		return 1
	},
}
