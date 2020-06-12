$(document).ready(function(){
    $("#bt1").click(function(){
        window.location="http://localhost:3001/images/api/v1/get?"+$("#in1").val();return false;
    });
});