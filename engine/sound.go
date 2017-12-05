package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
)

type SoundType int32

const (
	Sound = SoundType(0)
	Music = SoundType(1)
)

type aSound struct {
	Type SoundType
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
	
		
	if t == Sound {
		snd.Sound = s.(raylib.Sound)
	} else if t == Music {
		snd.Music = s.(raylib.Music)
	} else {
		// Something REALLY fucked up.
		return nil
	}
	
	ud.Value = snd
	
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
	"pause": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		
		if t.Type == Music {
			raylib.PauseMusicStream(t.Music)
			return 0
		}
		
		if t.Type == Sound {
			raylib.PauseSound(t.Sound)
			return 0
		}
		return 0
	},
	"pitch": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		if L.GetTop() == 2 {
			pitch := float32(L.CheckNumber(2))
			if t.Type == Music {
				raylib.SetMusicPitch(t.Music, pitch)
			} else if t.Type == Sound {
				raylib.SetSoundPitch(t.Sound, pitch)
			}
		}
		return 0
	},
	"volume": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		if L.GetTop() == 2 {
			pitch := float32(L.CheckNumber(2))
			if t.Type == Music {
				raylib.SetMusicVolume(t.Music, pitch)
			} else if t.Type == Sound {
				raylib.SetSoundVolume(t.Sound, pitch)
			}
		}
		return 0
	},
	"loop_mus": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		if L.GetTop() == 2 {
			count := float32(L.CheckNumber(2))
			if t.Type == Music {
				raylib.SetMusicLoopCount(t.Music, count)
			}
		}
		return 0
	},
	"is_playing": func (L *lua.LState) int {
		t := checkSound(L)
		if t == nil {
			return 0
		}
		playing := false
		if t.Type == Music {
				playing = raylib.IsMusicPlaying(t.Music)
		} else if t.Type == Sound {
				playing = raylib.IsSoundPlaying(t.Sound)
		}
		L.Push(lua.LBool(playing))
		return 1
	},
}