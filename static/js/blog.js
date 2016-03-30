$(document).ready(function() {
    var offset = 1;
    var limit = 10;
    $("#pagination_next").click(function() {
        offset++;
        return false;
    });
    $("#pagination_previous").click(function() {
        if(offset>1){
            alert(offset);
        }
        return false;
    });
    //初始化的时候调用   需要修改
    getBlogList("/test",offset,limit)
    getCategoryInfo("/category")
        //实现导航lan光标聚焦
    $("#navbar .nav li").click(function(e) {
        offset = 1;
        $("#navbar .nav li").removeClass();
        $(this).addClass("active");
        if ($(this).index() != 0) {
            var url = $(this).children().attr("href");
            getBlogList("/test",offset,limit) //此处的应该要换
            return false;

        }
    });
    $.scrollUp({
        scrollName: 'scrollUp',
        scrollDistance: 300,
        scrollFrom: 'top',
        scrollSpeed: 300,
        easingType: 'linear',
        animation: 'fade',
        animationSpeed: 200,
        scrollTrigger: false,
        scrollTarget: false,
        scrollText: '',
        scrollTitle: false,
        scrollImg: false,
        activeOverlay: false,
        zIndex: 2147483647
    });
});

function getBlogList(url,offset,limit) {
    $.ajax({
        url: url,
        type: 'get',
        data: {offset:offset, size:limit},
        success: function(data) {
            var item = $.parseJSON(data);
            $("#contentx").html("").append("<p></p><p></p>");
            $.each(item.data, function(idx, obj) {
                str = sprintf(`<div class="media col-md-12">
                        <div class="media-left">
                            <a href="%s">
                                <img class="media-object" src="%s" alt="logo">
                            </a>
                        </div>
                        <div class="media-body">
                            <h3 class="media-heading"><a href="%s">%s</a></h3>
                            <p>%s...</p>
                            <p></p>
                        </div>
                        <div class="col-md-12 col-sm-12">
                        <i class="fa fa-user">%s</i>&nbsp;&nbsp;&nbsp;&nbsp;<i class="fa fa-clock-o">%s</i>&nbsp;&nbsp;&nbsp;&nbsp;<i class="fa fa-folder-open">%s</i>
                        </div>
                        <hr></hr>
                    </div>`, obj.from, obj.image, obj.from, obj.title, obj.content, obj.user, obj.timestamp, obj.category)
                $("#contentx").append(str)
            });
            if (item.count < 10) {
                $("#pagination_next").addClass("disabled");
            }
            $("#pagination_next").children().attr("href", url);
            $("#pagination_previous").children().attr("href", url);
        },
        error: function() {
            alert("异常！");
        }
    })
}

function getCategoryInfo(url) {
    $.get(url, function(result) {
        Chart.defaults.global.responsive = true;
        Chart.defaults.global.responsiveAnimationDuration = 3000;
        var ctx = document.getElementById("myChart");
        var myDoughnutChart = new Chart(ctx, {
            type: 'doughnut',
            data: $.parseJSON(result),
            options: {
                animation: {
                    animateRotate: true,
                    animateScale: true
                }
            },
        });
    });
}
