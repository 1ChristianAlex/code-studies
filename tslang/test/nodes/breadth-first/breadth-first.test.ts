import { describe, expect, test } from "bun:test";
import { NodeBreadth } from "../../../src/nodes/breadth-first/breadth-first";
import { ListNode } from "../../../src/nodes/node-item";

describe("test breadth-first", () => {
	test("teste with node one ", () => {
		const five = new ListNode(5);
		const four = new ListNode(4);
		const one = new ListNode(1);
		const two = new ListNode(2);
		const three = new ListNode(3);
		const zero = new ListNode(0);

		four.next.push(five);

		one.next.push(four);
		two.next.push(four);
		three.next.push(four);

		zero.next.push(one, two, three);

		const instance = new NodeBreadth();

		expect(instance.breadthFirst(zero)).toEqual([0, 1, 2, 3, 4, 5]);
	});

	test("teste with node 6", () => {
		const five = new ListNode(5);
		const four = new ListNode(4);
		const six = new ListNode(6);
		const two = new ListNode(2);
		const three = new ListNode(3);
		const zero = new ListNode(0);

		four.next.push(five);

		two.next.push(four);

		zero.next.push(six, two, three);

		const instance = new NodeBreadth();

		expect(instance.breadthFirst(zero)).toEqual([0, 6, 2, 3, 4, 5]);
	});

	test("teste with node 7", () => {
		const zero = new ListNode(0);
		const one = new ListNode(1);
		const two = new ListNode(2);
		const three = new ListNode(3);
		const four = new ListNode(4);
		const five = new ListNode(5);
		const six = new ListNode(6);
		const seven = new ListNode(7);
		const eight = new ListNode(8);

		zero.next.push(one, two, three);
		one.next.push(four);
		two.next.push(four);
		three.next.push(four);

		four.next.push(seven, five);
		five.next.push(seven, six);
		seven.next.push(eight);

		const instance = new NodeBreadth();

		expect(instance.breadthFirst(zero)).toEqual([0, 1, 2, 3, 4, 7, 5, 8, 6]);
	});
});
