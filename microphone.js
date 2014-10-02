var b = require('bonescript');

setInterval(loop, 100);

function loop() {
    listen();
}

function listen() {
    b.analogRead('P9_39', print);
}

function print(reading) {
    console.log(reading.value);
}

// TODO: I stedet for å lese en verdi hvert 100. millisekund så burde vi lese kontinuerlig i 100ms og smoothe verdiene.