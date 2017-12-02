package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type SoundType int32

const (
	Wave = SoundType(0)
	Sound = SoundType(1)
	Music = SoundType(2)
)

type aSound struct {
	Type SoundType
	Wave raylib.Wave
	Sound raylib.Sound
	Music raylib.Music
}

func RegisterSoundType(state *lua.LState) {
	fmt.Println("[Polyplex:raylib]", "Register type: Sound...")
	mt := state.NewTypeMetatable("sound")
	
	state.SetGlobal("sound", mt)
	
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), soundMembers))
}


func NewSound(L *lua.LState, t SoundType, s interface{}) *aSound {
	snd := &aSound{ Type:t }
	ud := L.NewUserData()
	ud.Value = snd
	
	if t == Wave {
		snd.Wave = s.(raylib.Wave)
	} else if t == Sound {
		snd.Sound = s.(raylib.Sound)
	} else if t == Music {
		snd.Music = s.(raylib.Music)
	} else {
		// Something REALLY fucked up.
		return nil
	}
	
	L.SetMetatable(ud, L.GetTypeMetatable("sound"))
	L.Push(ud)
	return snd
}

func checkSound(L *lua.LState) *aSound {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*aSound); ok {
		return v
	}
	L.ArgError(1, "Expected Sound!")
	return nil
}

var soundMembers = map[string]lua.LGFunction {
	"play": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		
		if t.Type == Wave {
			return 0
		}
		
		if t.Type == Music {
			raylib.PlayMusicStream(t.Music)
			return 0
		}
		
		if t.Type == Sound {
			raylib.PlaySound(t.Sound)
			return 0
		}
		
		return 0
	},
	"stop": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		
		if t.Type == Wave {
			return 0
		}
		
		if t.Type == Music {
			raylib.StopMusicStream(t.Music)
			return 0
		}
		
		if t.Type == Sound {
			raylib.StopSound(t.Sound)
			return 0
		}
		return 0
	},
}