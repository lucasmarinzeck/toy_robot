import express from "express";
import { fileURLToPath } from "url";
import { dirname } from "path";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const app = express();

app.get("/", (req, res) => {
  res.sendFile("index.html", { root: __dirname });
});

app.get("/index.js", (req, res) => {
  res.sendFile("index.js", { root: __dirname });
});

app.listen(8080, function () {
  console.log("Server running on http://localhost:8080/");
});
