package infra
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"xorm.io/core"
)

var engine *xorm.Engine
