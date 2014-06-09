

public class Main {
    public static void main (String[] args) {
        int x = 12;
        int y = 12;

        for (int i = 1; i<= x; i++) {
            StringBuffer sb = new StringBuffer();
            for (int j = 1; j<= y; j++) {
                sb.append(String.format("%4d", i * j));
            }
            System.out.println(sb.toString().trim());
        }
    }
}
