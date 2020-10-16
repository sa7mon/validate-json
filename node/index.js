var currentLine = 1;

var lineReader = require('readline').createInterface({
    input: require('fs').createReadStream('sample.bson.json')
});
  
lineReader.on('line', function (line) {
    try {
        JSON.parse(line)
    } catch(ex) {
        console.log(`Line: ${currentLine}` + '\n' + line + '\n\n' + ex)
    }
    currentLine++;
    if (currentLine % 1000000 == 0) {
        console.log(`Parsed ${currentLine} lines`)
    }
});