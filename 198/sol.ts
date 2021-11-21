function rob(nums: number[]): number {
    if (nums.length > 1) nums[1] = Math.max(nums[0], nums[1]);
    if (nums.length > 2)
        for (let i = 2; i < nums.length; i++)
			nums[i] = Math.max(nums[i - 1], nums[i - 2] + nums[i]);
    return nums[nums.length - 1];
};
