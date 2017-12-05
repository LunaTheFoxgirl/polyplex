package engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/raylib-go/raylib"
	"fmt"
	"time"
)

func ModuleCore(L *lua.LState) int {
	exports := map[string]lua.LGFunction{
	    "begin_camera": func (L *lua.LState) int {
	    	cam := checkCamera(L)
	    	if cam == nil {
	    		return 0
	    	}
	    	raylib.Begin2dMode(cam.iCam)
		    return 1	
	    },
	    "end_camera": func(L *lua.LState) int {
	    	raylib.End2dMode()
	    	return 1
	    },
	    "exit": func(L *lua.LState) int {
	    	raylib.CloseWindow()
	    	return 0
	    },
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func ModuleInput(L *lua.LState) int {
	exports := map[string]lua.LGFunction{
	    "is_key_down": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	key := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsKeyDown(key)))
	    	return 1
	    },
	    "is_key_pressed": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	key := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsKeyPressed(key)))
	    	return 1
	    },
	    "is_key_released": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	key := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsKeyReleased(key)))
	    	return 1
	    },
	    "is_key_up": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	key := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsKeyUp(key)))
	    	return 1
	    },
	    
	    //Mouse
	    "is_mouse_down": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	button := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsMouseButtonDown(button)))
	    	return 1
	    },
	    "is_mouse_pressed": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	button := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsMouseButtonPressed(button)))
	    	return 1
	    },
	    "is_mouse_released": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	button := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsMouseButtonReleased(button)))
	    	return 1
	    },
	    "is_mouse_up": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		L.Push(lua.LBool(false))
	    		return 1
	    	}
	    	button := int32(L.CheckNumber(1))
	    	L.Push(lua.LBool(raylib.IsMouseButtonUp(button)))
	    	return 1
	    },
	    "mouse_position": func(L *lua.LState) int {
			pos := raylib.GetMousePosition()
	    	L.Push(lua.LNumber(pos.X))
	    	L.Push(lua.LNumber(pos.Y))
	    	return 2
	    },
	    "mouse_scroll": func(L *lua.LState) int {
			pos := raylib.GetMouseWheelMove()
	    	L.Push(lua.LNumber(pos))
	    	return 1
	    },
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	// TEXT KEYS
	L.SetField(mod, "key_a", lua.LNumber(raylib.KeyA))
	L.SetField(mod, "key_b", lua.LNumber(raylib.KeyB))
	L.SetField(mod, "key_c", lua.LNumber(raylib.KeyC))
	L.SetField(mod, "key_d", lua.LNumber(raylib.KeyD))
	L.SetField(mod, "key_e", lua.LNumber(raylib.KeyE))
	L.SetField(mod, "key_f", lua.LNumber(raylib.KeyF))
	L.SetField(mod, "key_g", lua.LNumber(raylib.KeyG))
	L.SetField(mod, "key_h", lua.LNumber(raylib.KeyH))
	L.SetField(mod, "key_i", lua.LNumber(raylib.KeyI))
	L.SetField(mod, "key_j", lua.LNumber(raylib.KeyJ))
	L.SetField(mod, "key_k", lua.LNumber(raylib.KeyK))
	L.SetField(mod, "key_l", lua.LNumber(raylib.KeyL))
	L.SetField(mod, "key_m", lua.LNumber(raylib.KeyM))
	L.SetField(mod, "key_n", lua.LNumber(raylib.KeyN))
	L.SetField(mod, "key_o", lua.LNumber(raylib.KeyO))
	L.SetField(mod, "key_p", lua.LNumber(raylib.KeyP))
	L.SetField(mod, "key_q", lua.LNumber(raylib.KeyQ))
	L.SetField(mod, "key_r", lua.LNumber(raylib.KeyR))
	L.SetField(mod, "key_s", lua.LNumber(raylib.KeyS))
	L.SetField(mod, "key_t", lua.LNumber(raylib.KeyT))
	L.SetField(mod, "key_u", lua.LNumber(raylib.KeyU))
	L.SetField(mod, "key_v", lua.LNumber(raylib.KeyV))
	L.SetField(mod, "key_w", lua.LNumber(raylib.KeyW))
	L.SetField(mod, "key_x", lua.LNumber(raylib.KeyX))
	L.SetField(mod, "key_y", lua.LNumber(raylib.KeyY))
	L.SetField(mod, "key_z", lua.LNumber(raylib.KeyZ))
	
	//NUMERIC KEYS
	L.SetField(mod, "key_1", lua.LNumber(raylib.KeyOne))
	L.SetField(mod, "key_2", lua.LNumber(raylib.KeyTwo))
	L.SetField(mod, "key_3", lua.LNumber(raylib.KeyThree))
	L.SetField(mod, "key_4", lua.LNumber(raylib.KeyFour))
	L.SetField(mod, "key_5", lua.LNumber(raylib.KeyFive))
	L.SetField(mod, "key_6", lua.LNumber(raylib.KeySix))
	L.SetField(mod, "key_7", lua.LNumber(raylib.KeySeven))
	L.SetField(mod, "key_8", lua.LNumber(raylib.KeyEight))
	L.SetField(mod, "key_9", lua.LNumber(raylib.KeyNine))
	L.SetField(mod, "key_0", lua.LNumber(raylib.KeyZero))
	
	//Mod keys
	L.SetField(mod, "key_lctrl", lua.LNumber(raylib.KeyLeftControl))
	L.SetField(mod, "key_lalt", lua.LNumber(raylib.KeyLeftAlt))
	L.SetField(mod, "key_lshift", lua.LNumber(raylib.KeyLeftShift))
	L.SetField(mod, "key_rctrl", lua.LNumber(raylib.KeyRightControl))
	L.SetField(mod, "key_ralt", lua.LNumber(raylib.KeyRightAlt))
	L.SetField(mod, "key_rshift", lua.LNumber(raylib.KeyRightShift))
	
	//Action Keys
	L.SetField(mod, "key_esc", lua.LNumber(raylib.KeyEscape))
	L.SetField(mod, "key_backspace", lua.LNumber(raylib.KeyBackspace))
	L.SetField(mod, "key_space", lua.LNumber(raylib.KeySpace))
	L.SetField(mod, "key_enter", lua.LNumber(raylib.KeyEnter))
	L.SetField(mod, "key_return", lua.LNumber(raylib.KeyEnter))
	
	
	L.SetField(mod, "key_F1", lua.LNumber(raylib.KeyF1))
	L.SetField(mod, "key_F2", lua.LNumber(raylib.KeyF2))
	L.SetField(mod, "key_F3", lua.LNumber(raylib.KeyF3))
	L.SetField(mod, "key_F4", lua.LNumber(raylib.KeyF4))
	L.SetField(mod, "key_F5", lua.LNumber(raylib.KeyF5))
	L.SetField(mod, "key_F6", lua.LNumber(raylib.KeyF6))
	L.SetField(mod, "key_F7", lua.LNumber(raylib.KeyF7))
	L.SetField(mod, "key_F8", lua.LNumber(raylib.KeyF8))
	L.SetField(mod, "key_F9", lua.LNumber(raylib.KeyF9))
	L.SetField(mod, "key_F10", lua.LNumber(raylib.KeyF10))
	L.SetField(mod, "key_F11", lua.LNumber(raylib.KeyF11))
	L.SetField(mod, "key_F12", lua.LNumber(raylib.KeyF12))
	
	//Special keys
	L.SetField(mod, "key_vol_up", lua.LNumber(raylib.KeyVolumeUp))
	L.SetField(mod, "key_vol_up", lua.LNumber(raylib.KeyVolumeUp))
	
	//Mouse Buttons
	L.SetField(mod, "mouse_left", lua.LNumber(raylib.MouseLeftButton))
	L.SetField(mod, "mouse_middle", lua.LNumber(raylib.MouseMiddleButton))
	L.SetField(mod, "mouse_right", lua.LNumber(raylib.MouseRightButton))
	L.Push(mod)
	return 1
}

func ModuleTranslation(L *lua.LState) int {
	exports := map[string]lua.LGFunction{
	    "translate": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		return 0
	    	}
	    	t := L.CheckString(1)
	    	if CURRENT_TRANSLATION == nil {
	    		L.Push(lua.LString(t))
	    		return 1
	    	}
	    	L.Push(lua.LString(CURRENT_TRANSLATION.Translate(string(t))))
	    	return 1
	    },
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func ModuleUtilities(L *lua.LState) int {
	exports := map[string]lua.LGFunction{
		"measure_string": func(L *lua.LState) int {
			t := string(L.CheckString(1))
	    	s := int32(L.CheckNumber(2))
			w := raylib.MeasureText(t, s)
			L.Push(lua.LNumber(s))
			L.Push(lua.LNumber(w))
			return 2
		},
	    "draw_string": func(L *lua.LState) int {
	    	if L.GetTop() < 2 {
	    		return 0
	    	}
	    	t := L.CheckString(1)
	    	x := L.CheckNumber(2)
	    	y := L.CheckNumber(3)
	    	z := L.CheckNumber(4)
	    	c := L.CheckTable(5)
	    	r := c.RawGetInt(1).(lua.LNumber)
	    	g := c.RawGetInt(2).(lua.LNumber)
	    	b := c.RawGetInt(3).(lua.LNumber)
	    	a := c.RawGetInt(4).(lua.LNumber)
	    	raylib.DrawText(string(t), int32(x), int32(y), int32(z), raylib.Color{uint8(r), uint8(g), uint8(b), uint8(a)})
	    	return 0
	    },
	    "clear_color": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		return 0
	    	}
	    	c := L.CheckTable(1)
	    	r := c.RawGetInt(1).(lua.LNumber)
	    	g := c.RawGetInt(2).(lua.LNumber)
	    	b := c.RawGetInt(3).(lua.LNumber)
	    	a := c.RawGetInt(4).(lua.LNumber)
			raylib.ClearBackground(raylib.Color{uint8(r), uint8(g), uint8(b), uint8(a)})
	    	return 0
	    },
	    "draw_fps": func(L *lua.LState) int {
			pos := L.CheckTable(1)
	    	x := pos.RawGetInt(1)
	    	y := pos.RawGetInt(2)
	    	raylib.DrawFPS(int32(x.(lua.LNumber)), int32(y.(lua.LNumber)))
	    	return 1
	    },
	    "draw_debug_sq": func(L *lua.LState) int {
	    	if L.GetTop() != 2 {
	    		return 0
	    	}
	    	p := L.CheckTable(1)
	    	x := int32(p.RawGetInt(1).(lua.LNumber))
	    	y := int32(p.RawGetInt(2).(lua.LNumber))
	    	w := int32(p.RawGetInt(3).(lua.LNumber))
	    	h := int32(p.RawGetInt(4).(lua.LNumber))
	    	c := L.CheckTable(2)
	    	r := uint8(c.RawGetInt(1).(lua.LNumber))
	    	g := uint8(c.RawGetInt(2).(lua.LNumber))
	    	b := uint8(c.RawGetInt(3).(lua.LNumber))
	    	a := uint8(c.RawGetInt(4).(lua.LNumber))
	    	raylib.DrawRectangleLines(x, y, w, h, raylib.Color{r,g,b,a})
	    	return 0
	    },
	    "set_raylib_logging": func(L *lua.LState) int {
	    	raylib.SetLogging(int(L.CheckNumber(1)))
	    	raylib.SetLoggingGo(int(L.CheckNumber(1)))
	    	return 0
	    },
		"get_time_ms": func(L *lua.LState) int {
			L.Push(lua.LNumber(time.Now().UnixNano() / int64(time.Millisecond)))
	    	return 1
	    },
		"master_volume": func(L *lua.LState) int {
			if L.GetTop() != 1 {
	    		return 0
			}
	    	v := L.CheckNumber(1)
	    	raylib.SetMasterVolume(float32(v))
	    	return 1
	    },
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "log_info", lua.LNumber(raylib.LogInfo))
	L.SetField(mod, "log_warning", lua.LNumber(raylib.LogWarning))
	L.SetField(mod, "log_error", lua.LNumber(raylib.LogError))
	L.SetField(mod, "log_debug", lua.LNumber(raylib.LogDebug))
	
	L.Push(mod)
	return 1
}

func ModuleResources(L *lua.LState) int {
	exports := map[string]lua.LGFunction{
	    "load_sprite": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		return 0
	    	}
	    	t := L.CheckString(1)
	    	CONTENT.LoadSprite(L, string(t))
	    	return 1
	    },
	    "load_shader": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
		    	CONTENT.LoadDefaultShader(L)
	    		return 0
	    	}
	    	t := L.CheckString(1)
	    	sh := CONTENT.LoadShader(L, string(t))
	    	if sh == nil {
	    		fmt.Println("[Polyplex:raylib:ERROR]", "Unkown error while loading shader", t)
		    	CONTENT.LoadDefaultShader(L)
	    		return 1
	    	}
	    	return 1
	    },
	    "load_sound": func(L *lua.LState) int {
	    	if L.GetTop() != 2 {
	    		return 0
	    	}
	    	t := L.CheckNumber(1)
	    	t2 := L.CheckString(2)
	    	
	    	CONTENT.LoadSound(L, SoundType(t), string(t2))
	    	return 1
	    },
	    "load_translation": func(L *lua.LState) int {
	    	if L.GetTop() != 1 {
	    		return 0
	    	}
	    	t := L.CheckString(1)
	    	ts, err := CONTENT.LoadTranslation(string(t))
	    	if err != nil {
				fmt.Println("[Polyplex:raylib:ERROR]", "Translation-load failed!\n", err)
	    		return 0
	    	}
	    	CURRENT_TRANSLATION = &ts
	    	fmt.Println("[Polyplex:raylib]", "Translation has been changed to ", string(t))
	    	return 0
	    },
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetField(mod, "type_sound", lua.LNumber(0))
	L.SetField(mod, "type_music", lua.LNumber(1))
	L.Push(mod)
	return 1
}
