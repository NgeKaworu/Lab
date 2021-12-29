const fs = require("fs");
const path = require("path");
const filename = "./text.txt";
const fullpath = path.join(__dirname, filename);
const content = "Some content!";

fs.writeFile(fullpath, content, { flag: "a+" }, (err) => {
  if (err) {
    console.error(err);
    return;
  }
  //file written successfully
});

try {
  const data = fs.writeFileSync(fullpath, content, { flag: "a+" });
  //file written successfully
} catch (err) {
  console.error(err);
}

fs.readFile(fullpath, "utf8", (err, data) => {
  if (err) {
    console.error(err);
    return;
  }
  console.log(data);
});

try {
  const data = fs.readFileSync(fullpath, "utf8");
  console.log(data);
} catch (err) {
  console.error(err);
}
