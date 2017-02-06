/* Import dependencies, declare constants */
const lib = require("lib");
const mustache = require("mustache");

const htmlTemplate = `
<html>
  <head>
    <!-- Google Fonts -->
    <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">

    <!-- CSS Reset -->
    <link rel="stylesheet" href="//cdn.rawgit.com/necolas/normalize.css/master/normalize.css">

    <!-- Milligram CSS minified -->
    <link rel="stylesheet" href="//cdn.rawgit.com/milligram/milligram/master/dist/milligram.min.css">
  </head>
  <body>
    <div class="container" style="align:center">
      <center>
      <h1>Printer Facts</h1>
      {{fact}}

      <footer>
        <br /> <br /> <br /> <br /> <br /> <br /> <br />
        <p>Mashed together by <a href="https://www.christine.website">Christine Dodrill</a>.</p>
      </footer>
      </center>
    </div>
  </body>
</html>
`;

/**
* Your function call
* @param {Object} params Execution parameters
*   Members
*   - {Array} args Arguments passed to function
*   - {Object} kwargs Keyword arguments (key-value pairs) passed to function
*   - {String} remoteAddress The IPv4 or IPv6 address of the caller
*
* @param {Function} callback Execute this to end the function call
*   Arguments
*   - {Error} error The error to show if function fails
*   - {Any} returnValue JSON serializable (or Buffer) return value
*/
module.exports = (params, callback) => {
    lib[".main"]((err, result) => {
        if(err != null) {
            callback(err);
        }

        callback(null, new Buffer(mustache.render(htmlTemplate, {fact: result.facts[0]})));
    });
};
