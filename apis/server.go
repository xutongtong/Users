package apis

func Init() {
	r := NewRouter()
	r.Run(":9000")
}