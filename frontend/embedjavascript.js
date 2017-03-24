var embed = require('./embedingo.js');

var package = "gui";
var property = "Javascript";
var destinationFileName = "../src/gui/JavascriptBundle.go";
var sourceFileName = "build/bundle.js";

embed.run(package, property, destinationFileName, sourceFileName);
