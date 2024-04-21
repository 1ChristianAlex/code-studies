const arrayInsertMiddle = <T>(list: T[], newValue: T): T[] => {
	const middleIndex = Math.ceil(list.length / 2);
	const newList = new Array<T>(list.length + 1);
	newList[middleIndex] = newValue;

	for (const [index, value] of list.entries()) {
		if (index >= middleIndex) {
			newList[index + 1] = value;
		} else {
			newList[index] = value;
		}
	}

	return newList;
};

export { arrayInsertMiddle };
