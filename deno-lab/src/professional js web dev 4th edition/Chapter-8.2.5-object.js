function Person(){
    
}

Person.prototype = {
    name: "test"
}

let p1 = new Person()

console.log(Object.entries(p1))

for (k in p1){
    console.log(k)
}