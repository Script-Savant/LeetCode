public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        var seen = new Dictionary<int, int>();          

        for (int i = 0; i < nums.Length; i++)
        {
            int complement = target - nums[i];
            if (seen.TryGetValue(complement, out int j))      
                return new[] { j, i };

            seen[nums[i]] = i;                          
        }

        return Array.Empty<int>();   
    }
}