const Emitter = require("events");
const event = new Emitter();

function cb() {
  console.log("cb");
}

function precb() {
  console.log("precb");
}

// 设置 最大次数
console.log(event.setMaxListeners(3))
// 获取 最大次数
console.log(event.getMaxListeners())

event.on("cb", cb);
event.on("cb", cb);
event.on("cb", cb);
event.on("cb", cb);
event.on("cb", cb);

// 类似 array shift
event.prependListener("cb", precb)
// 类似 array shift(一次版)
event.prependOnceListener("cb", precb)

console.log(event.eventNames("cb"));

event.emit("cb");

// 获取该事件回调数量
console.log(event.listenerCount("cb"));

// 获取该回调所有回调函数
console.log(event.listeners("cb"));

// 删一个
event.off("cb", cb);
console.log(event.listeners("cb"));
// 删所有
event.removeAllListeners("cb")
console.log(event.listeners("cb"));


