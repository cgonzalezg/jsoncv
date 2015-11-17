var jade = require('jade');
var fn = jade.compileFile('./index.jade', {});
var JSONcv = require('./cv.json');
// var fn = jade.compile(jadeTemplate);
var htmlOutput = fn({
  cv: JSONcv
});

console.log(htmlOutput);
