console.info("width:" + screen.width + ", height : " + screen.height);
// console.info("width:" + window.width + ", height : ");


function loadXMLDoc()
{
    // var url = "https:www.7ethan.top/article"
    const url = "http://api.money.126.net/data/feed/0000001,1399001?callback=refreshPrice"
	var xmlhttp;
	if (window.XMLHttpRequest)
	{
		//  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
		xmlhttp=new XMLHttpRequest();
	}
	else
	{
		// IE6, IE5 浏览器执行代码
		xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
	}
	xmlhttp.onreadystatechange=function()
	{
        console.log(xmlhttp.readyState)
		if (xmlhttp.readyState == 4 && xmlhttp.status==200)
		{
            console.log("request has been succeed")
			document.getElementById("myDiv").innerHTML=xmlhttp.responseText;
		}
	}
	xmlhttp.open("GET",url,true);
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    // xmlhttp.setRequestHeader("application/json; charset=utf-8","application/x-www-form-urlencoded");
	xmlhttp.send();
	// var p = document.getElementById("innerh5")
	// p.innerHTML = 'adf'
}


let update = (args) => {
	let p = document.getElementById("innerh5")
	p.innerHTML = `<li>inner html</li>
				   <li id="a">inner pytnon</li>
				   <p id="b">Scheme</p> `
	let t = document.getElementById("innertext")
	t.innerText = 'inner text test!!!'
	console.warn("Not use args "+ args)
}

insert = ()=>{
	console.log("insert() is running")
	let list = document.getElementById("list")
    let p = document.getElementById("ipython")
    list.appendChild(p)
    //节点首先会从原先的位置删除，再插入到新的位置
}

deleted = () =>{
    console.log("deleted() is running")
    let parent = document.getElementById("list")
    // parent.removeChild("java") "Error!!!"
    parent.removeChild(parent.children[3])
}

checkForm = ()=>{
    console.log("checkForm() is running..")
    //没有name属性的<input>的数据不会被提交。
    let username = document.getElementById("username")
    let input_password = document.getElementById("input-password")
    let md5_password = document.getElementById("md5-password")

    md5_password = toString(input_password)
    console.warn("username : " + username + "md5_pwd : " + md5_password)

    return true
}
