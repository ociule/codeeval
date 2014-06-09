

public class MultiplicationTables {
    public static void main (String[] args) {
        int x = 12;
        int y = 12;

        for (int i = 1; i<= x; i++) {
            for (int j = 1; j<= y; j++) {
                System.out.format("%4d", i * j);
            }
            System.out.println("");
        }
    }
}
