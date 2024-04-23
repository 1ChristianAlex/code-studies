import { describe, expect, test } from "bun:test";
import { leftJoin } from "../../../src/hashmap/left-join/left-join";

describe("test left join function", () => {
	test("test with map", () => {
		const hashMapOne = new Map([
			["diligent", "employed"],
			["font", "enamored"],
			["guide", "usher"],
			["outfit", "garb"],
			["wrath", "anger"],
		]);

		const hashMapTwo = new Map([
			["diligent", "idle"],
			["font", "averse"],
			["guide", "follow"],
			["flow", "jam"],
			["wrath", "delight"],
		]);

		const hashMapResult = [
			["diligent", "employed", "idle"],
			["font", "enamored", "averse"],
			["guide", "usher", "follow"],
			["outfit", "garb", null],
			["wrath", "anger", "delight"],
		];

		expect(leftJoin(hashMapOne, hashMapTwo)).toEqual(hashMapResult);
	});
});
