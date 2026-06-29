class Solution {

   public String encode(List<String> strs) {
        StringBuilder newStringBuilder = new StringBuilder();
        for (String s : strs){
            newStringBuilder.append(s.length());
            newStringBuilder.append("#");
            newStringBuilder.append(s);
        }
        return newStringBuilder.toString();
    }

    public List<String> decode(String str) {
        List<String> result = new ArrayList<>();
        int i = 0;

        while (i < str.length()) {
            int j = str.indexOf('#', i);

            int length = Integer.parseInt(str.substring(i, j));

            String res = str.substring(j + 1, j + 1 + length);
            result.add(res);

            i = j + 1 + length;
        }
        return result;
    }
}
