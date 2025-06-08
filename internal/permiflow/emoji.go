package permiflow

var NoEmoji bool

func Emoji(s string) string {
	if NoEmoji {
		return ""
	}
	return s
}
