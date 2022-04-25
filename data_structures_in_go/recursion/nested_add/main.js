function nestedAdd(array) {
	let sum = 0
	for (let i = 0; i < array.length; i++) {
		const current = array[i]
		// base case
		if (Array.isArray(current)) {
			sum += nestedAdd(current)
		} else {
			sum += current
		}
	}
	return sum
}

function init() {
	let array = [1, [4, 3], 12, 4, [1], [0, [0], 1]]
	let sum = nestedAdd(array)
	console.log("Sum of ", array, " is ", sum)

	array = []
	sum = nestedAdd(array)
	console.log("Sum of ", array, " is ", sum)

	array = [[[[[[[[[[[[[5]]]]]]]]]]]]]
	sum = nestedAdd(array)
	console.log("Sum of ", array, " is ", sum)
}

init()
