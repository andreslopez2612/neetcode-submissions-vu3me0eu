class Solution {
    public boolean isAnagram(String s, String t) {
        var newVal1 = Arrays.stream(s.split("")).sorted().collect(Collectors.joining());
        var newVal2 = Arrays.stream(t.split("")).sorted().collect(Collectors.joining());

        return newVal1.equals(newVal2);
    }
}
