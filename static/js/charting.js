
function LoadChartBar(labelLoad ,dataLoad,dataLoad2, textLoad, yText){
    if(myChart!=null){
        myChart.destroy();
    }
    var ctx = document.getElementById("myChart").getContext('2d');
        myChart = new Chart(ctx, {
        type: 'bar',
        data: {
        labels:  labelLoad,
        datasets: [{
            type: 'bar',
            label: textLoad,
            data: dataLoad,
            backgroundColor:'rgba(54, 162, 235, 0.2)',
            borderColor: 'rgba(54, 162, 235, 1)',
            borderWidth: 1
        },{
            type: 'line',
            label: 'thời gian xử lí cao nhất:',
            data: dataLoad2,
            backgroundColor:'#FFFFFF',
            borderColor: '#FF0000',
            borderWidth: 2,
            fill: false,
        }]
        },
        options: {
            maintainAspectRatio: false,
            scales: {
                yAxes: [{
                    ticks: {
                    beginAtZero: true,
                    callback: function(value, index, values) {
                        return formatNumber(value);
                    }
                    },    
                    scaleLabel: {
                    display: true,
                    labelString: yText,
                    }
                }]
            }
        }
        });
        
    }
function LoadCharOver(labelLoad, dataLoad){
    if(chartOver!=null){
        chartOver.destroy();
    }
    var pie = document.getElementById('chartOver1');
    chartOver = new Chart(pie, {
        type: 'pie',
        data: {
        labels:  labelLoad,
        datasets: [{
        label: '# số lần gọi',
        data: dataLoad,
        backgroundColor:['#29B0D0','#2A516E','#F07124','#CBE0E3','#979193',
        '#3366cc','#dc3d37','#fe9b3f','#3d961c','#9b4f99',
        '#804bed','#9f0b0e','#7e847c','#b6524a','#147e7e',
        '#1a8a50','#deba8a','#64df6e','#9b2e95','#60453f',
        '#a393f4','#0956e1','#12af1f','#88f8bf','#b6524a',
        '#b6bfba','#fdb53a','#fb7b69','#98d5e5','#08142d'
        ],
        borderColor: 'rgba(54, 162, 235, 1)',
        borderWidth: 1
        }]
        },
        options: {
            responsive: true,
            legend:{
                position: 'right',
            }
        }
    });
    
    }
function LoadCharOver2(labelLoad, dataLoad){
    if(chartOver2!=null){
        chartOver2.destroy();
    }
    var pie = document.getElementById('chartOver2');
    chartOver2 = new Chart(pie, {
        type: 'pie',
        data: {
        labels:  labelLoad,
        datasets: [{
        label: '# số lần gọi',
        data: dataLoad,
        backgroundColor:['#29B0D0','#2A516E','#F07124','#CBE0E3','#979193',
        '#3366cc','#dc3d37','#fe9b3f','#3d961c','#9b4f99',
        '#804bed','#9f0b0e','#7e847c','#b6524a','#147e7e',
        '#1a8a50','#deba8a','#64df6e','#9b2e95','#60453f',
        '#a393f4','#0956e1','#12af1f','#88f8bf','#b6524a',
        '#b6bfba','#fdb53a','#fb7b69','#98d5e5','#08142d'
        ],
        borderColor: 'rgba(54, 162, 235, 1)',
        borderWidth: 1
        }]
        },
        options: {
            responsive: true,
            legend:{
                position: 'right',
            }
        }
    });
    
    }
function LoadChartLine(labelLoad ,dataLoad, textLoad, yAxesLable){
    console.log(labelLoad)
    if(charRequest!=null){
        charRequest.destroy();
    }
    var ctx = document.getElementById("chartRequest").getContext('2d');
        charRequest = new Chart(ctx, {
        type: 'line',
        data: {
        labels:  labelLoad,
        datasets: [{
        label: textLoad,
        data: dataLoad,
        backgroundColor:'rgba(54, 162, 235, 0.2)',
        borderColor: 'rgba(54, 162, 235, 1)',
        borderWidth: 1
        }]
        },
        options: {
        scales: {
        yAxes: [{
            ticks: {
            beginAtZero: true,
            callback: function(value, index, values) {
                return formatNumber(value);
            }
            },    
            scaleLabel: {
            display: true,
            labelString: yAxesLable,
            }
        }]
        }
        }
        });
        
    }

 function formatNumber (value) {
    let result = `${value}`;
    const x0 = result.split('.');
    let x1 = x0[0];
    let x2 = x0.length > 1 ? `.${x0[1]}` : '';
    const rgx = /(\d+)(\d{3})/;
    while (rgx.test(x1)) {
        x1 = x1.replace(rgx, '$1,$2');
    }
    result = x1 + x2;
    return result;
}