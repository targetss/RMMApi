var urlGet = "/api/infoObject";

$(function (){
    $.getJSON({
        url: urlGet,
        success: function (data) {
            console.log(data)
        }
    });
});
