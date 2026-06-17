class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> map = new HashMap<>();
        for (String s : strs) {
            char[] word = s.toCharArray(); //Convert to array
            Arrays.sort(word); //sort the word
            if (!map.containsKey(String.valueOf(word))) { //Check if its already exist
                map.put(String.valueOf(word), new ArrayList<>());
            }
            map.get(String.valueOf(word)).add(s);
        }

        return new ArrayList<>(map.values());
    }
}
