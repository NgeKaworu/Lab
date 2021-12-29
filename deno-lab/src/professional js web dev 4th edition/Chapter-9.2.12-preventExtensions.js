const object1 = {};

Object.preventExtensions(object1);

try {
  object1.name ="233"
} catch (e) {
  console.log(e);
  // expected output: TypeError: Cannot define property property1, object is not extensible
}


console.log(object1)

// preventExtensions之后只是不允许添加新属性, 原有属性可修改可删除；
// seal之后不允许添加新属性， 不允许删除原有属性， 是否可修改由原属性的配置决定；
// freeze之后不允许添加新属性， 不允许删除原有属性, 不允许修改值以及属性描述符；
// 由此可见, 锁定深度是一层一层递进的；
// https://www.jianshu.com/p/b3eb24570c74
