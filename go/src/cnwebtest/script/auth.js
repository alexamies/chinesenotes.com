// JavaScript with jQuery
$(document).ready(function() {

    // Check session on page loading
    console.log("Checking session");
    $.ajax({
      url: "/loggedin/session",
      type: "GET",
      dataType : "json",
    })
    .done(function(json) {
       console.log("Checking session json.Authenticated: " +
                   json.Authenticated);
       if (json.Authenticated == 1) {
         $("#LoginBar").hide();
         $(".authenticated").show();
         if (json.User.Role == "admin") {
            $("#Menu").text("Admin")
         }
       }
    })
    .fail(function( xhr, status, errorThrown ) {
      console.log( "Error: " + errorThrown );
      console.log( "Status: " + status );
      $("#ErrorDiv").show();
    })
    .always(function( xhr, status ) {
      console.log( "Checking session status: " + status );
    });

  // Login
  $("#LoginForm").submit(function(event) {
    $.ajax({
      url: "/loggedin/login",
      data: $("#LoginForm").serialize(),
      type: "POST",
      dataType : "json",
    })
    .done(function(json) {
      if (json.Authenticated == 1) {
        $("#LoginBar").hide();
        $(".authenticated").show();
        if (json.User.Role == "admin") {
          $("#Menu").text("Admin")
        }
      } else {
        $("#ErrorDiv").replaceWith("Sorry, your user name or password did " + 
                                   "not match");
        $("#ErrorDiv").show();
      }
    })
    .fail(function( xhr, status, errorThrown ) {
      console.log( "Error: " + errorThrown );
      console.log( "Status: " + status );
      $("#ErrorDiv").replaceWith("Sorry, there was an error");
      $("#ErrorDiv").show();
    })
    .always(function( xhr, status ) {
      console.log( "LoginForm Status: " + status );
    });
    event.preventDefault();
  });

  // Logout
  $("#LogoutLink").click(function(event) {
    $.ajax({
      url: "/loggedin/logout",
      type: "POST",
      dataType : "json",
    })
    .done(function(json) {
       $("#LoginBar").show();
       $(".authenticated").hide();
    })
    .fail(function( xhr, status, errorThrown ) {
      console.log( "Error: " + errorThrown );
      console.log( "Status: " + status );
      $("#ErrorDiv").replaceWith("Sorry, there was an error");
      $("#ErrorDiv").show();
    })
    .always(function( xhr, status ) {
      console.log( "LogoutLink Status: " + status );
    });
    event.preventDefault();
  });

});