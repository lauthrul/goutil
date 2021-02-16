package lang

const (
	CN = "cn"
	EN = "en"
)

func Text(lang, key string) string {
	switch lang {
	case EN:
		return langEN[key]
	default:
		return langCN[key]
	}
}
