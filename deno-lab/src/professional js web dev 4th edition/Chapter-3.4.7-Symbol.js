// let genericSymbol = Symbol();
// let otherGenericSymbol = Symbol();

// let fooSymbol = Symbol("foo")
// let otherFooSymbol = Symbol("foo")

// console.log(genericSymbol == otherFooSymbol)
// console.log(fooSymbol == otherFooSymbol)

// let s1 = Symbol("foo"),
//   s2 = Symbol("bar");

// let o = {
//   [s1]: "foo val",
//   [s2]: "bar val",
//   baz: "baz val",
//   qux: "qux val",
// };

// let sn = Object.getOwnPropertySymbols(o);

// console.log(Object.getOwnPropertyDescriptors(o));
// console.log(Object.getOwnPropertyNames(o));
// console.log(sn);
// console.log(o[sn[0]]);
// console.log(Object.getPrototypeOf(o));

let initArr = ["test"];
let arrayLikeObject = { length: 1, 1: "233", k: "8s", 0: "baz" };
console.log(arrayLikeObject[Symbol.isConcatSpreadable]);
console.log(initArr.concat(arrayLikeObject));
arrayLikeObject[Symbol.isConcatSpreadable] = true;
console.log(initArr.concat(arrayLikeObject));
