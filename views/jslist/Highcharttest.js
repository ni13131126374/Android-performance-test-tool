function Highcharttest() {
    Highcharts.chart('container', {
        chart: {
            type: 'spline',
            animation: Highcharts.svg, // don't animate in old IE
            marginRight: 10,
            events: {
                load: function () {
                    // set up the updating of the chart each second
                    //
                    //mem_info=JSON.parse(info)
                    var series_Total = this.series[0],
                        series_graphics = this.series[1],
                        series_nativeheap = this.series[2],
                        series_Activities = this.series[3],
                        series_Cpu=this.series[4];
                    //alert(mem_info)
                    var setInterval1 = setInterval(function () {
                        var info = '',
                            Total = '',
                            graphics = '',
                            nativeheap = '',
                            Activities = '',
                            Cpu='';
                        var PackName = formtest.PackName.value
                        var Casename=formtest.Casename.value
                        $.ajax({
                            type: "get",//方法类型
                            dataType: "json",//预期服务器返回的数据类型
                            url: "/meminfo?PackName=" + PackName+"&Casename="+Casename,//url
                            async: false,
                            success: function (result) {
                                info = result;
                                Total = Number(info.TOTAL);
                                graphics = Number(info.Graphics);
                                nativeheap = Number(info.NativeHeap);
                                Activities = Number(info.Activities);

                            }
                        },);
                        $.ajax({
                            type: "get",//方法类型
                            dataType: "json",//预期服务器返回的数据类型
                            url: "/cpu?PackName=" + PackName+"&Casename="+Casename,//url
                            async: false,
                            success: function (result) {
                                info = result;
                                Cpu=Number(info.Cpunum);
                            }
                        });
                        //alert(nativeheap)
                        //alert(info.Total)
                        var x = (new Date()).getTime();// current time
                        series_Total.addPoint([x, Total], true, true);
                        //highcharts-legend-item-hidden
                        series_graphics.addPoint([x, graphics], true, true);
                        series_nativeheap.addPoint([x, nativeheap], true, true);
                        series_Activities.addPoint([x, Activities * 100], true, true);
                        series_Cpu.addPoint([x, Cpu], true, true);

                    }, 2000);
                    $("#stop").click(function (e) {
                        clearInterval(setInterval1);
                    });
                }
            }
        },

        time: {
            useUTC: false
        },

        title: {
            text: '内存详细数据'
        },
        accessibility: {
            announceNewData: {
                enabled: true,
                minAnnounceInterval: 15000,
                announcementFormatter: function (allSeries, newSeries, newPoint) {
                    if (newPoint) {
                        return 'New point added. Value: ' + newPoint.y;
                    }
                    return false;
                }
            }
        },
        xAxis: {
            type: 'datetime',
            tickPixelInterval: 50
        },
        yAxis: {
            title: {
                text: 'Value'
            },
            plotLines: [{
                value: 10,
                width: 1,
                color: '#808080'
            }]
        },
        tooltip: {
            headerFormat: '<b>{series.name}</b><br/>',
            pointFormat: '{point.x:%Y-%m-%d %H:%M:%S}<br/>{point.y:.2f}'
        },
        legend: {
            layout: 'vertical',
            align: 'left',
            verticalAlign: 'middle'
        },
        plotOptions: {
            series: {
                label: {
                    connectorAllowed: false
                },
                pointStart: 2010
            }
        },
        Credits: {
            enabled: false
        },
        exporting: {
            enabled: false
        },
        series: [{
            name: 'Total',

            data: (function () {
                // generate an array of random data
                var data = [],
                    time = (new Date()).getTime(),
                    i;
                for (i = -50; i <= 0; i += 1) {
                    data.push({
                        x: time + i * 1000,
                        y: 0
                    });
                }
                return data;
            }())
        }, {
            name: 'graphics',
            data: (function () {

                // generate an array of random data
                var data = [],
                    time = (new Date()).getTime(),
                    i;
                for (i = -50; i <= 0; i += 1) {
                    data.push({
                        x: time + i * 1000,
                        y: 0
                    });
                }
                return data;
            }()),

            visible:false,


        }, {
            name: 'nativeheap',

            data: (function () {
                // generate an array of random data
                var data = [],
                    time = (new Date()).getTime(),
                    i;
                for (i = -50; i <= 0; i += 1) {
                    data.push({
                        x: time + i * 1000,
                        y: 0
                    });
                }
                return data;
            }()),
            visible:false
        }, {
            name: 'Activities*100',
            data: (function () {
                // generate an array of random data
                var data = [],
                    time = (new Date()).getTime(),
                    i;
                for (i = -50; i <= 0; i += 1) {
                    data.push({
                        x: time + i * 1000,
                        y: 0
                    });
                }
                return data;
            }())
        },{
            name: 'Cpu',
            data: (function () {
                // generate an array of random data
                var data = [],
                    time = (new Date()).getTime(),
                    i;
                for (i = -50; i <= 0; i += 1) {
                    data.push({
                        x: time + i * 1000,
                        y: 0
                    });
                }
                return data;
            }())
        }]
    });
}

function stop() {
    clearInterval(setInterval1)
};