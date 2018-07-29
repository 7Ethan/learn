for (i = 2;i < 101;i++){
    fg = false
    for (j = 2;j < i;j++){
        if(i %j == 0){
            fg = true
            break
        }
    }
    if (fg == false){
        console.log(i)
    }
}