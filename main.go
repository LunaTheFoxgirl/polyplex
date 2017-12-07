package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/Member1221/polyplex/engine"
	//"github.com/Member1221/raylib-go/raylib"
	"os"
	"path/filepath"
)

var state *lua.LState

func main() {
	fmt.Println("[Polyplex:raylib]", "Loading Lua engine...")
	state = lua.NewState()
	fmt.Println("[Polyplex:raylib]", "Implementing simple class support...")
	state.DoString(`function class(base, init)
   local c = {}    -- a new class instance
   if not init and type(base) == 'function' then
      init = base
      base = nil
   elseif type(base) == 'table' then
    -- our new class is a shallow copy of the base class!
      for i,v in pairs(base) do
         c[i] = v
      end
      c._base = base
   end
   -- the class will be the metatable for all its objects,
   -- and they will look up their methods in it.
   c.__index = c

   -- expose a constructor which can be called by <classname>(<args>)
   local mt = {}
   mt.__call = function(class_tbl, ...)
   local obj = {}
   setmetatable(obj,c)
   if init then
      init(obj,...)
   else 
      -- make sure that any stuff from the base class is initialized!
      if base and base.init then
      base.init(obj, ...)
      end
   end
   return obj
   end
   c.init = init
   c.is_a = function(self, klass)
      local m = getmetatable(self)
      while m do 
         if m == klass then return true end
         m = m._base
      end
      return false
   end
   setmetatable(c, mt)
   return c
end`)
	
	
	engine.RegisterWindowType(state)
	engine.RegisterCameraType(state)
	engine.RegisterTextureType(state)
	engine.RegisterSpriteType(state)
	engine.RegisterShaderType(state)
	engine.RegisterSoundType(state)
	
	fmt.Println("[Polyplex:raylib]", "Preheating module core...")
	state.PreloadModule("core", engine.ModuleCore)
	
	fmt.Println("[Polyplex:raylib]", "Preheating module input...")
	state.PreloadModule("input", engine.ModuleInput)
	
	fmt.Println("[Polyplex:raylib]", "Preheating module translation...")
	state.PreloadModule("translathor", engine.ModuleTranslation)
	
	fmt.Println("[Polyplex:raylib]", "Preheating module content...")
	state.PreloadModule("content", engine.ModuleResources)
		
	fmt.Println("[Polyplex:raylib]", "Preheating module utils...")
	state.PreloadModule("utils", engine.ModuleUtilities)
	
	fmt.Println("[Polyplex:raylib]", "Lua engine has been loaded!")
	fmt.Println("[Polyplex:raylib]", "Preheating game.lua...")
	ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    
    exPath := filepath.Dir(ex)
    
	fPath := exPath+"/"+engine.CONTENT.BaseDirectory+"code/?.lua;"
    fmt.Println("[Polyplex:raylib]", "Set package path to:", fPath);
	state.DoString("package.path = [[" + fPath + "]] .. [[/subdir/?.lua;]] .. package.path")
	
	scr, err := engine.CONTENT.LoadScript("game")
	if err != nil {
		fmt.Println("[Polyplex:raylib:ERROR]", err)
		return
	}
	err = state.DoString(scr)
	if err != nil {
		fmt.Println("[Polyplex:raylib:ERROR]", err)
		return
	}
	fmt.Println("[Polyplex:raylib]", "Starting game ( running main() )....\n\n-----------APP BEGIN-----------")
	err = state.CallByParam(lua.P{ Fn: state.GetGlobal("main"), NRet: 0, Protect: true})
	if err != nil {
		fmt.Println("[Polyplex:raylib:ERROR]", "Polyplex failed finding an main function\n\n", err)
		return
	}
	fmt.Println("------------APP END------------\n\nHave a nice day! c:")
}
