const fs = require("fs");
const path = require("path");

// 创建写流 flag a+ 是append
const writeStream = fs.createWriteStream(path.join(__dirname, "./write.txt"), {
  flags: "a+",
});
// 创建读取流
const readStream = fs.createReadStream(path.join(__dirname, "./read.txt"));

// 交换数据
readStream.pipe(writeStream);
// 完成事件回调
writeStream.once("close", function () {
  console.log("All done!");
});
