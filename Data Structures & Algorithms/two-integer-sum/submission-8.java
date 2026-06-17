class Solution {
    public int[] twoSum(int[] nums, int target) {

        int[] result = new int[2];
        HashMap<Integer, Integer> seen  = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int value = nums[i];
            if (seen.containsKey(target - value)) {
                result[0] = seen.get(target-value);
                result[1] = i;
            }
            seen.put(value, i);
        }

        return result;
    }
}
