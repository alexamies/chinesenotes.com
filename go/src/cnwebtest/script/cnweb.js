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
       console.log( "Result: " + json);
       if (json.Authenticated == true) {
         $("#LoginBar").hide();
         $("#LogoutBar").show();
         $("#SessionSpan").show();
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
      console.log( "Status: " + status );
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
       $("#LoginBar").hide();
       $("#LogoutBar").show();
       $("#SessionSpan").show();
    })
    .fail(function( xhr, status, errorThrown ) {
      console.log( "Error: " + errorThrown );
      console.log( "Status: " + status );
      $("#ErrorDiv").replaceWith("Sorry, there was an error");
      $("#ErrorDiv").show();
    })
    .always(function( xhr, status ) {
      console.log( "Status: " + status );
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
       $("#LogoutBar").hide();
       $("#SessionSpan").hide();
    })
    .fail(function( xhr, status, errorThrown ) {
      console.log( "Error: " + errorThrown );
      console.log( "Status: " + status );
      $("#ErrorDiv").replaceWith("Sorry, there was an error");
      $("#ErrorDiv").show();
    })
    .always(function( xhr, status ) {
      console.log( "Status: " + status );
    });
    event.preventDefault();
  });

});