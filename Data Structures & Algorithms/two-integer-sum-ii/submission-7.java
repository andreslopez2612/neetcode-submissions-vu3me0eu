class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int l = 0;
        int r = numbers.length - 1;

        while (l < r) {

            int calc = numbers[r] + numbers[l];

            if (calc == target) {
                return new int[]{l + 1, r + 1};
            }
            if (calc > target) {
                r--;
            }
            if (calc < target) {
                l++;
            }
        }

        return new int[2];
    }
}
