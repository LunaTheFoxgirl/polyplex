package engine

var CURRENT_TRANSLATION *Translation = nil

type Translation struct {
	Language string `json:"lang"`
	Translations map[string]string `json:"translations"`
}

func (t *Translation) Translate(baseStr string) string {
	for base, translated := range t.Translations {
		if base == baseStr {
			return translated
		}
	}
	return baseStr
}