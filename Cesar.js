var Cesar = /** @class */ (function() {
  function Cesar(key) {
    var words = "abcdefghijklmnopqrstuvwxyz";

    var cypher = (k, w) => {
      return [words.indexOf(w) + k].map(i => (words + words)[i])[0];
    };
    var node = w => w[0] + w[0].toUpperCase();

    this.matrix = words
      .split("")
      .reduce((a, b) => [...a, [b, cypher(key, b)]], [])
      .reduce((a, b) => ({ ...a, [node(b[0])]: node(b[1]) }), {});
  }

  Cesar.prototype.encrypt = plainText => {
    return plainText
      .split("")
      .map(word =>
        /[a-zA-Z]/.test(word)
          ? Object.keys(this.matrix)
              .filter(node => node.includes(word))
              .map(node => this.matrix[node][node.indexOf(word)])[0]
          : word
      )
      .reduce((a, b) => (a += b), "");
  };

  Cesar.prototype.decrypt = cypherText => {
    var matrix = Object.keys(this.matrix).reduce(
      (a, b) => ({ ...a, [this.matrix[b]]: b }),
      {}
    );

    return cypherText
      .split("")
      .map(word =>
        /[a-zA-Z]/.test(word)
          ? Object.keys(matrix)
              .filter(node => node.includes(word))
              .map(node => matrix[node][node.indexOf(word)])[0]
          : word
      )
      .reduce((a, b) => (a += b), "");
  };
  return Cesar;
})();

var a = new Cesar(3);

export default Cesar;
