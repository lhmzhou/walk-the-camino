var nodeenvconfiguration = require('node-env-configuration');
var reporter = require('cucumber-html-reporter');

var args = {};
/*Example:
{
    "dir"       :"Sprint86",
    "iteration" :"P2 Iteration",
    "fileName":"report"
}*/
// print process.argv
process.argv.forEach(function (val, index, array) {
    //console.log(index + ': ' + val);
    if(val.indexOf("=") > -1){
        var keyValArray = val.split("=");
        args[keyValArray[0]] = keyValArray[1];
    }
});

console.log(args);

var defaults = {
    theme: 'bootstrap',
    jsonFile: args["directory"]+'report/in/'+args["fileName"]+'.json',
    output: args["output"]+'/report/'+args["fileName"]+'/'+args["fileName"]+'.html',
    directory:args["directory"],
    reportSuiteAsScenarios: true,
    launchReport: true,
    trimOuput :"gopro_bdd_scaffolding",
    brandTitle: 'Demo',
    name: "GoLang BDD Functional Report",
    metadata: {
        "platform": "MAC"
    }
};


// see https://github.com/whynotsoluciones/node-env-configuration 
var options = nodeenvconfiguration({
    defaults: defaults,
    prefix: 'chrApp' // read only env vars starting with CHR_APP prefix
});

reporter.generate(options);
