const dirTree = require("directory-tree");
const path = require("path");
const fs = require("fs");

const target = "docs";

function createIndex() {
  let items = [];

  dirTree(path.join(__dirname, target), { extensions: /\.md/ }, item =>
    items.push(item)
  );

  items = items.map(item => {
    const file = path.parse(item.path);
    const dir = file.dir.split("/");
    file.topics = dir.slice(dir.indexOf(target) + 1).join("/");
    return file.name !== "README" ? file : null;
  });

  let result = "## INDEX\n";
  const lastDir = {};
  items.forEach(item => {
    const ret = parseItem(lastDir, item);
    if (ret) result += ret;
  });

  fs.writeFileSync("index.md", result, { encoding: "utf8" });
}

function parseItem(last, item) {
  if (item === undefined || item === null) return;

  if (last.topic === null) last.topic = topicExtractor(item.topics);

  if (last.topic === topicExtractor(item.topics)) {
    return addFile(item);
  }

  last.topic = item.topics;
  return addDir(item) + addFile(item);
}

function topicExtractor(topics) {
  return topics.split("/")[0];
}

function addDir(item) {
  return `\n### ${item.topics}\n\n`;
}

function addFile(item) {
  const title = readTitle(item);
  const name = title
    ? title.replace(/[\n\r]/, "").replace("# ", "")
    : item.name;
  return `- [${name}](${item.topics}/${item.name}.html)\n`;
}

function readTitle(item) {
  const lines = fs
    .readFileSync(`${item.dir}/${item.base}`, "utf-8")
    .split(/[\n\r]/);
  return lines.filter(line => line.startsWith("# "))[0];
}

/* main */
createIndex();
