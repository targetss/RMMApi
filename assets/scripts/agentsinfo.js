var urlGet = "/api/infoObject";

$(function (){
    $.getJSON({
        url: urlGet,
        success: function (data) {
            $('#count_site').text(" " + data.response.count_site);
            $('#count_agent').text(" " + data.response.count_agent);
            $('#online_agent').text(" " + data.response.agent_online);
            $('#offline_agent').text(" " + data.response.agent_offline);
        }
    });
});
