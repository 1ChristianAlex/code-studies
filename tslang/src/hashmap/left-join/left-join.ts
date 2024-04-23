const leftJoin = (
	hashMapOne: Map<string, string>,
	hashMapSecond: Map<string, string>,
): (string | null)[][] => {
	const result: (string | null)[][] = [];

	for (const [key, value] of hashMapOne) {
		const secondItem = hashMapSecond.get(key) || null;
		result.push([key, value, secondItem]);
	}

	return result;
};

export { leftJoin };
