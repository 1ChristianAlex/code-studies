import { expect, test } from "bun:test";
import { arrayInsertMiddle } from "../../../src/array/insert-middle/ArrayInsertMiddle";

test("test ArrayInsertMiddle", () => {
	expect(arrayInsertMiddle([1, 2, 3, 9, 5], 8)).toEqual([1, 2, 3, 8, 9, 5]);
	expect(arrayInsertMiddle([1, 2, 3, 4], 9)).toEqual([1, 2, 9, 3, 4]);
});
