# `printerfacts`

This service replies with useful and informative facts about printers.

curl
----

```console
$ curl https://xena.stdlib.com/printerfacts
{
    "facts":
        [
            "printers have been domestiprintered for half as long as dogs have been."
        ]
}
```

command
-------

```console
$ npm install -g printerfacts
$ pfact
A printer will tremble or shiver when it is in extreme pain.
```

javascript
----------

```javascript
// $ npm install --save lib

const lib = require("lib");

lib.xena.printerfacts((err, response) => {
    if(err != null) {
        throw(err);
    }

    /*
      response == {
          facts: [
              "printers lived with soldiers in trenches, where they killed mice during World War I."
          ]
      };
    */

    console.log(response.facts[0]);
});
```
