function GetCpu() {
    var result_data = ""
    var PackName = formtest.PackName.value
    var Casename=formtest.Casename.value
    $.ajax({
        type: "get",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: "/cpu?PackName=" + PackName+"&Casename="+Casename,//url
        success: function (result) {
            console.log('123', result);
            if (result.resultCode == 200) {
                result_data = result
                //console.log('123', result.resultCode)//打印服务端返回的数据(调试用)
            }
            ;
        },
        error: function () {
            alert("异常！");
        }
    });
    return result_data
}

function GetMeminfo() {
    var PackName = formtest.PackName.value
    var Casename=formtest.Casename.value
    result_data = ""
    $.ajax({
        type: "get",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: "/meminfo?PackName=" + PackName+"&Casename="+Casename,//url
        success: function (result) {

            if (result.resultCode == 200) {
                result_data = result
                //console.log('123', result.resultCode)//打印服务端返回的数据(调试用)
            }
            ;
        },
        error: function () {
            alert("异常！");
        }
    });
    return result_data
}