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
	xmlhttp.send();
};
p();