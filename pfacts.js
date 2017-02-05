#!/usr/bin/env node

const lib = require("lib");

lib.xena.printerfacts((err, response) => {
    if(err != null) {
        throw(err);
    }

    console.log(response.facts[0]);
});
