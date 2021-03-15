$(document).ready(function(){
    $("#bt1").click(function(){
        window.location="http://localhost:80/images/api/v1/get?"+$("#in1").val();return false;
    });
});