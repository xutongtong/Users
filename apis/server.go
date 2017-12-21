package apis

func Init() {
	//config := config.GetConfig()
	//r := NewRouter()
	//r.Run(config.GetString("server.port"))
	r := NewRouter()
	r.Run("9000")
}