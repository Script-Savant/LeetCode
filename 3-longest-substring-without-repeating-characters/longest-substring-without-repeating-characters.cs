public class Solution {
    public int LengthOfLongestSubstring(string s) {
        if (string.IsNullOrEmpty(s)) return 0;

        var lastSeen = new Dictionary<char, int>();

        int left  = 0;       
        int maxLen = 0;

        for (int right = 0; right < s.Length; right++) {
            char c = s[right];

            if (lastSeen.ContainsKey(c) && lastSeen[c] >= left) {
                left = lastSeen[c] + 1;
            }

            lastSeen[c] = right;                  
            maxLen = Math.Max(maxLen, right - left + 1);
        }

        return maxLen;
    }
}