var embed = require('./embedingo.js');

var package = "gui";
var property = "Css";
var destinationFileName = "../src/gui/CssBundle.go";
var sourceFileName = "./dist/static/css/app.css";

embed.run(package, property, destinationFileName, sourceFileName);
