class Solution {
    public boolean isPalindrome(String s) {

        char[] text = s.replaceAll(" ", "").toCharArray();
        List<Character> newText = new ArrayList<>();
        for (char str : text) {
            if (letter(str)) {
                newText.add(Character.toLowerCase(str));
            }
        }

        int l = 0;
        int r = newText.size() - 1;

        while (l < r) {
            if (newText.get(l).equals(newText.get(r))) {
                l++;
                r--;
            } else {
                return false;
            }
        }

        return true;
    }

    public boolean letter(char c) {

        if (c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '0' && c <= '9') {
            return true;
        }
        return false;
    }
}
