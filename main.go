package main

import (
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"

	"./common"
	"./usdt"
	"./ethereum"
	"./bitcoin"
	//  "fmt"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}


type User struct {
	UserId    int
	UserName  string
}


func main(){
	
	router := gin.Default()
	
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "chains_api",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userID := 1
			userName := loginVals.Username
			password := loginVals.Password

			if (userName == "admin" && password == "admin123@") {
				return &User{
					UserId:    userID,
					UserName:  userName,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	if err != nil {
		// log.Fatal("JWT Error:" + err.Error())
	}
    router.POST("/login", authMiddleware.LoginHandler)

	auth := router.Group("/api")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
	    auth.GET("/bitcoin/validateaddress/:address", func(c *gin.Context) {
		address := c.Param("address")
		    validInfo,err := bitcoin.ValidateAddress(address)
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "address": validInfo.Address,
				    "isvalid": validInfo.IsValid,
				    "ismine":  validInfo.IsMine,
				    "timestamp": validInfo.TimeStamp,
			    })
		    }
		})
		
	    auth.GET("/bitcoin/getbalance", func(c *gin.Context) {
		    balance,err := bitcoin.GetBalance()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "balance": balance,
			    })
		    }
	    })
    
	    auth.GET("/bitcoin/getnewaddress", func(c *gin.Context) {
		    newAddress,err := bitcoin.GetNewaAddress()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "newaddress": newAddress,
			    })
		    }
		})
		
	    auth.GET("/bitcoin/getblockcount", func(c *gin.Context) {
		    blockCount,err := bitcoin.GetBlockCount()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "blockcount": blockCount,
			    })
		    }
		})
		
	    auth.GET("/bitcoin/listaccounts", func(c *gin.Context) {
		    accounts,err := bitcoin.ListAccounts()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "accounts": accounts,
			    })
		    }
		})
		
	    auth.GET("/bitcoin/listtransactions", func(c *gin.Context) {
		    transactions,err := bitcoin.ListTransactions()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				     "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "transactions": transactions,
			    })
		    }
		})
		
	    auth.GET("/bitcoin/listaddressgroupings", func(c *gin.Context) {
		    addressGroupings,err := bitcoin.ListAddressGroupings()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
		        })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "addressGroupings": addressGroupings,
			    })
		    }
		})
		
	    auth.GET("/bitcoin/gettransaction/:txid", func(c *gin.Context) {
		    txid := c.Param("txid")
		    transaction,err := bitcoin.GetTransaction(txid)
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "transaction": transaction,
			    })
		    }
		})
		
	    auth.GET("/ethereum/gethavebalancewithaddress", func(c *gin.Context) {
		    balances,err := ethereum.GetHaveBalanceWithAddress()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "balances": balances,
			    })
		    }
		})
		
	    auth.GET("/ethereum/gettransactioncount/:address", func(c *gin.Context) {
		    address := c.Param("address")
		    result,err := ethereum.GetTransactionCount(address)
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "result": common.HexDec(result),
			    })
		    }
		})
		
	    auth.GET("/ethereum/getgasprice", func(c *gin.Context) {
		    gasPrice,err := ethereum.GetGasPrice()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "gasprice": common.HexDec(gasPrice),
			    })
		    }
		})
		
	    auth.GET("/ethereum/newblockfilter", func(c *gin.Context) {
		    filterCode,err := ethereum.NewBlockFilter()
		    if err!=nil{
			    c.JSON(500, gin.H{
			        "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "filtercode": filterCode,
			    })
		    }
		})
		
	    auth.GET("/ethereum/getfilterchanges", func(c *gin.Context) {
		    filterCode := c.Query("filtercode") 
		    ft,err := strconv.Atoi(c.Query("filtertype"))
		    var filterType ethereum.FilterType
		    switch ft{
			    case 0:
				    filterType = ethereum.BlockFilter
			    case 1:
				    filterType = ethereum.PendingTransactionFilter
			    case 2:
				    filterType = ethereum.BlockFilter
		    }
		    result,err := ethereum.GetFilterChanges(filterType,filterCode)
		    if err!=nil{
		        c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "result": result,
			    })
		    }
		})
		
	    auth.GET("/ethereum/getblockbyhash/:hash", func(c *gin.Context) {
		    hash := c.Param("hash")
		    result,err := ethereum.GetBlockByHash(hash)
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			   })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "result":result,
			    })
		    }
		})
		
	    auth.GET("/ethereum/gettransactionbyhash/:hash", func(c *gin.Context) {
		    hash := c.Param("hash")
		    result,err := ethereum.GetTransactionByHash(hash)
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "result":result,
			    })
		    }
		})
		
	    auth.GET("/usdt/getwalletaddressbalances", func(c *gin.Context) {
		    balances,err := usdt.GetWalletaddressBalances()
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "balances": balances,
			    })
		    }
	    })
	    auth.GET("/usdt/listtransactions/:address", func(c *gin.Context) {
		    address := c.Param("address")
		    transactions,err := usdt.ListTransactions(address)
		    if err!=nil{
			    c.JSON(500, gin.H{
				    "status":  "failure",
				    "err": err,
			    })
		    }else{
			    c.JSON(200, gin.H{
				    "status":  "success",
				    "transactions": transactions,
			    })
		    }
	    })
		auth.GET("/usdt/getinfo", func(c *gin.Context) {
			nodeInfo,err := usdt.Getinfo()
			if err!=nil{
				c.JSON(500, gin.H{
					"status":  "failure",
					"err": err,
				})
			}else{
				c.JSON(200, gin.H{
					"status":  "success",
					"nodeinfo": nodeInfo,
				})
			}
		})
		auth.GET("/usdt/gettransaction/:txid", func(c *gin.Context) {
			txid := c.Param("txid")
			transaction,err := usdt.GetTransaction(txid)
			if err!=nil{
				c.JSON(500, gin.H{
					"status":  "failure",
					"err": err,
				})
			}else{
				c.JSON(200, gin.H{
					"status":  "success",
					"nodeinfo": transaction,
				})
			}
		})
		auth.GET("/usdt/listblocktransactions/:index", func(c *gin.Context) {
			index,err := strconv.Atoi(c.Param("index"))
			blockTransaction,err := usdt.ListBlockTransactions(index)
			if err!=nil{
				c.JSON(500, gin.H{
					"status":  "failure",
					"err": err,
				})
			}else{
				c.JSON(200, gin.H{
					"status":  "success",
					"blocktransaction": blockTransaction,
				})
			}
		})
	}
	router.Run(":8080")

	// result,err := ethereum.GetFilterChanges(ethereum.BlockFilter,"0xfcda5ef155f385216bc0f881eebd098d")
	// if err == nil && result != nil{
		
	// }
}