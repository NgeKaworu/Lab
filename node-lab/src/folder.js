const fs = require("fs");
const path = require("path");

// 相当于ls
console.log(fs.readdirSync(__dirname));

const isFile = (fileName) => {
  // fs.lstatSync 同步返回文件详细信息
  return fs.lstatSync(fileName).isFile();
};

console.log(
  fs
    .readdirSync(__dirname)
    .map((fileName) => {
      return path.join(__dirname, fileName);
    })
    .filter(isFile)
);

fs.realpath("./app.js", console.log);
