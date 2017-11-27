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
          console.info("first",data);
          if(data.message){
            $("#output").text(data.message);

          }
          else
          $("#output").text(data.Message);
    }, failure: function (err) {
      $("#output").text(err.message);
    }
  });
})
 </script>
</body>
</html>
