// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

const (
	defaultAuthPlugin       = "mysql_native_password"
	defaultMaxAllowedPacket = 4 << 20 // 4 MiB
	minProtocolVersion      = 10
	maxPacketSize           = 1<<24 - 1
	timeFormat              = "2006-01-02 15:04:05.999999"
)

// MySQL constants documentation:
// http://dev.mysql.com/doc/internals/en/client-server-protocol.html

// 包头标志
const (
	iOK byte = 0x00 //ok 包括 PREPARE_OK

	iAuthMoreData byte = 0x01 //
	iLocalInFile  byte = 0xfb

	iEOF byte = 0xfe
	iERR byte = 0xff
)

// https://dev.mysql.com/doc/internals/en/capability-flags.html#packet-Protocol::CapabilityFlags
// capability flags 能力标记 用于表明支持的能用的特性
type clientFlag uint32

const (
	clientLongPassword clientFlag = 1 << iota //强密码
	clientFoundRows
	clientLongFlag
	clientConnectWithDB
	clientNoSchema
	clientCompress
	clientODBC
	clientLocalFiles
	clientIgnoreSpace
	clientProtocol41
	clientInteractive
	clientSSL
	clientIgnoreSIGPIPE
	clientTransactions
	clientReserved
	clientSecureConn
	clientMultiStatements
	clientMultiResults
	clientPSMultiResults
	clientPluginAuth
	clientConnectAttrs
	clientPluginAuthLenEncClientData
	clientCanHandleExpiredPasswords
	clientSessionTrack
	clientDeprecateEOF
)

//命令列表 https://dev.mysql.com/doc/internals/en/text-protocol.html
const (
	// COM_SLEEP 内部服务器命令
	comQuit             byte = iota + 1 //1关闭/退出连接
	comInitDB                           //2切换数据库
	comQuery                            //3文本形式 立刻执行 SQL 查询 包括 select 和 增删改
	comFieldList                        //4获取数据表字段信息
	comCreateDB                         //5创建数据库
	comDropDB                           //6删除数据库
	comRefresh                          //7清楚缓存
	comShutdown                         //8停止服务器
	comStatistics                       //9获取服务器统计信息
	comProcessInfo                      //a获取当前连接的列表
	comConnect                          //b (服务器内部命令)
	comProcessKill                      //c中断一个连接
	comDebug                            //d设置调试模式,保存服务器调试信息
	comPing                             //e测试 连通性
	comTime                             // (服务器内部命令)
	comDelayedInsert                    // (服务器内部命令)
	comChangeUser                       // 重新登录(不断连接)
	comBinlogDump                       //获取二进制日志信息
	comTableDump                        //获取数据表结构信息
	comConnectOut                       // (服务器内部命令)
	comRegisterSlave                    // (从服务器向主服务器注册)
	comStmtPrepare                      //预处理 SQL 语句
	comStmtExecute                      //执行预处理语句
	comStmtSendLongData                 //发送 BLOB 类型的数据
	comStmtClose                        //销毁预处理语句
	comStmtReset                        //清楚预处理语句参数缓存
	comSetOption                        //设置语句选项
	comStmtFetch                        //获取预处理语句的执行结果
)

var cmdName = map[byte]string{
	0: "comSleep",
	1: "comQuit",      //1关闭/退出连接
	2: "comInitDB",    //2切换数据库
	3: "comQuery",     //3文本形式 立刻执行 SQL 查询 包括 select 和 增删改
	4: "comFieldList", //4获取数据表字段信息
	5: "comCreateDB",  //5创建数据库

	6:  "comDropDB",      //6删除数据库
	7:  "comRefresh",     //7清楚缓存
	8:  "comShutdown",    //8停止服务器
	9:  "comStatistics",  //9获取服务器统计信息
	10: "comProcessInfo", //a获取当前连接的列表

	11: "comConnect",     //b (服务器内部命令)
	12: "comProcessKill", //c中断一个连接
	13: "comDebug",       //d设置调试模式,保存服务器调试信息
	14: "comPing",        //e测试 连通性
	15: "comTime",        // (服务器内部命令)

	16: "comDelayedInsert", // (服务器内部命令)
	17: "comChangeUser",    // 重新登录(不断连接)
	18: "comBinlogDump",    //获取二进制日志信息
	19: "comTableDump",     //获取数据表结构信息
	20: "comConnectOut",    // (服务器内部命令)

	21: "comRegisterSlave",    // (从服务器向主服务器注册)
	22: "comStmtPrepare",      //预处理 SQL 语句
	23: "comStmtExecute",      //执行预处理语句
	24: "comStmtSendLongData", //发送 BLOB 类型的数据
	25: "comStmtClose",        //销毁预处理语句
	26: "comStmtReset",        //清楚预处理语句参数缓存
	27: "comSetOption",        //设置语句选项
	28: "comStmtFetch",        //获取预处理语句的执行结果
}

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType
type fieldType byte

// 字段类型
const (
	fieldTypeDecimal   fieldType = iota //decimal
	fieldTypeTiny                       //tiny 1B
	fieldTypeShort                      //short 2B
	fieldTypeLong                       //long
	fieldTypeFloat                      //float
	fieldTypeDouble                     //double
	fieldTypeNULL                       //null
	fieldTypeTimestamp                  //timestamp
	fieldTypeLongLong                   //long long
	fieldTypeInt24                      // 3B int24
	fieldTypeDate                       // date
	fieldTypeTime                       // time
	fieldTypeDateTime                   //datetime
	fieldTypeYear                       //year
	fieldTypeNewDate                    // newdate
	fieldTypeVarChar                    //varchar
	fieldTypeBit                        //bit
)

const (
	fieldTypeJSON       fieldType = iota + 0xf5 //json
	fieldTypeNewDecimal                         //new decimal
	fieldTypeEnum                               //enum
	fieldTypeSet                                // set
	fieldTypeTinyBLOB                           //tiny blob
	fieldTypeMediumBLOB                         // medium blob
	fieldTypeLongBLOB                           // long blob
	fieldTypeBLOB                               //blob
	fieldTypeVarString                          //var string
	fieldTypeString                             //string
	fieldTypeGeometry                           //geometry
)

type fieldFlag uint16

// 字段 标志
const (
	flagNotNULL       fieldFlag = 1 << iota // not null
	flagPriKey                              // primary key
	flagUniqueKey                           // unique
	flagMultipleKey                         // multiple key
	flagBLOB                                //blob
	flagUnsigned                            //unsigned
	flagZeroFill                            //zero fill
	flagBinary                              //binary
	flagEnum                                //enum 类型
	flagAutoIncrement                       //自增
	flagTimestamp                           //时间戳
	flagSet                                 //集合
	flagUnknown1
	flagUnknown2
	flagUnknown3
	flagUnknown4
)

// http://dev.mysql.com/doc/internals/en/status-flags.html
type statusFlag uint16

// 状态标志
const (
	statusInTrans            statusFlag = 1 << iota //a transaction is active
	statusInAutocommit                              //auto commit is enabled
	statusReserved                                  // Not in documentation
	statusMoreResultsExists                         // 更多资源
	statusNoGoodIndexUsed                           //  没有好的索引可用
	statusNoIndexUsed                               // 没有用索引
	statusCursorExists                              // 存在游标
	statusLastRowSent                               //最后一行已发送
	statusDbDropped                                 // 库已删
	statusNoBackslashEscapes                        // 无 \ 转义
	statusMetadataChanged                           //元数据已经改变
	statusQueryWasSlow                              // 慢查询
	statusPsOutParams
	statusInTransReadonly     //in a read-only transaction
	statusSessionStateChanged //connection state information has changed
)

const (
	cachingSha2PasswordRequestPublicKey          = 2
	cachingSha2PasswordFastAuthSuccess           = 3
	cachingSha2PasswordPerformFullAuthentication = 4
)
