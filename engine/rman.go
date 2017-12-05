package engine

import (
	"github.com/Member1221/raylib-go/raylib"
	"github.com/yuin/gopher-lua"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Image = raylib.Image

var CONTENT *ResourceManager = &ResourceManager{"game/"}

type ResourceManager struct {
	BaseDirectory string
}

func (r *ResourceManager) SetBaseDirectory(bDir string) {
	r.BaseDirectory = bDir
	if strings.HasSuffix(r.BaseDirectory, "/") {
		r.BaseDirectory += "/"
	}
}

func (r *ResourceManager) LoadSprite(L *lua.LState, name string) *Sprite {
	t := raylib.LoadImage(r.BaseDirectory+"textures/"+name+".png")
	return NewSprite(L, t)
}

func (r *ResourceManager) LoadSound(L *lua.LState, stype SoundType, name string) *aSound {
	fmt.Println("Going to load",r.BaseDirectory+"audio/"+name, "as", stype)
	if int(stype) > 2 {
		return nil
	}
	if stype == 0 {
		af := raylib.LoadSound(r.BaseDirectory+"audio/"+name)
		return NewSound(L, stype, af)
	} else {
		af := raylib.LoadMusicStream(r.BaseDirectory+"audio/"+name)
		return NewSound(L, stype, af)
	}
}

func (r *ResourceManager) LoadShader(L *lua.LState, name string) *Shader {
	s := raylib.LoadShader(r.BaseDirectory+"shaders/"+name+"/"+name+".vsh", r.BaseDirectory+"shaders/"+name+"/"+name+".fsh")
	return NewShader(L, s) 
}

func (r *ResourceManager) LoadDefaultShader(L *lua.LState) *Shader {
	return NewShader(L, raylib.GetShaderDefault()) 
}

func (r *ResourceManager) LoadTranslation(name string) (Translation, error) {
	f, err := ioutil.ReadFile(r.BaseDirectory+"translations/"+name+".json")
	if err != nil {
		fmt.Println("Translation-load failed!\n", err)
		return Translation{}, err
	}
	var n Translation
	json.Unmarshal(f, &n)
	return n, nil
}

func (r *ResourceManager) LoadScript(name string) (string, error) {
	f, err := ioutil.ReadFile(r.BaseDirectory+"code/"+name+".lua")
	if err != nil {
		return "", err
	}
	return string(f), nil
}
