import { expect, test } from "bun:test";
import { arrayBinarySearch } from "../../../src/array/binary-search/ArrayBinarySearch";

test("test arrayBinarySearch", () => {
	expect(arrayBinarySearch([1, 2, 3, 4, 5, 9, 12], 8)).toBe(-1);
	expect(arrayBinarySearch([1, 2, 3, 4, 5], 5)).toBe(4);
	expect(arrayBinarySearch([1, 2, 3, 5, 9, 12, 32, 41], 41)).toBe(7);
	expect(arrayBinarySearch([1, 2, 3, 12, 32, 5], 1)).toBe(0);
});
