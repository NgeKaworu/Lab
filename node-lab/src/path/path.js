const path = require("path")

const notes = "./text.txt"
// 文件夹
console.log(path.dirname(notes))
// 文件名
console.log(path.basename(notes))
// 拓展名
console.log(path.extname(notes))
// 去除扩展名
console.log(path.basename(notes, path.extname(notes)) )

// 跨平台 拼接路径名
console.log(path.join("/", "user", notes))

// 相当于process.cwd() + notes
console.log("resolve", path.resolve(notes))

// 相当于/etc + notes
console.log("etc resolve", path.resolve('/etc', notes))

// 下面两条一样相当于process.cwd()+ /etc + notes
console.log("./etc resolve", path.resolve('./etc', notes))
console.log("./etc resolve", path.resolve('etc', notes))

// 尝试计算路径
console.log(path.normalize('/users/joe/..//test.txt'))

// node 执行目录
console.log(process.cwd())
// argv前二分别是
// node所在地
// node运行的文件
console.log(process.argv)

// 当前文件夹名
console.log(__dirname)
// 当前文件名
console.log(__filename)


