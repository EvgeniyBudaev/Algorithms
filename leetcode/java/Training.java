


public class Training {
    public static void main(String[] args) {
        UserStatus userStatus = UserStatus.ACTIVE;
        System.out.println("userStatus: " + userStatus);
    }
}

enum UserStatus implements Describable {
    ACTIVE(1, "Активный"),
    INACTIVE(2, "Неактивный"),
    BANNED(3, "Заблокирован");

    private final int code;
    private final String description;

    UserStatus(int code, String description) {
        this.code = code;
        this.description = description;
    }

    public int getCode() { return code; }

    public String getDescription() { return description; }

    @Override
    public String getInfo() {
        return "";
    }
}

interface Describable {
    String getInfo();
}