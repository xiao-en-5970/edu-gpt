package bcrypts
import "golang.org/x/crypto/bcrypt"

// 加密密码（自动生成salt）
func HashPassword(password string) (string, error) {
    HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(HashPassword), err
}

// 验证密码
func CheckPasswordHash(password, HashPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(password))
    return err == nil
}