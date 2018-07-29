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
	xmlhttp.open("GET","https://api.douban.com/v2/book/1220562",true);
	xmlhttp.send();
};
p();