package utils

func title(titleEsc string) string {
	switch titleEsc {
	case "Sequential programming":
		return "顺序编程"
	case "Structured programming":
		return "结构化编程"
	}
	return ""
}
