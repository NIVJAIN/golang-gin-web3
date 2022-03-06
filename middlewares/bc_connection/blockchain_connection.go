// package middlewares

// import (
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/gin-contrib/gzip"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	"github.com/nivjain/golang-gin-web3/api"
// )

// //Connection is the connection created
// type Connection struct {
// 	name     string
// 	conn     *api.Api

// 	channel  *amqp.Channel
// 	exchange string
// 	queues   []string
// 	err      chan error
// 	logras   map[string]*logrus.Logger
// }

// var (
// 	connectionPool = make(map[string]*Connection)
// 	logras         map[string]*logrus.Logger
// 	log            *logrus.Logger
// )

// //NewConnection returns the new connection object
// func NewConnection(name, exchange string, queues []string, logpool map[string]*logrus.Logger) *Connection {
// 	if c, ok := connectionPool[name]; ok {
// 		return c
// 	}
// 	c := &Connection{
// 		exchange: exchange,
// 		queues:   queues,
// 		err:      make(chan error),
// 		logras:   logpool,
// 	}
// 	connectionPool[name] = c
// 	logras = logpool
// 	log = logpool["info"]
// 	return c
// }

// //GetConnection returns the connection which was instantiated
// func GetConnection(name string) *Connection {
// 	return connectionPool[name]
// }

// // Connect ...
// func (c *Connection) Connect() error {
// 	var err error

// 	var rabbitURL = ""
// 	rabbitURL = os.Getenv("RABBIT_URL")
// 	if rabbitURL == "" {
// 		// rabbitURL = "amqp://jain:jain@18.140.220.74:5672//german"
// 	}
// 	log.Info(`RabbitURL:`, rabbitURL)
// 	c.conn, err = amqp.Dial(rabbitURL)
// 	if err != nil {
// 		return err
// 		// return fmt.Errorf("Error in creating rabbitmq connection with %s : %s", err.Error())
// 	}
// 	go func() {
// 		<-c.conn.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
// 		c.err <- errors.New("Connection Closed")
// 	}()
// 	c.channel, err = c.conn.Channel()
// 	if err != nil {
// 		return fmt.Errorf("Channel: %s", err)
// 	}
// 	if err := c.channel.ExchangeDeclare(
// 		c.exchange, //exchange name
// 		// "my-exchange", //exchange name
// 		"direct", // type
// 		true,     // durable
// 		false,    // auto-deleted
// 		false,    // internal
// 		false,    // noWait
// 		nil,      // arguments
// 	); err != nil {
// 		log.Error("Error in Exchange Declare ", err)
// 		return fmt.Errorf("Error in Exchange Declare: %s", err)
// 	}
// 	return nil
// }
// client, err := ethclient.Dial("http://127.0.0.1:7545")
// if err != nil {
// 	panic(err)
// }
// conn, err := api.NewApi(common.HexToAddress("0x9Aa23BcdB51785e79EaFF318E7Db61B0befbd905"), client)
// if err != nil {
// 	panic(err)
// }