var b = require('octalbonescript');
b.pinMode('P9_12', b.INPUT);

setInterval(copyInputToOutput, 100);

function copyInputToOutput() {
    b.digitalRead('P9_12', writeToOutput);
    function writeToOutput(x) {
        console.log(x);
    }
}
