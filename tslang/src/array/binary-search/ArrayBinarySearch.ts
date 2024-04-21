const arrayBinarySearch = (list: number[], target: number) => {
	let highestIndex = list.length - 1;
	let lowestIndex = 0;

	while (lowestIndex <= highestIndex) {
		const middleIndex = Math.round((lowestIndex + highestIndex) / 2);
		const middleItem = list[middleIndex];

		if (target === middleItem) {
			return middleIndex;
		}

		if (middleItem > target) {
			highestIndex = middleIndex - 1;
		} else {
			lowestIndex = middleIndex + 1;
		}
	}

	return -1;
};

export { arrayBinarySearch };
