var embed = require('./embedingo.js');

var package = "gui";
var property = "Html";
var destinationFileName = "../src/gui/Html.go";
var sourceFileName = "index.html";

embed.run(package, property, destinationFileName, sourceFileName);
