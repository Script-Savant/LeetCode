public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        int n = nums.Length;

        Dictionary<int, int> numsToIndex = new();

        if (n == 0){
            return new int[0]{};
        }

        for (int i = 0; i < n; i++)
        {
            int complement = target - nums[i];
            if (numsToIndex.ContainsKey(complement))
            {
                return new int[2]{numsToIndex[complement], i};
            }
            numsToIndex[nums[i]] = i;
        }

        return new int[0]{};
    }
}