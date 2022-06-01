/* Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]  */

/* Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]  */

const initialData = [3,2,2,3];

const removeElement = (nums, val) => {
	if (!nums) return 0;
	for (let i = 0; i < nums.length; i++) {
		if (nums[i] === val) {
			nums.splice(i, 1);
			i--;
		}
	}

	return nums.length;
};

console.log(removeElement(initialData, 3));
