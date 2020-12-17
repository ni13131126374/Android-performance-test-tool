function ResultMeminfo() {
var dom = document.getElementById("Result");
var myChart = echarts.init(dom,'light');
var CPUdom = document.getElementById("CpuResult");
var myCPUChart = echarts.init(CPUdom,'light');
var PackName = formtest.PackName.value;
var Casename=formtest.Casename.value;
var TOTALMAP={};
var CpuTOTAL={};
var CpuAverage=0.00;
var Meminfolist=new Array();
var Meminfokeylist=new Array();
var Meminfovalue=new Array();
var CPUkeylist=new Array();
var CPUvalue=new Array();
    $.ajax({
        type: "get",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: "/ResultMeminfo?PackName=" + PackName+"&Casename="+Casename,//url
        async: false,
        success: function (result) {
            TOTALMAP=result.TOTALMAP
            Meminfolist=result.Meminfolist
            for (var key in TOTALMAP){
                Meminfokeylist.push(key*100)
                Meminfovalue.push(TOTALMAP[key])
            }

            document.getElementById('TotalMaxvalue').innerText=Meminfolist[0]/1024;
            document.getElementById('TotalMinvalue').innerText=Meminfolist[1]/1024;
            document.getElementById('TotalMedian').innerText=Meminfolist[2]/1024;
            document.getElementById('TotalAverage').innerText=Meminfolist[3]/1024;
            document.getElementById('DiffMax').innerText=Meminfolist[4]/1024;
            document.getElementById('GraphicsMaxvalue').innerText=Meminfolist[5]/1024;
            document.getElementById('activetMax').innerText=Meminfolist[6];

        }
    },);
    $.ajax({
        type: "get",//方法类型
        dataType: "json",//预期服务器返回的数据类型
        url: "/Resultcpu?PackName=" + PackName+"&Casename="+Casename,//url
        async: false,
        success: function (result) {
            CpuAverage=result.CpuAverage
            CpuTOTAL=result.CpuTOTAL
            for (var key in CpuTOTAL){
                CPUkeylist.push(key)
                CPUvalue.push(CpuTOTAL[key])
            }
        }
    },);
    document.getElementById('CpuAverage').innerText=CpuAverage;
    option = null;
    option = {
    title: {
        text: '内存数据分布（MB）',
        subtext: '数据来自测试数据'
    },
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'shadow'
        }
    },
    legend: {
        data: ['2011年', '2012年']
    },
    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
    },
    xAxis: {
        type: 'value',
        boundaryGap: [0, 0.01]
    },
    yAxis: {
        type: 'category',
        data: Meminfokeylist
    },
    series: [
        {
            name: '内存',
            type: 'bar',
            data: Meminfovalue

        },

    ]
};
;


    if (option && typeof option === "object") {
    myChart.setOption(option, true);
    }
    cpuoption = null;
    cpuoption = {
        title: {
            text: 'CPU数据分布（%）',
            subtext: '数据来自测试数据'
        },
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'shadow'
            }
        },
        legend: {
            data: ['2011年', '2012年']
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        xAxis: {
            type: 'value',
            boundaryGap: [0, 0.01]
        },
        yAxis: {
            type: 'category',
            data: CPUkeylist
        },
        series: [
            {
                name: 'Cpu',
                type: 'bar',
                data: CPUvalue

            },

        ]
    };
    ;
    if (cpuoption && typeof cpuoption === "object") {
        myCPUChart.setOption(cpuoption, true);
    }
}