(() => {
  const fs = require("node:fs");
  const path = require("node:path");
  const readline = require("node:readline");
  const vm = require("node:vm");

  const output = [];
  Object.assign(globalThis, {
    load(filename) {
      if (filename === undefined) throw new Error("load: missing filename");
      const fullPath = path.join("v8-v7", String(filename));
      vm.runInThisContext(fs.readFileSync(fullPath, "utf8"), {
        filename: fullPath,
      });
    },
    print(...args) {
      const line = [];
      args.forEach((arg, i) => {
        if (i > 0) line.push(" ");
        line.push(String(arg));
      });
      output.push(line);
    },
  });

  function reply(message) {
    process.stdout.write(JSON.stringify(message) + "\n");
  }

  const input = readline.createInterface({ input: process.stdin });
  input.on("line", (line) => {
    try {
      const request = JSON.parse(line);
      vm.runInThisContext(request.source, { filename: request.filename });
      reply({ output });
    } catch (error) {
      reply({ output, error: String((error && error.stack) || error) });
    }
  });

  reply({ ready: true, jitless: process.execArgv.includes("--jitless") });
})();
