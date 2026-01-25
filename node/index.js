const http = require("http");
const crypto = require("crypto");

function cpuWork() {
  const buf = crypto.randomBytes(32 * 1024); // 32KB
  return crypto.createHash("sha256").update(buf).digest("hex");
}

const server = http.createServer((req, res) => {
  if (req.url !== "/ping") {
    res.statusCode = 404;
    res.end();
    return;
  }

  // ---- CPU + memory ----
  const hash = cpuWork();

  // Random IO jitter
  const baseDelay = 5 + Math.floor(Math.random() * 15);

  setTimeout(() => {
    // slow path (1%)
    if (Math.random() < 0.01) {
      setTimeout(() => {}, 100 + Math.random() * 200);
    }

    // error simulation
    if (Math.random() < 0.002) {
      res.statusCode = 500;
      res.end("upstream failure");
      return;
    }

    console.log("Hash in Node: ", hash);
    res.setHeader("Content-Type", "application/json");
    res.end(
      JSON.stringify({
        status: "ok",
        hash,
      })
    );
  }, baseDelay);
});

server.listen(3000, () => {
  console.log("Node realistic server on :3000");
});
