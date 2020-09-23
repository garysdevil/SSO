function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]);
    return null; //返回参数值
}

function getCookie(name){
    var strcookie = document.cookie;//获取cookie字符串
    var arrcookie = strcookie.split("; ");//分割
    //遍历匹配
    for ( var i = 0; i < arrcookie.length; i++) {
        var arr = arrcookie[i].split("=");
        if (arr[0] == name){
            return arr[1];
        }
    }
    return "";
}
/**
 * [setCookie 设置cookie]
 * [key value t 键 值 时间(分钟)]
 */
function setCookie(key, value, time = 60, domain = window.location.hostname, path = window.location.pathname){

    var Days = time * 60 * 1000; // 分钟
    var exp = new Date(); 
    exp.setTime(exp.getTime() + Days); 
    document.cookie = key + "=" + value + ";expires="+ exp.toString() + ";path=" + path + ";domain=" + domain;
}