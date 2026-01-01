const fs = require('fs');
const path = require('path');

class Parser {
  constructor() {
    this.data = [];
  }

  readFile(filePath) {
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const lines = fileContent.split('\n');
    this.data = lines.map(line => {
      const parts = line.split(',');
      return {
        id: parts[0],
        name: parts[1],
        email: parts[2]
      };
    });
  }

  getData() {
    return this.data;
  }
}

module.exports = Parser;