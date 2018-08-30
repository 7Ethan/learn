var p = function loadJsonDoc(){
    var xmlhttp;
    xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange=function()
	{
		if (xmlhttp.readyState==4 && xmlhttp.status==200)
		{
            console.log(xmlhttp.responseText)
		}
	}
	xmlhttp.open("GET","https://www.7ethan.top/article",true);

	/**
	 * 
	==============  gin server =============
	router := gin.Default() // *gin.Engine
	api := router.Group("v1")
	api.GET("ping", func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	s := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	*/

	let sender = ()=>xmlhttp.send();
	setTimeout(sender(),10);
	console.log(xmlhttp.responseText);
};
p();