<html>
<head>
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>

</head>
<body>
<h2> HI </h2>
<input class="form-control in" id="message" name="" value="">
<button id="submit">Submit</button>
<h4 id ="output"></h4>
<script>
$("#submit").on('click',() => {
  console.log('hello');
  $.ajax({
      url: 'chat',
      type: 'post',
      data: JSON.stringify({
          message: $("#message").val()
      }),
      headers: {
        Authorization: "fwefeuiwfwejkfjkwefjkew"
      },
      contentType:"application/json; charset=utf-8",
      dataType: 'json',
      success: function (data) {
        $("#output").css("color", "black")

          console.info("first",data);
          if(data.message){
            $("#output").text(data.message);

          }
          else
            $("#output").text(data.Message);
    }, error: function (err) {
      console.info("error",err);
      $("#output").css("color", "red")
      $("#output").text(err.responseJSON.message);
    }
  });
})
 </script>
</body>
</html>
