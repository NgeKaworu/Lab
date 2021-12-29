// emit是同步的 


// 输出hi 1和hi 2
// const EventEmitter = require('events');
// let emitter = new EventEmitter();
// emitter.on('myEvent', () => {
//   console.log('hi 1');
// });
// emitter.on('myEvent', () => {
//   console.log('hi 2');
// });
// emitter.emit('myEvent');

// 死循环
// const EventEmitter = require('events');
// let emitter = new EventEmitter();
// emitter.on('myEvent', () => {
//   console.log('hi');
//   emitter.emit('myEvent');
// });
// emitter.emit('myEvent');

// 输出hi 一次 不死循环
const EventEmitter = require('events');
let emitter = new EventEmitter();
emitter.on('myEvent', function sth () {
  emitter.on('myEvent', sth);
  console.log('hi');
});
emitter.emit('myEvent');
// 但每次都额外注册一次函数
// emitter.emit('myEvent');
// emitter.emit('myEvent');