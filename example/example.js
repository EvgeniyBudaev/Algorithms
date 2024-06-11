const example = (nums, k) => {
   let left = 0, curr = 1, ans = 0;

   for (let right = 0; right < nums.length; right++) {
       curr *= nums[right];
       console.log(curr);

       while (curr >= k) {
           curr /= nums[left];
           left++;
       }

       ans += right - left + 1;
   }

   return ans;
}

console.log(example([10, 5, 2, 6], 100)); // 5