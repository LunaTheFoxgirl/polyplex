package engine
import (
	"github.com/yuin/gopher-lua"
)

type Drawable interface {
	Draw(L *lua.LState) int
	Update(L *lua.LState) int
	Init(L *lua.LState) int
}