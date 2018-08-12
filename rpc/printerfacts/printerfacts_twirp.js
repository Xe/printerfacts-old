// Code generated by protoc-gen-twirp_browserjs v0.0.1, DO NOT EDIT.
// source: printerfacts.proto

// _request takes the HTTP method, URL path (this will usually contain the domain name)
// json body that will be sent to the server, a callback on successful requests and a
// callback for requests that error out.
var _request = function(method, path, body, onSuccess, onError) {
  var xhr = new XMLHttpRequest();
  xhr.open(method, path, true);
  xhr.setRequestHeader("Accept","application/json");
  xhr.setRequestHeader("Content-Type","application/json");

  xhr.onreadystatechange = function (e) {
    if (xhr.readyState == 4) {
      if (xhr.status == 204 || xhr.status == 205) {
        onSuccess();
      } else if (xhr.status == 200) {
        var value = JSON.parse(xhr.responseText);
        onSuccess(value);
      } else {
        var value = JSON.parse(xhr.responseText);
        onError(value);
      }
    }
  };

  if (body != null) {
    xhr.send(JSON.stringify(body));
  } else {
    xhr.send(null);
  }
};

// methods for PrinterfactsClient
/*
    Printerfacts manages amusing facts about printers.
*/


/*
    Fact gives the client one or more facts about printers.
*/
var Printerfacts_fact = function(server_address, fact_params, onSuccess, onError) {
  var full_method = server_address + "/twirp/" + "us.xeserv.api.Printerfacts" + "/" + "Fact";
  _request("POST", full_method, fact_params, onSuccess, onError);
};
