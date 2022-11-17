var urlGetinfo = "/api/infoObject";
let urlGetSite = "/api/site"

$(function (){
    $.getJSON({
        url: urlGetinfo,
        success: function (data) {
            $('#count_site').text(" " + data.response.count_site);
            $('#count_agent').text(" " + data.response.count_agent);
            $('#online_agent').text(" " + data.response.agent_online);
            $('#offline_agent').text(" " + data.response.agent_offline);
        }
    });
});

let arrSite = [];

$(function (){
   $.getJSON({
       url: urlGetSite,
       success: function (data){
           arrSite = data;
           for(var i = 0; i < arrSite.length; i++){
               let oldHtml = $('#site').html();
               $('#site').html(oldHtml + "<li>" + arrSite[i].name + "</li>");
           };
       }
   });
});