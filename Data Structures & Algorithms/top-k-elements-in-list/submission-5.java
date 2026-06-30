class Solution {
    public int[] topKFrequent(int[] nums, int k) {
         //Almacenamos los numeros existentes
        Map<Integer, Integer> mapNums = new HashMap<>();
        //Definimos el array para retornar
        int[] result = new int[k];

        //Iteramos el array para comenzar almacenar los datos en HashMap
        for (int i : nums) {

            //Usamos mapNums.getOrDefault para prevenir un valor nulo, key, defaultValue + 1
            mapNums.put(i, mapNums.getOrDefault(i, 0) + 1);
        }

        //Creamos el minHeap para organizar de menor mayor con Map.Entry
        //PriorityQueue<Map.Entry<Integer, Integer>> minHeap = new PriorityQueue<>((a, b) -> a.getValue() - b.getValue());
        PriorityQueue<Map.Entry<Integer, Integer>> minHeap = new PriorityQueue<>(Map.Entry.comparingByValue());

        for (Map.Entry<Integer, Integer> en : mapNums.entrySet()) {
            minHeap.add(en);
            //Validamos el tamaño del minHeap
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }

        //Almacenamos en el array result
        for (int i = k - 1; i >= 0; i--) {
            result[i] = minHeap.poll().getKey();
        }

        return result;
    }
}