class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> check = new HashSet<>();

        for (int i : nums) {
            check.add(i);
        }

        int res = 0;

        for (int num : check) {
            if (!check.contains(num - 1)) {
                int current = 1;
                int next = num + 1;

                while (check.contains(next)) {
                    current++;
                    next++;
                }

                res = Math.max(res, current);
            }
        }

        return res;
    }
}
