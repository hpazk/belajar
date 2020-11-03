// function div(num) {
//     if (Number.isFinite(1000 / num)) {
//         return "not infinite"
//     }
//     return "Infnite"
// }
// console.log(div(1))

// console.log(Math.trunc(2.8))

// let a = +"5"

// console.log(a)
// console.log(typeof a)

// use callback
// const greeting = (name) => {
//     console.log("yourname is:", name)
// }

// const userInput = callback => {
//     let name = "Hpazk";
//     callback(name)
// }

// userInput(greeting)


// basic error handling
// try {
//     textInfo()
// } catch (err) {
//     console.log('Jenis Error:', err.name)
//     console.log('Pesan Error:', err.message)
//     let stack = err.stack
//     console.log(typeof stack)
//     // console.log('Tempat Error:', stack)
// } finally {
//     console.log('Block Finally')
// }

// const textInfo = () => {
//     console.log('This is information')
// }

// custom error
class PromramClass {
    constructor(id) {
        this.id = id;
        if (id === undefined) {
            try {
                console.log(this.id);
                throw new Error('Tidak dapat membuat objek tanpa Id');
            } catch (err) {
                console.log(err.message);
            }
        }
    }
}

const program1 = new PromramClass(); // undefined
const program2 = new PromramClass(2); // 2

console.log(program1.id)
console.log(program2.id)