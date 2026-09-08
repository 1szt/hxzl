import javax.swing.*;
import java.awt.*;

public class Main {
    public static void main(String[] args) {
        SwingUtilities.invokeLater(() -> {
            JFrame frame = new JFrame("Hello mcup-server");
            frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
            frame.setSize(480, 320);
            frame.setLocationRelativeTo(null);

            JPanel panel = new JPanel();
            panel.setBackground(new Color(241, 241, 252));
            panel.setLayout(new GridBagLayout());

            GridBagConstraints gbc = new GridBagConstraints();
            gbc.gridx = 0;
            gbc.gridy = 0;
            gbc.insets = new Insets(10, 0, 10, 0);

            JLabel title = new JLabel("Hello mcup-server!");
            title.setFont(new Font("Microsoft YaHei", Font.BOLD, 28));
            title.setForeground(new Color(87, 85, 217));
            panel.add(title, gbc);

            gbc.gridy = 1;
            JLabel sub = new JLabel("欢迎使用 Java Swing 客户端");
            sub.setFont(new Font("Microsoft YaHei", Font.PLAIN, 16));
            sub.setForeground(new Color(69, 80, 96));
            panel.add(sub, gbc);

            gbc.gridy = 2;
            gbc.insets = new Insets(20, 0, 0, 0);
            JButton btn = new JButton("点击我");
            btn.setFont(new Font("Microsoft YaHei", Font.PLAIN, 14));
            btn.setBackground(new Color(87, 85, 217));
            btn.setForeground(Color.WHITE);
            btn.setFocusPainted(false);
            btn.setBorderPainted(false);
            btn.setCursor(Cursor.getPredefinedCursor(Cursor.HAND_CURSOR));
            btn.addActionListener(e -> JOptionPane.showMessageDialog(frame, "你好！来自 mcup-server 😊"));
            panel.add(btn, gbc);

            frame.setContentPane(panel);
            frame.setVisible(true);
        });
    }
}
