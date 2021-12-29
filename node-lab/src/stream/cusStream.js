const Stream = require('stream')

const readableStream = new Stream.Readable({
  read() {}
})
const writableStream = new Stream.Writable()

writableStream._write = (chunk, encoding, next) => {
  console.log(chunk.toString())
  next()
}

readableStream.on('readable', () => {
    console.log(readableStream.read())
  })


// 通过读流 忘 os输出流写数据
readableStream.pipe(process.stdout)
readableStream.push('hi!')
readableStream.push('ho!')


writableStream.write('hey!\n')
writableStream.end()
