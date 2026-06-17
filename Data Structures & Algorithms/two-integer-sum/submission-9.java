class Solution {
    public int[] twoSum(int[] nums, int target) {

        HashMap<Integer, Integer> seen  = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int value = nums[i];
            if (seen.containsKey(target - value)) {
                return new int[] { seen.get(target - value), i };
            }
            seen.put(value, i);
        }
        return new int[0];
    }
}
