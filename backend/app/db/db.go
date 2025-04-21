package db
import(
	"github.com/xiao-en-5970/edu-gpt/backend/app/db/sql"
	"github.com/xiao-en-5970/edu-gpt/backend/app/db/redis"
)
func InitDB(){
	sql.InitMySQL()
	redis.InitRedis()
}