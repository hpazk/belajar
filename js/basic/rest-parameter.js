// function rest_parameter(...args) {
//     // console.log(args)

//     let value = args.reduce((acc, arg) => {
//         return acc + arg
//     })

//     return value
// }

// const nil = rest_parameter(1, 2, 3, 4, 5)

// console.log(nil)

class RestParameter {
    constructor() {}
    basic(...args) {
        return args
    }
}

const arr = new RestParameter()

console.log(arr.basic(1, 2, 3, 4, 5))

const obj = new RestParameter()

console.log(obj.with_object({
    ID: 1,
    Name: 'hpazk',
}))